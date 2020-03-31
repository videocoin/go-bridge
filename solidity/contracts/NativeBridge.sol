pragma solidity 0.5.16;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

/// @title VID homd bridge contract: handles VID transfers on alpha network
contract NativeBridge is Ownable {

  event TransferBridged(address indexed from, address indexed to, bytes32 indexed txHash, uint256 value);

  uint256 internal lastKnownBlock;
  mapping (bytes32 => bool) public transfers;

  constructor() public {}

  /// @notice Method for executing VID transfers coming from ERC20 token.
  function transfer(address payable to, bytes32 txHash) external payable {
    require(!transfers[txHash]);
    require(to != address(0));

    (bool success,) = to.call.value(msg.value)("");
    require(success, "transfer failed");
    emit TransferBridged(msg.sender, to, txHash, msg.value);

    transfers[txHash] = true;
  }

  function setLastBlock(uint256 lastBlock) public onlyOwner {
    lastKnownBlock = lastBlock;
  }

  function getLastBlock() public view returns(uint256) {
    return lastKnownBlock;
  }
}
