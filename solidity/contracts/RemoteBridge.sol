pragma solidity 0.5.16;
pragma experimental ABIEncoderV2;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract RemoteBridge is Ownable {

  event TransferRegistered(bytes32 indexed hash, address indexed signer, uint64 nonce);

  struct Transfer {
    bytes32 hash;
    address signer;
    uint64 nonce;
    bool exist;
  }

  uint256 internal lastKnownBlock;
  mapping (bytes32 => Transfer) public transfers;

  constructor() public {}

  function register(bytes32 local, bytes32 remote, address signer, uint64 nonce) external onlyOwner {
    require(!transfers[local].exist, "transfer already registered");
    transfers[local] = Transfer(remote, signer, nonce, true);
    emit TransferRegistered(remote, signer, nonce);
  }

  function update(bytes32 local, bytes32 remote) external onlyOwner {
    require(transfers[local].exist, "transfer not registered");
    transfers[local].hash = remote;
    Transfer storage transfer = transfers[local];
    emit TransferRegistered(remote, transfer.signer, transfer.nonce);
  }

  function setLastBlock(uint256 lastBlock) external onlyOwner {
    lastKnownBlock = lastBlock;
  }

  function getLastBlock() external view returns(uint256) {
    return lastKnownBlock;
  }
}
