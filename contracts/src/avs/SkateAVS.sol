// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import {IDelegationManager} from "./interfaces/IDelegationManager.sol";

import {SkateAVSStorage} from "./SkateAVSStorage.sol";
import {Errors} from "./Errors.sol";

/**
 * @notice StakeAVS contract is an implementation of AVS by Skate chain. It allows operator to opt-in to Skate AVS.
 * The operators can validate and sign the data that is later submitted to this contract upon reaching of quorum of
 * 2/3 of operators.
 */
contract SkateAVS is Initializable, UUPSUpgradeable, OwnableUpgradeable, PausableUpgradeable, SkateAVSStorage {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    uint256 public constant WEIGHTING_DIVISOR = 1e18;
    IAVSDirectory internal immutable _avsDirectory;
    IDelegationManager internal immutable _delegationManager;

    /**
     * @notice runs on contract deployment and initializes avs directory and delegation manager instances.
     * @param avsDirectory_ instance of avs directory.
     * @param delegationManager_ instance of delegation manager.
     */
    constructor(IAVSDirectory avsDirectory_, IDelegationManager delegationManager_) {
        _avsDirectory = avsDirectory_;
        _delegationManager = delegationManager_;
        _disableInitializers();
    }

    /**
     * @notice called on contract initialization though proxy.
     * @param owner_ address of owner.
     * @param strategies_ list of strategies supported by avs.
     * @param metadataURI_ url to metadata.
     * @param allowlistEnabled_ if operator allowlist is enabled or not.
     */
    function initialize(
        address owner_,
        StrategyParams[] calldata strategies_,
        string calldata metadataURI_,
        bool allowlistEnabled_
    ) external initializer {
        if (allowlistEnabled_) {
            _enableAllowlist();
        }

        if (bytes(metadataURI_).length != 0) {
            _avsDirectory.updateAVSMetadataURI(metadataURI_);
        }
        _setMaxOperatorCount(5);
        _setMinOperatorStake(1);
        _setStrategies(strategies_);
        _transferOwnership(owner_);
    }

    /**
     * @notice registers operator on AVS.
     * @param operator address of operator
     * @param operatorSignature operator signature to verify the validity of operator
     * requirements
     * - can only be called when AVS is not paused.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override whenNotPaused {
        if (operator != msg.sender) revert Errors.OnlyOperatorAllowedToCall();
        if (_allowlistEnabled && !_allowlist[operator]) revert Errors.OperatorNotAllowed();
        if (isOperator(operator)) revert Errors.AlreadyAnOperator();
        if (_operators.length == _maxOperatorCount) revert Errors.MaxOperatorCountReached();
        if (_operatorTotalDelegations(operator) < _minOperatorStake) revert Errors.MinOperatorStakeNotSatisfied();

        _isOperator[operator] = true;
        _operators.push(operator);
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);

        emit OperatorAdded(operator);
    }

    /**
     * @notice Returns the EigenLayer AVSDirectory contract.
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function avsDirectory() external view override returns (address) {
        return address(_avsDirectory);
    }

    /**
     * @notice Returns the currrent list of operator registered as OmniAVS.
     *         Operator.addr        = The operator's ethereum address
     *         Operator.staked      = The total amount staked by the operator, not including delegations
     *         Operator.delegated   = The total amount delegated, not including operator stake
     */
    function operators() public view override returns (Operator[] memory operators_) {
        operators_ = new Operator[](_operators.length);

        for (uint256 i = 0; i < operators_.length; i++) {
            address operator = _operators[i];
            uint96 total = _operatorTotalDelegations(operator);
            uint96 staked = _getSelfDelegations(operator);
            operators_[i] = Operator(operator, total > staked ? total - staked : 0, staked);
        }
    }

    /**
     * @notice Returns the list of strategies that the AVS supports for restaking.
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function getRestakeableStrategies() external view override returns (address[] memory) {
        return _getRestakeableStrategies();
    }

    /**
     * @notice Returns the list of strategies that the operator has potentially restaked on the AVS
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     *
     *      This function is intended to be called off-chain
     *
     *      No guarantee is made on whether the operator has shares for a strategy. The off-chain
     *      service should do that validation separately. This matches the behavior defined in
     *      eigenlayer-middleware's ServiceManagerBase.
     *
     * @param operator The address of the operator to get restaked strategies for
     */
    function getOperatorRestakedStrategies(address operator) external view override returns (address[] memory) {
        if (!isOperator(operator)) return new address[](0);
        return _getRestakeableStrategies();
    }

    /**
     * @notice Check if an operator can register to the AVS.
     *         Returns true, with no reason, if the operator can register to the AVS.
     *         Returns false, with a reason, if the operator cannot register to the AVS.
     * @dev This function is intended to be called off-chain.
     * @param operator The operator to check
     * @return canRegister True if the operator can register, false otherwise
     */
    function canRegister(address operator) external view override returns (bool) {
        if (!_delegationManager.isOperator(operator)) revert Errors.NotAnOperator();
        if (_allowlistEnabled && !_allowlist[operator]) revert Errors.OperatorNotAllowed();
        if (isOperator(operator)) revert Errors.AlreadyAnOperator();
        if (_operators.length >= _maxOperatorCount) revert Errors.MaxOperatorCountReached();
        if (_operatorTotalDelegations(operator) < _minOperatorStake) revert Errors.MinOperatorStakeNotSatisfied();
        return true;
    }

    /**
     * @notice Sets AVS metadata URI with the AVSDirectory.
     * @param metadataURI metadata uri to set.
     */
    function updateAVSMetadataURI(string memory metadataURI) external override onlyOwner {
        _avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    /**
     * @notice sets the new stratgies.
     * @param strategies_ the list of new strategies to set
     */
    function setStrategies(StrategyParams[] calldata strategies_) external override onlyOwner {
        _setStrategies(strategies_);
    }

    /**
     * @notice Set the minimum operator stake.
     * @param stake The minimum operator stake, not including delegations
     */
    function setMinOperatorStake(uint96 stake) external override onlyOwner {
        _setMinOperatorStake(stake);
    }

    /**
     * @notice Set the maximum operator count.
     * @param count The maximum operator count
     */
    function setMaxOperatorCount(uint32 count) external override onlyOwner {
        _setMaxOperatorCount(count);
    }

    /**
     * @notice Add an operator to the allowlist.
     * @param operator The operator to add
     */
    function addToAllowlist(address operator) external override onlyOwner {
        if (operator == address(0x0)) revert Errors.ZeroOperatorAddress();
        if (_allowlist[operator]) revert Errors.OperatorAlreadyInAllowlist();
        _allowlist[operator] = true;
        emit OperatorAllowed(operator);
    }

    /**
     * @notice Remove an operator from the allowlist.
     * @param operator The operator to remove
     */
    function removeFromAllowlist(address operator) external override onlyOwner {
        if (!_allowlist[operator]) revert Errors.OperatorNotInAllowlist();
        _allowlist[operator] = false;
        emit OperatorDisallowed(operator);
    }

    /**
     * @notice Enable the allowlist.
     */
    function enableAllowlist() external override onlyOwner {
        _enableAllowlist();
    }

    /**
     * @notice Disable the allowlist.
     */
    function disableAllowlist() external override onlyOwner {
        _disableAllowlist();
    }

    /**
     * @notice Eject an operator from the AVS.
     * @param operator the address of operator to eject.
     */
    function deregisterOperatorFromAVS(address operator) external override onlyOwner {
        if (!isOperator(operator)) revert Errors.NotAnOperator();

        for (uint256 i = 0; i < _operators.length; i++) {
            if (_operators[i] == operator) {
                _operators[i] = _operators[_operators.length - 1];
                _operators.pop();
                break;
            }
        }
        _avsDirectory.deregisterOperatorFromAVS(operator);

        emit OperatorRemoved(operator);
    }

    /**
     * @notice Pause the contract.
     * @dev This pauses registerOperatorToAVS, deregisterOperatorFromAVS, and syncWithOmni.
     */
    function pause() external override onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause the contract.
     */
    function unpause() external override onlyOwner {
        _unpause();
    }

    /**
     * @notice submits data after verification from operators. It validates the passed data though signature
     * verification. The quorum of 2/3 of operators must be reached for the successful submission.
     * @param taskId The id of the task.
     * @param messageData the message data validated by the operators.
     * @param signatureTuples the list of operator signatures to be validated.
     */
    function submitData(uint256 taskId, bytes calldata messageData, SignatureTuple[] calldata signatureTuples)
        public
        override
    {
        bytes32 digest = keccak256(abi.encodePacked(taskId, messageData)).toEthSignedMessageHash();
        uint256 sigsVerified;
        bool quorumSuccessful;
        for (uint256 i = 0; i < signatureTuples.length; i++) {
            SignatureTuple memory sigTuple = signatureTuples[i];
            if (!isOperator(sigTuple.operator)) revert Errors.NotAnOperator();

            if (i > 0) {
                SignatureTuple memory prevSigTuple = signatureTuples[i - 1];
                if (sigTuple.operator == prevSigTuple.operator) revert Errors.DuplicateSignature();
                if (sigTuple.operator < prevSigTuple.operator) revert Errors.SignaturesAreNotOrdered();
            }

            if (ECDSA.recover(digest, sigTuple.signature) != sigTuple.operator) revert Errors.InvalidSignature();
            sigsVerified++;

            // 2/3 of operators must submit the data.
            if (sigsVerified * 10_000 >= operators().length * 6666) {
                quorumSuccessful = true;
                break;
            }
        }

        if (!quorumSuccessful) revert Errors.QuorumNotReached();
        emit DataSubmitted(taskId, messageData);
    }

    /**
     * @notice submits data in batch and verify it. Internally calls {submitData} function.
     * @param taskIds the list of task ids.
     * @param messageDatas the list of message datas.
     * @param signaturesTuples the list of signature tuples by operators.
     */
    function batchSubmitData(
        uint256[] calldata taskIds,
        bytes[] calldata messageDatas,
        SignatureTuple[][] calldata signaturesTuples
    ) external override {
        for (uint256 i = 0; i < taskIds.length; i++) {
            submitData(taskIds[i], messageDatas[i], signaturesTuples[i]);
        }
    }

    /**
     * @notice Set the minimum operator stake.
     * @param stake The minimum operator stake, not including delegations
     */
    function _setMinOperatorStake(uint96 stake) private {
        _minOperatorStake = stake;
        emit MinOperatorStakeSet(stake);
    }

    /**
     * @notice Set the maximum operator count.
     * @param count The maximum operator count
     */
    function _setMaxOperatorCount(uint32 count) private {
        _maxOperatorCount = count;
        emit MaxOperatorCountSet(count);
    }

    /**
     * @notice Enable the allowlist.
     */
    function _enableAllowlist() private {
        if (_allowlistEnabled) revert Errors.AllowlistAlreadyEnabled();
        _allowlistEnabled = true;
        emit AllowlistEnabled();
    }

    /**
     * @notice Disable the allowlist.
     */
    function _disableAllowlist() private {
        if (!_allowlistEnabled) revert Errors.AllowlistAlreadyDisabled();
        _allowlistEnabled = false;
        emit AllowlistDisabled();
    }

    /**
     * @notice sets the new strategies.
     * @param strategies_ the new strategies
     */
    function _setStrategies(StrategyParams[] calldata strategies_) private {
        delete _strategies;

        for (uint256 i = 0; i < strategies_.length; i++) {
            if (address(strategies_[i].strategy) == address(0x0)) revert Errors.ZeroStrategyAddress();

            // ensure no duplicates
            for (uint256 j = i + 1; j < strategies_.length; j++) {
                if (strategies_[i].strategy == strategies_[j].strategy) revert Errors.StrategyAlreadyAdded();
            }
            _strategies.push(strategies_[i]);
        }

        emit StrategiesSet(strategies_);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Internal views                              //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Returns the operator's self-delegations
     * @param operator The operator address
     */
    function _getSelfDelegations(address operator) internal view returns (uint96 staked) {
        (IStrategy[] memory strategiesByOperator, uint256[] memory shares) =
            _delegationManager.getDelegatableShares(operator);

        for (uint256 i = 0; i < strategiesByOperator.length; i++) {
            // find the strategy params for the strategy
            StrategyParams memory params;
            for (uint256 j = 0; j < _strategies.length; j++) {
                if (_strategies[j].strategy == strategiesByOperator[i]) {
                    params = _strategies[j];
                    break;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(params.strategy) == address(0)) continue;

            staked += uint96(shares[i] * params.multiplier / WEIGHTING_DIVISOR);
        }
    }

    /**
     * @notice Returns total delegations to the operator, including self delegations
     * @param operator The operator address
     */
    function _operatorTotalDelegations(address operator) internal view returns (uint96 delegation) {
        for (uint256 i = 0; i < _strategies.length; i++) {
            uint256 shares = _delegationManager.operatorShares(operator, _strategies[i].strategy);
            delegation += uint96(shares * _strategies[i].multiplier / WEIGHTING_DIVISOR);
        }
    }

    /**
     * @notice Returns the list of restakeable strategy addresses
     */
    function _getRestakeableStrategies() internal view returns (address[] memory strategies_) {
        strategies_ = new address[](_strategies.length);
        for (uint256 i = 0; i < _strategies.length; i++) {
            strategies_[i] = address(_strategies[i].strategy);
        }
    }

    /**
     * @notice called by the UUPS function to validate the msg.sender when upgrading the AVS contract's implementation.
     */
    function _authorizeUpgrade(address) internal override onlyOwner {}
}
