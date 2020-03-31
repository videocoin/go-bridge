Bridge from ERC20 tokens to native coin
=======================================

ERC20 token is deployed on PoW "slow" blockchain with probabilistic finality.

Bridge performs its work in 3 actions:
1.
1.1. Fetch last known processed block from NativeBridge contract.
1.2. Fetch head of the chain with ERC20 contract
2. Fetch all transfers between last known block and head block using configurable step.
3. Execute transfers on local chain using NativeBridge contract.

NativeBridge contract maintains a map of completed transfers with a key of the original
transaction where ERC20 event was found. This guarantees that we will never execute same
transfer multiple times

Because bridged chain provides only probabilistic finality,  as a workaround for possible
reverted transactions, due to reorg, bridge executes transactions that were mined only N blocks ago.
Technically it is accomplished by substracting N blocks from head fetched in step 1.