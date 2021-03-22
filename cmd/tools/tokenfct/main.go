package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/videocoin/go-bridge/cmd/common"
	"github.com/videocoin/go-bridge/erc20"
	"golang.org/x/sync/errgroup"
)

var (
	ownerkey = pflag.String("owner", "owner.json", "file with account with private keys")
	ownerpw  = pflag.String("ownerpw", "",
		"file with owner key password. if file is not provided - app will try to decrypt with empty pw")
	config = pflag.StringP("config", "c", "faucetcfg.json", "configuration for faucet helper")
	num    = pflag.IntP("number", "n", 10, "number of keys to generate and to faucet (if key exist it will just fauceted)")
	gas    = pflag.IntP("gas", "g", 100, "one unit is 1e12 wei")
	tokens = pflag.IntP("tokens", "t", 100, "one unit is 1e18 wei")

	gasPrecision   = big.NewInt(1e12)
	tokenPrecision = big.NewInt(1e18)
)

type conf struct {
	DataDir       string
	URL           string
	LogLevel      string
	TokenContract ethcommon.Address
}

func must(log *logrus.Entry) func(err error) {
	return func(err error) {
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func main() {
	pflag.Parse()

	c := conf{}
	common.MustParseConfig(&c, "", *config)

	logger := logrus.New()
	lvl, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		fmt.Printf("failed to parse level %s: %v", c.LogLevel, err)
		os.Exit(1)
	}
	logger.SetLevel(lvl)
	log := logrus.NewEntry(logger)
	fatal := must(log)

	key := common.MustDecryptKey(log, *ownerkey, *ownerpw)

	client, err := ethclient.Dial(c.URL)
	fatal(err)

	erc, err := erc20.NewERC20(c.TokenContract, client)
	fatal(err)

	gn := new(big.Int).SetInt64(int64(*gas))
	gn = gn.Mul(gn, gasPrecision)

	tn := new(big.Int).SetInt64(int64(*tokens))
	tn = tn.Mul(tn, tokenPrecision)

	f := faucet{
		log:    log,
		erc:    erc,
		client: client,
		owner:  *bind.NewKeyedTransactor(key.PrivateKey),
	}
	ctx := context.Background()
	fatal(f.validate(ctx, gn, tn, *num))

	k := keys{
		log:     log,
		dir:     c.DataDir,
		pattern: "worker-",
	}
	existing, err := k.existing()
	fatal(err)
	pkeys := make(chan *ecdsa.PrivateKey, *num)
	group, gctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return k.scan(gctx, *num, pkeys)
	})
	group.Go(func() error {
		return k.generate(gctx, existing, *num, pkeys)
	})
	go func() {
		if err := group.Wait(); err != nil {
			fatal(err)
		}
		close(pkeys)
	}()
	fatal(f.fund(ctx, pkeys, gn, tn))
	fmt.Println("Keys generated and funded with gas and tokens.")
}

type keys struct {
	log     *logrus.Entry
	dir     string
	pattern string
}

func (k keys) existing() (int, error) {
	n := 0
	return n, filepath.Walk(k.dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(info.Name(), k.pattern) {
			n++
		}
		return nil
	})
}

