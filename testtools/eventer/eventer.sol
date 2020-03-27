pragma solidity ^0.5.16;

contract Eventer {

    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor() public {}

    function transfer(address recipient, uint256 amount) external returns (bool) {
      emit Transfer(msg.sender, recipient, amount);
      return true;
    }
}
