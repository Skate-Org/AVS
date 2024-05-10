// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.20;

import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

interface ISkateAVS {
    struct StrategyParams {
        IStrategy strategy;
        uint96 multiplier;
    }

    event StrategiesSet(StrategyParams[] strategies);
    event OperatorAdded(address indexed operator);
    event OperatorRemoved(address indexed operator);
    event MaxOperatorCountSet(uint32 maxOperatorCount);
    event OperatorAllowed(address indexed operator);
    event OperatorDisallowed(address indexed operator);
    event MinOperatorStakeSet(uint96 minOperatorStake);
    event AllowlistEnabled();
    event AllowlistDisabled();
    event DataSubmitted(uint256 taskId, bytes messageData);

    struct Operator {
        address addr;
        uint96 delegated;
        uint96 staked;
    }

    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;
    function updateAVSMetadataURI(string memory metadataURI) external;
    function setStrategies(StrategyParams[] calldata strategies_) external;
    function setMinOperatorStake(uint96 stake) external;
    function setMaxOperatorCount(uint32 count) external;
    function addToAllowlist(address operator) external;
    function removeFromAllowlist(address operator) external;
    function enableAllowlist() external;
    function disableAllowlist() external;
    function deregisterOperatorFromAVS(address operator) external;
    function pause() external;
    function unpause() external;

    struct SignatureTuple {
        address operator;
        bytes signature;
    }

    function submitData(uint256 taskId, bytes calldata messageData, SignatureTuple[] calldata signatureTuples)
        external;

    function batchSubmitData(
        uint256[] calldata taskIds,
        bytes[] calldata messageDatas,
        SignatureTuple[][] calldata signaturesTuples
    ) external;

    function isInAllowlist(address operator) external view returns (bool);
    function avsDirectory() external view returns (address);
    function operators() external view returns (Operator[] memory);
    function isOperator(address operator) external view returns (bool);
    function strategies() external view returns (StrategyParams[] memory);
    function getRestakeableStrategies() external view returns (address[] memory);
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory);
    function canRegister(address operator) external view returns (bool);
    function maxOperatorCount() external view returns (uint32);
    function allowlistEnabled() external view returns (bool);
    function minOperatorStake() external view returns (uint96);

    //    function operatorPubkeys(address operator) external view returns (bytes32);
}
