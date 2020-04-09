package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/videocoin/go-bridge/client"
)

// worker is one of the prepared key for testing workers, prefunded with tokens on goerli network.
const worker = `{"address":"7bd5ee00a271dbd99f1897768e02e04110a111d6","crypto":{"cipher":"aes-128-ctr","ciphertext":"3b5cd9da1f959d96ccc2426cdfa949ec4e96c360b9b4944ab0cd556016e981ad","cipherparams":{"iv":"59109f4099ec653ca5b49016c8a193a5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bd583b0242a8ac1f059f2ddf470e22cfd80d7ff0eae2daa865cdaf287ad717e8"},"mac":"f6b287123c4f222772ea08ec4e5d317d75b08b5c604fa98ca1c0c598dd5f3d14"},"id":"0105a013-e01e-4442-bc05-cb9b2ab31ef4","version":3}`

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// this is not a regular testable example. this example is using development data that may
	// get outdated. definitely should not be executed on CI.
	key, err := keystore.DecryptKey([]byte(worker), "")

	goerli, err := ethclient.Dial("https://goerli.infura.io/v3/977d81ca036540f48c4ab19f1927dcd0")
	must(err)
	vid, err := ethclient.Dial("https://dev1:D6msEL93LJT5RaPk@rpc.dev.kili.videocoin.network")
	must(err)
	config := client.Config{
		ProxyAddress:         common.HexToAddress("0x68896bAEcc5186284d5985903e86158F8e803Ca9"),
		ERC20Address:         common.HexToAddress("0x7bc345BE17AF288a0CFcFF8E673714635C847aa0"),
		LocalBridgeAddress:   common.HexToAddress("0xfD8c66B99919F291AefE756DA506446FB227f17D"),
		ForeignBridgeAddress: common.HexToAddress("0x06C8031ACa8B5d91Ae52F297511Ca4809aB55e29"),
	}
	client, err := client.Dial(vid, goerli, config)
	must(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	tokenbank := common.HexToAddress("0xf90d1852a344a67ca8dabfa1d307b483d89856c4")

	info, err := client.WaitDeposit(ctx, key.PrivateKey, tokenbank, big.NewInt(77))
	must(err)
	fmt.Printf("Deposit:\nAddress 0x%x\nVID tx hash 0x%x\nETH tx hash 0x%x\n",
		key.Address, info.LocalTxHash, info.ForeignTxHash)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	nativebank := common.HexToAddress("0xa2A92CeB62B447A6223fB4f3D2833eE9a40ED9B7")
	info, err = client.WaitWithdraw(ctx, key.PrivateKey, nativebank, big.NewInt(77))
	must(err)
	fmt.Printf("Withdraw:\nAddress 0x%x\nVID tx hash 0x%x\nETH tx hash 0x%x\n",
		key.Address, info.LocalTxHash, info.ForeignTxHash)
}
