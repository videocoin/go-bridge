package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/videocoin/go-bridge/client"
	"github.com/videocoin/oauth2/google"
	"golang.org/x/oauth2"
)

// worker is one of the prepared key for testing workers, prefunded with tokens on goerli network.
const worker = `{"address":"8bc56ed72f5159d6ee385b8fdf0d395279e4f7e6","crypto":{"cipher":"aes-128-ctr","ciphertext":"e3b4f41becf47584ce037376b0eb4c32fb7f268f5d1c3e6a2733d4785c77a711","cipherparams":{"iv":"79557a50a0fb46adcf503f46a682085d"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"a10ec781526d51a93f0a72e8114f3eb97a26cf622ec73f7cd060b57e1a112603"},"mac":"4b805279d58ed1191ecfb9ba65c4bc8dc59b17d92379f2bf78ef941f27346fa4"},"id":"ff2832a5-b98e-4a69-b3bd-661f86062554","version":3}`

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func Dial(ctx context.Context, url, clientID, tokenfile string) (*ethclient.Client, error) {
	ts, err := google.IAPTokenSource(ctx, clientID, tokenfile)
	if err != nil {
		return nil, err
	}
	r, err := rpc.DialHTTPWithClient(url, oauth2.NewClient(ctx, ts))
	if err != nil {
		return nil, err
	}

	client := ethclient.NewClient(r)
	return client, nil
}

func main() {
	// this is not a regular testable example. this example is using development data that may
	// get outdated. definitely should not be executed on CI.
	key, err := keystore.DecryptKey([]byte(worker), "secret")

	goerli, err := ethclient.Dial("https://goerli.infura.io/v3/23b942b35e224604ae314b9ec3f1277c")
	must(err)
	vid, err := Dial(context.TODO(), "https://symphony.dev.videocoin.net/",
		"47928468404-hfuqhrb6lhtv9sem30rkjc1djcrlpt4v.apps.googleusercontent.com",
		"/home/dd/Downloads/videocoin-experiments-e8b4d2d00010.json",
	)
	must(err)
	config := client.Config{
		ProxyAddress:         common.HexToAddress("0xc16De466447e348b6Cd1B678d604990e6DB3057C"),
		ERC20Address:         common.HexToAddress("0x7bc345BE17AF288a0CFcFF8E673714635C847aa0"),
		LocalBridgeAddress:   common.HexToAddress("0xb067b9A2eb0bd087F859F836e0AC23E0691Ca62e"),
		ForeignBridgeAddress: common.HexToAddress("0x3CC38A35E3F93B7C57F44330c9584A48ef98E239"),
	}
	client, err := client.Dial(vid, goerli, config)
	must(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	tokenbank := common.HexToAddress("0x4d80ad6305b893a329039765134ddd436a87ff08")

	info, err := client.WaitDeposit(ctx, key.PrivateKey, tokenbank, big.NewInt(77))
	must(err)
	fmt.Printf("Deposit:\nAddress 0x%x\nVID tx hash 0x%x\nETH tx hash 0x%x\n",
		key.Address, info.LocalTxHash, info.ForeignTxHash)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	nativebank := common.HexToAddress("0xb8f52379ff40fe8ca57dc60ff24cea17bce043aa")
	info, err = client.WaitWithdraw(ctx, key.PrivateKey, nativebank, big.NewInt(77))
	must(err)
	fmt.Printf("Withdraw:\nAddress 0x%x\nVID tx hash 0x%x\nETH tx hash 0x%x\n",
		key.Address, info.LocalTxHash, info.ForeignTxHash)

}
