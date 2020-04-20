package nativetotoken

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/videocoin/go-bridge/cmd/common"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/nativeproxy"
	"github.com/videocoin/go-bridge/remotebridge"
	"github.com/videocoin/go-bridge/service"
	"github.com/videocoin/go-bridge/service/blocks"
	"github.com/videocoin/go-bridge/service/nativetotoken"
)

type config struct {
	DataDir         string
	LocalChainRPC   string
	ForeignChainRPC string

	ERC20Address       ethcommon.Address
	NativeProxyAddress ethcommon.Address
	BridgeAddress      ethcommon.Address

	LogLevel string

	ScanStep          *big.Int
	ScanPeriodSeconds int64

	EnablePrometheus   bool
	PrometheusListener string

	Banks []ethcommon.Address
}

func Command() *cobra.Command {
	var (
		confpath                    string
		tokenKeypath, tokenPwpath   string
		bridgeKeypath, bridgePwpath string
	)
	cmd := &cobra.Command{
		Use:   "native2token",
		Short: "Bridge from native coin to erc20 token.",
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

			key := common.MustDecryptKey(log, tokenKeypath, tokenPwpath)
			tokenopts := bind.NewKeyedTransactor(key.PrivateKey)

			key = common.MustDecryptKey(log, bridgeKeypath, bridgePwpath)
			bridgeopts := bind.NewKeyedTransactor(key.PrivateKey)

			lclient, err := ethclient.Dial(conf.LocalChainRPC)
			if err != nil {
				log.Fatal(err.Error())
			}
			fclient, err := ethclient.Dial(conf.ForeignChainRPC)
			if err != nil {
				log.Fatal(err.Error())
			}

			proxy, err := nativeproxy.NewNativeProxy(conf.NativeProxyAddress, lclient)
			if err != nil {
				log.Fatal(err.Error())
			}
			bridge, err := remotebridge.NewRemoteBridge(conf.BridgeAddress, lclient)
			if err != nil {
				log.Fatal(err.Error())
			}
			erc, err := erc20.NewERC20(conf.ERC20Address, fclient)
			if err != nil {
				log.Fatalf(err.Error())
			}

			blocks, err := blocks.NewWriterBlockResource(log, filepath.Join(conf.DataDir, "last_block_n2t"))
			if err != nil {
				log.Fatalf(err.Error())
			}

			engine := nativetotoken.NewTransferEngine(log,
				lclient, fclient,
				*bridgeopts, *tokenopts,
				bridge, erc)
			svc := service.NewService(log, lclient,
				engine,
				blocks,
				nativetotoken.NewNativeTransferAccess(log, proxy),
				service.StaticSource(conf.Banks),
				big.NewInt(0), conf.ScanStep,
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

			var (
				errors = make(chan error, 2)
				wg     sync.WaitGroup
			)

			if conf.EnablePrometheus {
				wg.Add(1)
				go func() {
					log.Infof("started prometheus on %s", conf.PrometheusListener)
					errors <- common.BootstrapPrometheus(ctx, conf.PrometheusListener)
					wg.Done()
				}()
			}

			wg.Add(1)
			go func() {
				errors <- service.PollForever(ctx, period, 10*time.Minute, func(ctx context.Context) {
					err := svc.Run(ctx)
					if err != nil {
						log.Debugf("poll failed with %v", err)
					}
				})
				wg.Done()
			}()
			go func() {
				wg.Wait()
				close(errors)
			}()
			for err := range errors {
				if err != nil {
					log.Fatal(err.Error())
				}
			}
			log.Infof("bridge stopped")
		},
	}
	cmd.Flags().StringVarP(&confpath, "config", "c", "", "config for application")
	cmd.Flags().StringVar(&tokenKeypath, "tokenkey", "",
		"File with encrypted private key with tokens on the remote side of the bridge")
	cmd.Flags().StringVar(&tokenPwpath, "tokenpw", "",
		"File with password for token private key. If file is not provided application will use empty password")
	cmd.Flags().StringVar(&bridgeKeypath, "bridgekey", "",
		"File with encrypted private key with gas to cover state modifications in the RemoteBridge contract.")
	cmd.Flags().StringVar(&bridgePwpath, "bridgepw", "",
		"File with password for bridge private key. If file is not provided application will use empty password")
	return cmd
}
