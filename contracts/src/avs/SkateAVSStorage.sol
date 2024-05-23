// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.20;

import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {ISkateAVS} from "./interfaces/ISkateAVS.sol";
import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IDelegationManager} from "./interfaces/IDelegationManager.sol";
import {BN254} from "./libraries/BN254.sol";

abstract contract SkateAVSStorage is ISkateAVS {
    StrategyParams[] internal _strategies;
    address[] internal _operators;
    mapping(address => bool) internal _isOperator;
    mapping(address => bool) internal _allowlist;
    uint32 internal _maxOperatorCount;
    bool public _allowlistEnabled;
    uint96 public _minOperatorStake;
    mapping(address => BN254.G2Point) internal _blsPubKey;
    bytes32 public constant PUBKEY_REGISTRATION_TYPEHASH = keccak256("BLSPubkey(address operator)");

    function strategies() external view override returns (StrategyParams[] memory) {
        return _strategies;
    }

    function isOperator(address operator) public view override returns (bool) {
        return _isOperator[operator];
    }

    function isInAllowlist(address operator) external view override returns (bool) {
        return _allowlist[operator];
    }

    function maxOperatorCount() external view override returns (uint32) {
        return _maxOperatorCount;
    }

    function allowlistEnabled() external view override returns (bool) {
        return _allowlistEnabled;
    }

    function minOperatorStake() external view override returns (uint96) {
        return _minOperatorStake;
    }
}
