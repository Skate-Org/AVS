// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.20;

library Errors {
    error OnlyOperatorAllowedToCall();
    error OperatorNotAllowed();
    error AlreadyAnOperator();
    error MaxOperatorCountReached();
    error MinOperatorStakeNotSatisfied();
    error ZeroOperatorAddress();
    error OperatorAlreadyInAllowlist();
    error OperatorNotInAllowlist();
    error NotAnOperator();
    error AllowlistAlreadyEnabled();
    error AllowlistAlreadyDisabled();
    error ZeroStrategyAddress();
    error StrategyAlreadyAdded();
    error DuplicateSignature();
    error SignaturesAreNotOrdered();
    error QuorumNotReached();
    error InvalidBLSSignature();
}
