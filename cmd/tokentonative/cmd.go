package tokentonative

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/videocoin/go-bridge/cmd/common"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/nativebridge"
	"github.com/videocoin/go-bridge/service"
	"github.com/videocoin/go-bridge/service/tokentonative"
)

type config struct {
	LocalChainRPC   string
	ForeignChainRPC string

	ERC20Address  ethcommon.Address
	BridgeAddress ethcommon.Address

	LogLevel string

	BlockDelay        *big.Int
	ScanStep          *big.Int
	ScanPeriodSeconds int64

	Banks []ethcommon.Address
}

func Command() *cobra.Command {
	var (
		confpath, keypath, pwpath string
	)
	cmd := &cobra.Command{
		Use:   "token2native",
		Short: "Bridge from erc20 token to native coin.",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config{}
			common.MustParseConfig(&conf, confpath)

			logger := logrus.New()
			lvl, err := logrus.ParseLevel(conf.LogLevel)
			if err != nil {
				fmt.Printf("failed to parse level %s: %v", conf.LogLevel, err)
				os.Exit(1)
			}

			logger.SetLevel(lvl)
			log := logrus.NewEntry(logger)

			key := common.MustDecryptKey(log, keypath, pwpath)

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
			_ = service.PollForever(ctx, period, 10*time.Minute, func(ctx context.Context) {
				err := svc.Run(ctx)
				if err != nil {
					log.Debugf("poll failed with %v", err)
				}
			})
			log.Infof("bridge stopped")
		},
	}
	cmd.Flags().StringVarP(&confpath, "config", "c", "", "config for application")
	cmd.Flags().StringVarP(&keypath, "key", "k", "", "File with encrypted private key for funds")
	cmd.Flags().StringVarP(&pwpath, "password", "p", "",
		"File with password for private key. If file is not provided application will use empty password")
	return cmd
}
