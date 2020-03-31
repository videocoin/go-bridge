Native coin to ERC20 token
==========================

Bridge monitors native coin transfers to one of the bank address that are proxied with NativeProxy.
There are two reasons to use NativeProxy instead of making regular transactions:
- filter out transfers that shouldn't be proxied, such as faucet transactions.
- provide efficient api to fetch transfers that should be proxied, e.g. using getLogs.
  The alternative is to scan blockchain for transfers that are sent to a particular address.

Once bridge observed TransferProxied event, it will try to execute this transfer on the network
with ERC20 token. Additionally to track state we are using RemoteBridge contract that is deployed
on local network (videocoin), this contract is required to prevent accidental transfer of the same
value twice/more.

The whole protocol works as follows:

1. Once local transfer is observed - check if transfer exists in RemoteBridge contract
2. If transfers exists:
2.1. Check if transfer exists on remote blockchain, if it does - wait for transfer to be mined.
2.2. If it doesn't resent transfer with nonce that we recorded in RemoveBridge.
3. If transfer doesn't exist
3.1. Pick pending nonce from blockchain with ERC20 token.
3.2. Register pending transfer in RemoteBridge
3.3. Submit transfer on chain.
4. Wait for trasnfer to get mined.
