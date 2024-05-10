// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

contract SkateApp {
    mapping(uint256 => bytes32) public tasks;
    uint256 private curTaskId;
    address owner;

    constructor(address _owner) {
        owner = _owner;
    }

    event TaskCreated(
        uint256 indexed taskId,
        bytes32 taskHash,
        string message,
        address signer,
        uint32 chain
    );
    struct Task {
        string message;
        uint32 chain;
    }

    function createMsg(string memory message, uint32 chain) public {
        Task memory task = Task({message: message, chain: chain});
        //Create onchain keccak hash as onchain proof
        tasks[curTaskId++] = keccak256(abi.encode(task));

        emit TaskCreated(
            curTaskId - 1,
            tasks[curTaskId - 1],
            message,
            msg.sender,
            chain
        );
    }

    function getProof(uint taskId) external view returns (bytes32) {
        return tasks[taskId];
    }
}
