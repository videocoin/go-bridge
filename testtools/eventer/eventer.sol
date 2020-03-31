pragma solidity ^0.5.16;

contract Eventer {

    event Transfer(address indexed from, address indexed to, uint256 value);

    mapping (address => uint256) private balances;

    constructor() public {}

    function balanceOf(address account) external view returns (uint256) {
      return balances[account];
    }

    function transfer(address recipient, uint256 amount) external returns (bool) {
      emit Transfer(msg.sender, recipient, amount);
      balances[recipient] += amount;
      return true;
    }
}
