package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/nativebridge"
	"github.com/videocoin/go-bridge/service"
	"github.com/videocoin/go-bridge/service/tokentonative"
)

var (
	config = pflag.StringP("config", "c", "", "config for application")

	keypath = pflag.StringP("key", "k", "", "File with encrypted private key for funds")
	pwpath  = pflag.StringP("password", "p", "",
		"File with password for private key. If file is not provided application will try to decrypt with empty password")
)

type Config struct {
	LocalChainRPC   string
	ForeignChainRPC string

	ERC20Address  common.Address
	BridgeAddress common.Address

	LogLevel string

	BlockDelay        *big.Int
	ScanStep          *big.Int
	ScanPeriodSeconds int64

	Banks []common.Address
}

func decryptKey(log *logrus.Entry, path, pwpath string) *keystore.Key {
	f, err := os.Open(maybeSymlink(path))
	if err != nil {
		log.Fatalf("faield to open keyfile %s: %v", path, err)
	}
	data, err := ioutil.ReadAll(f)
	_ = f.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	password := ""
	if len(password) != 0 {
		f, err = os.Open(pwpath)
		if err != nil {
			log.Fatalf("can't open %s: %v", pwpath, err)
		}
		data, err := ioutil.ReadAll(f)
		_ = f.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		password = string(data)
	}
	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		log.Fatalf("failed to decrypt a key: %v", err)
	}
	return key
}

func maybeSymlink(path string) string {
	path = strings.TrimSpace(path)
	sympath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return path
	}
	return sympath
}

func parseConfig() Config {
	conf := Config{}

	f, err := os.Open(maybeSymlink(*config))
	if err != nil {
		fmt.Printf("can't open config at %s. err %v\n", *config, err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("can't read config at %s. err %v\n", *config, err)
		os.Exit(1)
	}
	_ = f.Close()
	if err := json.Unmarshal(data, &conf); err != nil {
		fmt.Printf("can't unmarshal config %s data. err %v\n", *config, err)
		os.Exit(1)
	}
	return conf
}

func main() {
	pflag.Parse()

	conf := parseConfig()

	logger := logrus.New()
	lvl, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		fmt.Printf("failed to parse level %s: %v", conf.LogLevel, err)
		os.Exit(1)
	}

	logger.SetLevel(lvl)
	log := logrus.NewEntry(logger)

	key := decryptKey(log, *keypath, *pwpath)

	lclient, err := ethclient.Dial(conf.LocalChainRPC)
	if err != nil {
		log.Fatal(err.Error())
	}
	fclient, err := ethclient.Dial(conf.ForeignChainRPC)
	if err != nil {
		log.Fatal(err.Error())
	}
	erc, err := erc20.NewERC20(conf.ERC20Address, fclient)
	if err != nil {
		log.Fatal(err.Error())
	}
	bridge, err := nativebridge.NewNativeBridge(conf.BridgeAddress, lclient)
	if err != nil {
		log.Fatal(err.Error())
	}
	opts := bind.NewKeyedTransactor(key.PrivateKey)

	engine := tokentonative.NewTransferEngine(log, lclient, *opts, bridge)
	svc := service.NewService(log, fclient,
		engine,
		engine,
		tokentonative.NewERC20Access(log, erc),
		service.StaticSource(conf.Banks),
		conf.BlockDelay, conf.ScanStep,
	)

	ctx, cancel := context.WithCancel(context.Background())
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	go func() {
		<-sigint
		cancel()
	}()
	period := time.Duration(conf.ScanPeriodSeconds) * time.Second
	log.Infof("bridge is started. scan period %v", period)
	_ = service.PollForever(ctx, period, 10*time.Minute, func(ctx context.Context) { _ = svc.Run(ctx) })
	log.Infof("bridge stopped")
}
