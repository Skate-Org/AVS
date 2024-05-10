// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

contract SkateGateway {
    address owner;
    mapping(uint256 => MessageData) public messages;
    address public relayer;

    struct MessageData {
        string message;
        address signer;
    }

    constructor(address _owner) {
        owner = _owner;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyRelayer() {
        require(relayer == msg.sender);
        _;
    }

    function registerRelayer(address newRelayer) public onlyOwner {
        relayer = newRelayer;
    }

    function postMsg(
        uint taskId,
        string memory message,
        address signer
    ) public onlyRelayer {
        // TODO: Verify that message has appeared in Skate AVS
        messages[taskId] = MessageData(message, signer);
    }

    function getMsg(uint taskId) public view returns (string memory) {
        return messages[taskId].message;
    }
}
