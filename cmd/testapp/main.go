package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/videocoin/go-bridge/client"
	"github.com/videocoin/go-bridge/cmd/common"
	"github.com/videocoin/go-bridge/testapp"
)

var (
	confpath = pflag.StringP("config", "c", "", "Path to the config file")
)

type config struct {
	LocalURL, ForeignURL string

	ProxyAddress         ethcommon.Address
	ERC20Address         ethcommon.Address
	LocalBridgeAddress   ethcommon.Address
	ForeignBridgeAddress ethcommon.Address
	TokenBank            ethcommon.Address
	CoinBank             ethcommon.Address

	Period      time.Duration
	LogLevel    string
	DatabaseURL string

	KeysDir string
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func readKeys(log *logrus.Entry, dir string) (keys []*ecdsa.PrivateKey) {
	must(filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fpath := filepath.Join(dir, info.Name())
		key := common.MustDecryptKey(log, fpath, "")
		keys = append(keys, key.PrivateKey)
		return nil
	}))
	return keys
}

func main() {
	pflag.Parse()
	conf := config{}
	common.MustParseConfig(&conf, "test_app", *confpath)

	vid, err := ethclient.Dial(conf.LocalURL)
	must(err)
	eth, err := ethclient.Dial(conf.ForeignURL)
	must(err)
	client, err := client.Dial(vid, eth, client.Config{
		ProxyAddress:         conf.ProxyAddress,
		ERC20Address:         conf.ERC20Address,
		LocalBridgeAddress:   conf.LocalBridgeAddress,
		ForeignBridgeAddress: conf.ForeignBridgeAddress,
	})
	must(err)

	db, err := testapp.Open(conf.DatabaseURL)
	must(err)

	logger := logrus.New()
	lvl, err := logrus.ParseLevel(conf.LogLevel)
	must(err)
	logger.SetLevel(lvl)
	log := logrus.NewEntry(logger)

	manager := testapp.NewManager(
		log,
		conf.Period,
		rand.New(rand.NewSource(time.Now().Unix())),
		db,
		readKeys(log, conf.KeysDir),
		client,
		conf.TokenBank,
		conf.CoinBank,
	)

	ctx, cancel := context.WithCancel(context.Background())
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	go func() {
		<-sigint
		cancel()
	}()
	if err := manager.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		must(err)
	}
}
