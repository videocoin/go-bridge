Supported deployments
=====================

## Goerli bridge

Please create your own URL for Goerli test network. Infura provides free account
with a 100 000 requests/day limit.

### Bridge from ERC20 to native

#### ERC20

- Address: 0xB269eA0cc5270dF258B5e4c4Fdc7F35E50cb91Ae
  Etherscan: https://goerli.etherscan.io/address/0xB269eA0cc5270dF258B5e4c4Fdc7F35E50cb91Ae
- Bank: 0xf90d1852a344a67ca8dabfa1d307b483d89856c4
  Etherscan: https://goerli.etherscan.io/address/0xf90d1852a344a67ca8dabfa1d307b483d89856c4

Bank is funded with 1 000 000 tokens. If it will run out bridge won't be making progress
until bank is refunded.

To deposit on videocoin chain submit ERC20 transfer to bank address.
Bindings for ERC20 are available in this repository [erc 20 bindings](./erc20).

#### NativeBridge

- Address: 0xfD8c66B99919F291AefE756DA506446FB227f17D

NativeBridge is managed by one of the services in this repo.
On the user side, you can verify if erc20 transfer was completed on videocoin chain by
using call to the [executed transfers](./nativebridge/nativebridge.go#L273). Transaction hash
that was used to create ERC20 transfer should be used as an input.

You can check this [test](./service/tokentonative/service_test.go#L44-L84) for an example how it works.

### Bridge from native to erc20

#### NativeProxy

- Address: 0x68896bAEcc5186284d5985903e86158F8e803Ca9
- Bank address: 0xa2A92CeB62B447A6223fB4f3D2833eE9a40ED9B7

Is a proxy for native transfers. It records event for each transfer that simplifies blockchain monitoring
and prevents executing unneccassary transfers (e.g transfer from faucet to videocoin bank).

Bindings are available at (./nativeproxy).

You can check this [test](./https://github.com/videocoin/go-bridge/blob/master/service/nativetotoken/resources_test.go#L40-L60) to see how to work with NativeProxy.

#### RemoteBridge

- Address: 0x06C8031ACa8B5d91Ae52F297511Ca4809aB55e29

Tracks state on videocoin chain related to ERC20 transfers on ethereum blockchain.

Bindings are available at [bridge](./remotebridge).

### Prefunded keys

Each key is funded with 1 000 000 tokens. For gas please use goerli faucets:

- https://goerli-faucet.slock.it/
- https://faucet.goerli.mudit.blog/

Key is encrypted with empty passphrase.

```json
{"address":"7bd5ee00a271dbd99f1897768e02e04110a111d6","crypto":{"cipher":"aes-128-ctr","ciphertext":"3b5cd9da1f959d96ccc2426cdfa949ec4e96c360b9b4944ab0cd556016e981ad","cipherparams":{"iv":"59109f4099ec653ca5b49016c8a193a5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bd583b0242a8ac1f059f2ddf470e22cfd80d7ff0eae2daa865cdaf287ad717e8"},"mac":"f6b287123c4f222772ea08ec4e5d317d75b08b5c604fa98ca1c0c598dd5f3d14"},"id":"0105a013-e01e-4442-bc05-cb9b2ab31ef4","version":3}

{"address":"279f6c49a3d34abe7b15e85aaae842e42b9c0092","crypto":{"cipher":"aes-128-ctr","ciphertext":"de709cb9f77a4951d42c4df96e0324aa238a4186c4a070fee8aed42a4435aa17","cipherparams":{"iv":"b35f332d2c702541a2135353c7ad286e"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"6f96fc1f195a60e94a69531e8eca23bcc2fd47db39354c06bcde77ba77d40c94"},"mac":"8a57f57e147bfba3ba4ae853f4d9695f8497ba18097959f8d4d29ca48aec0415"},"id":"80971801-2176-4ef6-abee-cf82c24cf564","version":3}

{"address":"c00872c549d6727c07e4ddc6c17cd2734972208e","crypto":{"cipher":"aes-128-ctr","ciphertext":"01cc35f295d4050251f351cd04572ce33beeff3bbad12fe4453a50052dc91a2b","cipherparams":{"iv":"1b59cf11f6eccbb89818f19051db66e5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"6ca614c7d3f4dbec410ccaa5baf50615298a88e5609d2deee4b3810c5f0887c7"},"mac":"1df9daf0c373d8fcfba9b48f3912b0967cf9f49e788746480293ff70b278ad65"},"id":"dfc380ef-5b6b-47e6-9a43-887cbb0b8f30","version":3}
```