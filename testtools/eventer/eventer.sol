pragma solidity ^0.5.16;

contract Eventer {

    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor() public {}

    function transfer(address from, address to, uint256 value) external {
        emit Transfer(from, to, value);
    }
}
