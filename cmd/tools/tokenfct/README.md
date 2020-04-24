Utility to generate required number of accounts and faucet them withs tokens and gas
===

#### To install:

```
go get github.com/videocoin/go-bridge/cmd/tools/tokenfct
```

Or clone repository and run:

```
make tokenfct
```

#### Configuration:

```json
{
    "TokenContract": "0x7bc345BE17AF288a0CFcFF8E673714635C847aa0",
    "LogLevel": "debug",
    "DataDir": "/tmp/t1",
    "URL": "https://goerli.infura.io/v3/9bd08b37cf1d465c8bf05a3ace5f9d4a"
}
```

Configuration that is reused across multiple runs must be stored in config file.
Most of the options are self-explanatory.

Command line options

```
Usage of ./build/tokenfct:
  -c, --config string    configuration for faucet helper (default "faucetcfg.json")
  -g, --gas int          one unit is 1e12 wei (default 10)
  -n, --number int       number of keys to generate and to faucet (if key exist it will just fauceted) (default 10)
      --owner string     file with account with private keys (default "owner.json")
      --ownerpw string   file with owner key password. if file is not provided - app will try to decrypt with empty pw
  -t, --tokens int       one unit is 1e18 wei (default 100)
pflag: help requeste
```

#### Example:

Following example will generate 3 keys and faucet them with default amount gas and tokens.
If keys were generated in previous run, command will faucet existing keys again.

```
./build/tokenfct -c _assets/faucetcfg.json --owner token-bank.json -n 3
```

Expected output:

```
DEBU[0001] generated key with address 0xa1abe8ea885c59b2b92a775f216a0fcce77a8eee. stored at /tmp/t1/worker-1.json
DEBU[0001] generated key with address 0x82714528172f6aff7022e86d9818d469e84857b8. stored at /tmp/t1/worker-0.json
DEBU[0001] generated key with address 0x969f184f726502a7965474b4c43f0dab334c3fe6. stored at /tmp/t1/worker-2.json
INFO[0002] sending 10000000000000 wei to 0xa1abe8ea885c59b2b92a775f216a0fcce77a8eee
INFO[0002] sending 100000000000000000000 tokens to 0xa1abe8ea885c59b2b92a775f216a0fcce77a8eee
INFO[0003] sending 10000000000000 wei to 0x82714528172f6aff7022e86d9818d469e84857b8
INFO[0003] sending 100000000000000000000 tokens to 0x82714528172f6aff7022e86d9818d469e84857b8
INFO[0003] sending 10000000000000 wei to 0x969f184f726502a7965474b4c43f0dab334c3fe6
INFO[0003] sending 100000000000000000000 tokens to 0x969f184f726502a7965474b4c43f0dab334c3fe6
DEBU[0015] 100000000000000000000 tokens to 0xa1abe8ea885c59b2b92a775f216a0fcce77a8eee were sent. tx 0x1f129224b9b4783b63374ad2e4763451a39a7e24ec165e4ca874b08bd8d527f9
DEBU[0020] 10000000000000 wei to 0xa1abe8ea885c59b2b92a775f216a0fcce77a8eee were sent. tx 0x1f129224b9b4783b63374ad2e4763451a39a7e24ec165e4ca874b08bd8d527f9
DEBU[0027] 10000000000000 wei to 0x82714528172f6aff7022e86d9818d469e84857b8 were sent. tx 0x434922e79f8b87d2ea224fc98e92f63d525c3da61e110645ac08357e1720b269
DEBU[0027] 100000000000000000000 tokens to 0x969f184f726502a7965474b4c43f0dab334c3fe6 were sent. tx 0x69a3dfb325754719f9afc884c3e2c5d6c3fed1cc8452142296f3492a6d4a2e83
DEBU[0027] 100000000000000000000 tokens to 0x82714528172f6aff7022e86d9818d469e84857b8 were sent. tx 0x434922e79f8b87d2ea224fc98e92f63d525c3da61e110645ac08357e1720b269
DEBU[0027] 10000000000000 wei to 0x969f184f726502a7965474b4c43f0dab334c3fe6 were sent. tx 0x69a3dfb325754719f9afc884c3e2c5d6c3fed1cc8452142296f3492a6d4a2e83
Keys generated and funded with gas and tokens.
```

#### Tokens owner

If owner runs out of gas, please faucet it on:

- https://faucet.goerli.mudit.blog/
- https://goerli-faucet.slock.it/

```json
{"address":"218431c2e6b7a8d6af5ef18f3ad63b23c614889f","crypto":{"cipher":"aes-128-ctr","ciphertext":"4b272ba7aee44f6725d7abfc339f5120760ba8c7faf7f94208bfab7ea290a4a3","cipherparams":{"iv":"5d7808c00586f603994faf8b56924750"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8651643b518ecfbb510870137686ff3a5b6971c4b49ea310b0ec1c5efb3b764e"},"mac":"5817ec372694abaeef8cd7891e26281561cf68d713ef687a90eee7f7c6a52bce"},"id":"e8f3d517-4f93-439c-a98e-f2315fa1de19","version":3}
```