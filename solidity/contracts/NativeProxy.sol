pragma solidity 0.5.16;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract NativeProxy {

  event TransferProxied(address indexed from, address indexed to, uint256 value);

  constructor() public {}

  function proxy(address payable to) external payable {
    (bool success,) = to.call.value(msg.value)("");
    require(success, "transfer proxy failed");
    emit TransferProxied(msg.sender, to, msg.value);
  }
}
