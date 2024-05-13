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
        uint32 chainType,
        uint32 chainId
    );

    struct Task {
        string message;
        uint32 chainId;
        uint32 chainType;
    }

    function createMsg(
        string memory message,
        uint32 chainType,
        uint32 chainId
    ) public {
        Task memory task = Task({
            message: message,
            chainType: chainType,
            chainId: chainId
        });

        // task proof
        tasks[curTaskId++] = keccak256(abi.encode(task));

        emit TaskCreated(
            curTaskId - 1,
            tasks[curTaskId - 1],
            message,
            msg.sender,
            chainType,
            chainId
        );
    }

    function getProof(uint taskId) external view returns (bytes32) {
        return tasks[taskId];
    }
}