func (k keys) scan(ctx context.Context, total int, scanned chan<- *ecdsa.PrivateKey) error {
	group, ctx := errgroup.WithContext(ctx)
	n := 0
	if err := filepath.Walk(k.dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if n == total {
			return nil
		}
		n++
		fpath := filepath.Join(k.dir, info.Name())
		group.Go(func() error {
			key := common.MustDecryptKey(k.log, fpath, "")
			k.log.Debugf("found with address 0x%x. at %s", key.Address, fpath)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case scanned <- key.PrivateKey:
			default:
			}
			return nil
		})
		return nil
	}); err != nil {
		return err
	}
	return group.Wait()
}
func (k keys) generate(ctx context.Context, current int, total int, generated chan<- *ecdsa.PrivateKey) error {
	group, ctx := errgroup.WithContext(ctx)
	for i := current; i < total; i++ {
		name := fmt.Sprintf("%s%d.json", k.pattern, i)
		group.Go(func() error {
			pkey, err := crypto.GenerateKey()
			if err != nil {
				return err
			}
			id := uuid.NewRandom()
			key := &keystore.Key{
				Id:         id,
				Address:    crypto.PubkeyToAddress(pkey.PublicKey),
				PrivateKey: pkey,
			}
			keyjson, err := keystore.EncryptKey(key, "", keystore.LightScryptN, keystore.LightScryptP)
			if err != nil {
				return err
			}
			fpath := filepath.Join(k.dir, name)
			if err := ioutil.WriteFile(fpath, keyjson, 0600); err != nil {
				return err
			}
			k.log.Debugf("generated key with address 0x%x. stored at %s", key.Address, fpath)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case generated <- key.PrivateKey:
			default:
			}
			return nil
		})
	}
	return group.Wait()
}

type faucet struct {
	log    *logrus.Entry
	erc    *erc20.ERC20
	client *ethclient.Client
	owner  bind.TransactOpts
}

func (f faucet) validate(ctx context.Context, gn, tn *big.Int, n int) error {
	totalGas := new(big.Int).Mul(gn, big.NewInt(int64(n)))
	totalTokens := new(big.Int).Mul(tn, big.NewInt(int64(n)))
	balance, err := f.client.BalanceAt(ctx, f.owner.From, nil)
	if err != nil {
		return err
	}
	if balance.Cmp(totalGas) <= 0 {
		return fmt.Errorf("total balance is %v, less than requested %v", balance, totalGas)
	}
	tbalance, err := f.erc.BalanceOf(&bind.CallOpts{Context: ctx}, f.owner.From)
	if err != nil {
		return err
	}
	if tbalance.Cmp(totalTokens) <= 0 {
		return fmt.Errorf("token balance is %v, less than requested %v", tbalance, totalTokens)
	}
	return nil
}

func (f faucet) fund(ctx context.Context, keys <-chan *ecdsa.PrivateKey, gn, tn *big.Int) error {
	nonce, err := f.client.PendingNonceAt(ctx, f.owner.From)
	if err != nil {
		return err
	}
	price, err := f.client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	group, ctx := errgroup.WithContext(ctx)
	for key := range keys {
		to := crypto.PubkeyToAddress(key.PublicKey)
		tx := types.NewTransaction(nonce, to, gn, 40000, price, nil)
		tx, err := f.owner.Signer(f.owner.From, tx)
		if err != nil {
			return err
		}
		nonce++
		if err := f.client.SendTransaction(ctx, tx); err != nil {
			return err
		}
		f.log.Infof("sending %v wei to 0x%x", gn, to)
		group.Go(func() error {
			receipt, err := bind.WaitMined(ctx, f.client, tx)
			if err != nil {
				return err
			}
			if receipt.Status == types.ReceiptStatusFailed {
				return fmt.Errorf("gas transfer 0x%x failed", tx.Hash())
			}
			f.log.Debugf("%v wei to 0x%x were sent. tx 0x%x", gn, to, tx.Hash())
			return nil
		})

		opts := &f.owner
		opts.Nonce = new(big.Int).SetUint64(nonce)
		opts.GasPrice = price
		opts.Context = ctx
		f.log.Infof("sending %v tokens to 0x%x", tn, to)
		ttx, err := f.erc.Transfer(opts, to, tn)
		if err != nil {
			return err
		}
		group.Go(func() error {
			receipt, err := bind.WaitMined(ctx, f.client, ttx)
			if err != nil {
				return err
			}
			if receipt.Status == types.ReceiptStatusFailed {
				return fmt.Errorf("token transfer 0x%x failed", ttx.Hash())
			}
			f.log.Debugf("%v tokens to 0x%x were sent. tx 0x%x", tn, to, ttx.Hash())
			return nil
		})
		nonce++
	}
	return group.Wait()
}
