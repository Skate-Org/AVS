// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.20;

import {IDelegationManager as IDM} from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

/**
 * @title Delegation Manager
 * @notice Add unexported public functions to the Delegation Manager interface.
 */
interface IDelegationManager is IDM {
    function getDelegatableShares(address operator) external view returns (IStrategy[] memory, uint256[] memory);
}
