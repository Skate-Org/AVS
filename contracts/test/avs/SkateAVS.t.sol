// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.20;

import {Test, console2} from "forge-std/Test.sol";
import {ISkateAVS} from "../../src/avs/interfaces/ISkateAVS.sol";
import {SkateAVS} from "../../src/avs/SkateAVS.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import {IStrategyManager, IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";

import {IDelegationManager} from "../../src/avs/interfaces/IDelegationManager.sol";

import {Errors} from "../../src/avs/Errors.sol";
import {BN254} from "../../src/avs/libraries/BN254.sol";
import {BLS} from "../../src/avs/libraries/BLS.sol";

import {Vm} from "forge-std/Vm.sol";

contract SkateAVSTest is Test {
    IAVSDirectory avsDirectory = IAVSDirectory(0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF);
    IDelegationManager delegationManager = IDelegationManager(0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A);
    IStrategyManager strategyManager = IStrategyManager(0x858646372CC42E1A627fcE94aa7A7033e7CF075A);
    ISkateAVS.StrategyParams strategyParam = ISkateAVS.StrategyParams(IStrategy(0x93c4b944D05dfe6df7645A86cd2206016c51564D), 1e18);
    IERC20 strategyToken = IERC20(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
    ISkateAVS avs;

    /////////////////////////////////////////////////////////////////////////////////////////////////////////
    /////////////////////////////// BLS constant generated off-chain ////////////////////////////////////////

    // Private key for 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 - anvil test wallet 0
    address addr0 = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    uint256 privateKey0 = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
    uint256 blsPrivateKey0 = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 % BN254.R;
    BN254.G2Point pubKey0 =
        BN254.G2Point(
            [
                11829690919498240492717914986943835720956660440597679675829055482637656108571,
                2332637959155764237114944033008161729347763881858201243151775165598748610464
            ],
            [
                16844159549615964720138902510035506195836911611113351665735906210620230569546,
                16479896805003260112387785137189286363984636687122540735037453711640344896777
            ]
        );

    // Private key for 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 - anvil test wallet 1
    address addr1 = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    uint256 privateKey1 = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d;
    uint256 blsPrivateKey1 = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d % BN254.R;
    BN254.G2Point pubKey1 =
        BN254.G2Point(
            [
                16553675810662402913332642109180320946410252990138787612031442987029708414951,
                3950887887877429316505810968405340209351770580495254664272056647280823110872
            ],
            [
                19767199458582240114868020780170509926238616450007142624546444508397469746103,
                17996234567193988288700669608184288873172459020548483946003084471765254786022
            ]
        );

    // Private key for 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC - anvil test wallet 1
    address addr2 = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    uint256 privateKey2 = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a;
    uint256 blsPrivateKey2 = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a % BN254.R;
    BN254.G2Point pubKey2 =
        BN254.G2Point(
            [
                9792872845262080300649079268741298251040032822358403315276048280727343356779,
                15847210887164936986146813501197001106156225637036600642599857938920002570208
            ],
            [
                14792701865640318097083912335496875808840104722124871661180771905152637430638,
                15078585932495272831110558597484648487865512330810900597399063864832124095287
            ]
        );

    // just be explicit
    BN254.G2Point zeroG2 = BN254.G2Point([uint256(0), 0], [uint256(0), 0]);
    BN254.G1Point zeroG1 = BN254.G1Point(0, 0);

    //////////////////////////////////////// END BLS constants //////////////////////////////////////////////
    /////////////////////////////////////////////////////////////////////////////////////////////////////////

    error EnforcedPause();
    error OwnableUnauthorizedAccount(address);

    function setUp() external {
        vm.createSelectFork(vm.rpcUrl("https://eth-mainnet.g.alchemy.com/v2/Kc1-RaFsfCaZwuK0HpYO7QE07jaZZnfv"));
        ISkateAVS.StrategyParams[] memory strategies = new ISkateAVS.StrategyParams[](1);
        strategies[0] = ISkateAVS.StrategyParams(strategyParam.strategy, strategyParam.multiplier);

        address avsImpl = address(new SkateAVS(avsDirectory, delegationManager));
        avs = ISkateAVS(
            address(
                new ERC1967Proxy(
                    avsImpl,
                    abi.encodeWithSignature("initialize(address,(address,uint96)[],string,bool)", address(this), strategies, "", false)
                )
            )
        );
    }

    function testDeployment() external view {
        assertEq(avs.maxOperatorCount(), 5);
        assertEq(avs.minOperatorStake(), 1);
        assertEq(avs.strategies().length, 1);
        assertEq(address(avs.avsDirectory()), address(avsDirectory));
        assertEq(avs.operators().length, 0);
        assertEq(avs.getRestakeableStrategies().length, 1);
    }

    function testRegisterOperatorWhenAVSPaused() external {
        avs.pause();
        _fundAddress(addr0);
        ISignatureUtils.SignatureWithSaltAndExpiry memory emptyOperatorSignature;
        vm.expectRevert(EnforcedPause.selector);
        vm.prank(addr0);
        avs.registerOperatorToAVS(addr0, emptyOperatorSignature, zeroG2); // registration doesn't matter
    }

    function testRegisterOperatorWithNonOperatorAddress() external {
        ISignatureUtils.SignatureWithSaltAndExpiry memory emptyOperatorSignature;
        address wallet = addr0;
        _fundAddress(wallet);
        vm.expectRevert(Errors.OnlyOperatorAllowedToCall.selector);
        vm.prank(wallet);
        avs.registerOperatorToAVS(address(0x1), emptyOperatorSignature, zeroG2);
    }

    function testRegisterWithOperatorNotAllowed() external {
        avs.enableAllowlist();
        Vm.Wallet memory wallet;
        wallet.addr = addr0;
        wallet.privateKey = privateKey0;

        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(wallet);
        vm.expectRevert(Errors.OperatorNotAllowed.selector);
        vm.prank(wallet.addr);
        avs.registerOperatorToAVS(wallet.addr, operatorSignature, zeroG2);
    }

    // function testRegisterOperatorWithOperatorAlreadyRegistered() external {
    //     Vm.Wallet memory wallet = _setupWallet();
    //     _stakeAndRegisterAsOperator(wallet);
    //
    //     ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(wallet);
    //     vm.prank(wallet.addr);
    //     avs.registerOperatorToAVS(wallet.addr, operatorSignature);
    //
    //     vm.prank(wallet.addr);
    //     vm.expectRevert(Errors.AlreadyAnOperator.selector);
    //     avs.registerOperatorToAVS(wallet.addr, operatorSignature);
    // }

    function testRegisterOperatorWithMaxCountReached() external {
        //        Vm.Wallet memory wallet = _setupWallet();
        //        _stakeAndRegisterAsOperator(wallet);
        //
        //        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(wallet);
        //        bytes memory pubKey = _getPubKey(wallet);
        //
        //        vm.prank(wallet.addr);
        //        avs.register(pubKey, operatorSignature);
        //
        //        Vm.Wallet memory wallet1 = vm.createWallet(uint256(keccak256(abi.encode("testname", uint256(2)))));
        //        vm.prank(0xE53FFF67f9f384d20Ebea36F43b93DC49Ed22753);
        //        strategyToken.transfer(wallet1.addr, 100 ether);
        //
        //        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature1 = _prepareOperatorSignature(wallet1);
        //        bytes memory pubKey1 = _getPubKey(wallet1);
        //
        //        vm.prank(wallet1.addr);
        //        vm.expectRevert(Errors.MaxOperatorCountReached.selector);
        //        avs.register(pubKey1, operatorSignature1);
    }

    // function testRegisterOperatorWithLessThanMinOperatorStake() external {
    //     Vm.Wallet memory wallet = _setupWallet();
    //     _stakeAndRegisterAsOperator(wallet);
    //
    //     ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(wallet);
    //     avs.setMinOperatorStake(100_000 ether);
    //
    //     vm.prank(wallet.addr);
    //     vm.expectRevert(Errors.MinOperatorStakeNotSatisfied.selector);
    //     avs.registerOperatorToAVS(wallet.addr, operatorSignature);
    // }

    function testRegister() external {
        Vm.Wallet memory wallet;
        wallet.addr = addr0;
        wallet.privateKey = privateKey0;

        _fundAddress(addr0);
        _stakeAndRegisterAsOperator(wallet.addr);

        assertEq(avs.canRegister(wallet.addr), true);
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(wallet);

        // BN254.G1Point memory registrationSignature = BLS.signMessage(blsPrivateKey0, avs.pubkeyRegistrationMessage(wallet.addr));
        vm.prank(wallet.addr);
        avs.registerOperatorToAVS(wallet.addr, operatorSignature, pubKey0);

        assertEq(avs.operators()[0].addr, wallet.addr);
    }

    function testDeregisterOperatorByNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.deregisterOperatorFromAVS(address(0x1));
    }

    // function testderegister() external {
    //     Vm.Wallet memory wallet;
    //     wallet.addr = wallet0;
    //     wallet.privateKey = privateKey0;
    //     _stakeAndRegisterAsOperator(wallet);
    //
    //     ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(wallet);
    //     assertEq(avs.operators().length, 0);
    //
    //     vm.prank(wallet.addr);
    //
    //     avs.registerOperatorToAVS(wallet.addr, operatorSignature);
    //     assertEq(avs.operators().length, 1);
    //
    //     avs.deregisterOperatorFromAVS(wallet.addr);
    //     assertEq(avs.operators().length, 0);
    // }

    function testSetMetadataURIByNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.updateAVSMetadataURI("");
    }

    function testSetMetadataURI() external {
        avs.updateAVSMetadataURI("123");
    }

    function testSetStrategiesByNonOwner() external {
        ISkateAVS.StrategyParams[] memory strategies;

        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.setStrategies(strategies);
    }

    function testSetStrategies() external {
        assertEq(avs.strategies().length, 1);
        ISkateAVS.StrategyParams[] memory strategies;
        avs.setStrategies(strategies);
        assertEq(avs.strategies().length, 0);
    }

    function testSetMinOperatorStakeByNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.setMinOperatorStake(0);
    }

    function testSetMinOperatorStake() external {
        assertEq(avs.minOperatorStake(), 1);
        avs.setMinOperatorStake(2);
        assertEq(avs.minOperatorStake(), 2);
    }

    function testSetMaxOperatorCountByNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.setMaxOperatorCount(0);
    }

    function testSetMaxOperatorCount() external {
        assertEq(avs.maxOperatorCount(), 5);
        avs.setMaxOperatorCount(2);
        assertEq(avs.maxOperatorCount(), 2);
    }

    function testAddToAllowlistZeroAddress() external {
        vm.expectRevert(Errors.ZeroOperatorAddress.selector);
        avs.addToAllowlist(address(0x0));
    }

    function testAddToAllowlistNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.addToAllowlist(address(0x2));
    }

    function testAddToAllowlist() external {
        address operator = address(0x2);
        assertEq(avs.isInAllowlist(operator), false);
        avs.addToAllowlist(operator);
        assertEq(avs.isInAllowlist(operator), true);

        vm.expectRevert(Errors.OperatorAlreadyInAllowlist.selector);
        avs.addToAllowlist(operator);
    }

    function testRemoveFromAllowlistByNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.removeFromAllowlist(address(0x2));
    }

    function testRemoveFromAllowlist() external {
        address operator = address(0x2);
        avs.addToAllowlist(operator);
        assertEq(avs.isInAllowlist(operator), true);

        avs.removeFromAllowlist(operator);
        assertEq(avs.isInAllowlist(operator), false);
    }

    function testEnableAllowlistByNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.enableAllowlist();
    }

    function testEnableAllowlist() external {
        assertEq(avs.allowlistEnabled(), false);
        avs.enableAllowlist();
        assertEq(avs.allowlistEnabled(), true);

        vm.expectRevert(Errors.AllowlistAlreadyEnabled.selector);
        avs.enableAllowlist();
    }

    function testDisableAllowlistByNonOwner() external {
        vm.prank(address(0x1));
        vm.expectRevert(abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x1)));
        avs.disableAllowlist();
    }

    function testDisableAllowlist() external {
        avs.enableAllowlist();
        assertEq(avs.allowlistEnabled(), true);

        avs.disableAllowlist();
        assertEq(avs.allowlistEnabled(), false);

        vm.expectRevert(Errors.AllowlistAlreadyDisabled.selector);
        avs.disableAllowlist();
    }

    function testSubmitDataWhenAggregationOrderMixed() external {
        uint256 taskId = 1;
        bytes memory messageData = bytes("hello, world!");
        bytes32 messageDigest = keccak256(abi.encodePacked(taskId, messageData));

        // NOTE: order the wallets beforehand
        uint256[] memory privateKeys = new uint256[](3);
        privateKeys[0] = privateKey2;
        privateKeys[1] = privateKey1;
        privateKeys[2] = privateKey0;

        BN254.G2Point[] memory pubKeys = new BN254.G2Point[](3);
        pubKeys[0] = pubKey2;
        pubKeys[1] = pubKey1;
        pubKeys[2] = pubKey0;

        address[] memory operators = new address[](3);
        operators[0] = addr2;
        operators[1] = addr1;
        operators[2] = addr0;

        for (uint256 i = 0; i < operators.length; i++) {
            // create account
            vm.prank(0xE53FFF67f9f384d20Ebea36F43b93DC49Ed22753);
            strategyToken.transfer(operators[i], 10 ether);

            // register to EigenLayer
            _stakeAndRegisterAsOperator(operators[i]);

            // register as operator
            // bytes32 blsRegistrationMessage = avs.pubkeyRegistrationMessage(operators[i]);
            // BN254.G1Point memory registrationSignature = BLS.signMessage(privateKeys[i] % BN254.R, blsRegistrationMessage);

            ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(privateKeys[i], operators[i]);
            vm.prank(operators[i]);
            avs.registerOperatorToAVS(operators[i], operatorSignature, pubKeys[i]);

            // sign the task message
        }

        BN254.G1Point memory aggSignature;
        // NOTE: signature order is tampered
        aggSignature = BN254.addG1(aggSignature, BLS.signMessage(privateKeys[1] % BN254.R, messageDigest));
        aggSignature = BN254.addG1(aggSignature, BLS.signMessage(privateKeys[0] % BN254.R, messageDigest));
        aggSignature = BN254.addG1(aggSignature, BLS.signMessage(privateKeys[2] % BN254.R, messageDigest));

        // vm.expectRevert(Errors.InvalidBLSSignature.selector);
        avs.submitData(taskId, messageData, operators, aggSignature);
    }

    function testSubmitData_RevertTamperedSignature() external {
        uint256 taskId = 1;
        bytes memory messageData = bytes("hello, world!");
        bytes32 messageDigest = keccak256(abi.encodePacked(taskId, messageData));

        // NOTE: order the wallets beforehand
        uint256[] memory privateKeys = new uint256[](1);
        privateKeys[0] = privateKey2;

        BN254.G2Point[] memory pubKeys = new BN254.G2Point[](1);
        pubKeys[0] = pubKey2;

        address[] memory operators = new address[](1);
        operators[0] = addr2;

        for (uint256 i = 0; i < operators.length; i++) {
            // create account
            vm.prank(0xE53FFF67f9f384d20Ebea36F43b93DC49Ed22753);
            strategyToken.transfer(operators[i], 10 ether);

            // register to EigenLayer
            _stakeAndRegisterAsOperator(operators[i]);

            // register as operator
            // bytes32 blsRegistrationMessage = avs.pubkeyRegistrationMessage(operators[i]);
            // BN254.G1Point memory registrationSignature = BLS.signMessage(privateKeys[i] % BN254.R, blsRegistrationMessage);

            ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(privateKeys[i], operators[i]);
            vm.prank(operators[i]);
            avs.registerOperatorToAVS(operators[i], operatorSignature, pubKeys[i]);

            // sign the task message
        }

        // NOTE: signature is tampered
        BN254.G1Point memory aggSignature = BN254.G1Point(1, 2); // start from G1 instead of 0 -> invalid signature
        aggSignature = BN254.addG1(aggSignature, BLS.signMessage(privateKeys[0] % BN254.R, messageDigest));

        vm.expectRevert(Errors.InvalidBLSSignature.selector);
        avs.submitData(taskId, messageData, operators, aggSignature);
    }

    function testSubmitData_RevertQuorumNotReached() external {
        uint256 taskId = 1;
        bytes memory messageData = bytes("hello, world!");

        // NOTE: order the wallets beforehand
        uint256[] memory privateKeys = new uint256[](2);
        privateKeys[0] = privateKey2;
        privateKeys[1] = privateKey1;

        BN254.G2Point[] memory pubKeys = new BN254.G2Point[](2);
        pubKeys[0] = pubKey2;
        pubKeys[1] = pubKey1;

        address[] memory operators = new address[](2);
        operators[0] = addr2;
        operators[1] = addr1;

        BN254.G1Point memory aggSignature;
        for (uint256 i = 0; i < operators.length; i++) {
            // create account
            vm.prank(0xE53FFF67f9f384d20Ebea36F43b93DC49Ed22753);
            strategyToken.transfer(operators[i], 10 ether);

            // register to EigenLayer
            _stakeAndRegisterAsOperator(operators[i]);

            // bytes32 blsRegistrationMessage = avs.pubkeyRegistrationMessage(operators[i]);
            // BN254.G1Point memory registrationSignature = BLS.signMessage(privateKeys[i] % BN254.R, blsRegistrationMessage);

            // register as operator
            ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(privateKeys[i], operators[i]);
            vm.prank(operators[i]);
            avs.registerOperatorToAVS(operators[i], operatorSignature, pubKeys[i]);
        }

        address[] memory submitOperators = new address[](1);
        operators[0] = addr2;

        vm.expectRevert(Errors.QuorumNotReached.selector);
        avs.submitData(taskId, messageData, submitOperators, aggSignature);
    }

    function testSubmitData() external {
        uint256 taskId = 1;
        bytes memory messageData = bytes("hello, world!");
        bytes32 messageDigest = keccak256(abi.encodePacked(taskId, messageData));

        // NOTE: order the wallets beforehand
        uint256[] memory privateKeys = new uint256[](3);
        privateKeys[0] = privateKey2;
        privateKeys[1] = privateKey1;
        privateKeys[2] = privateKey0;

        BN254.G2Point[] memory pubKeys = new BN254.G2Point[](3);
        pubKeys[0] = pubKey2;
        pubKeys[1] = pubKey1;
        pubKeys[2] = pubKey0;

        address[] memory operators = new address[](3);
        operators[0] = addr2;
        operators[1] = addr1;
        operators[2] = addr0;

        BN254.G1Point memory aggSignature;
        for (uint256 i = 0; i < operators.length; i++) {
            // create account
            vm.prank(0xE53FFF67f9f384d20Ebea36F43b93DC49Ed22753);
            strategyToken.transfer(operators[i], 10 ether);

            // register to EigenLayer
            _stakeAndRegisterAsOperator(operators[i]);

            // register as operator
            // bytes32 blsRegistrationMessage = avs.pubkeyRegistrationMessage(operators[i]);
            // BN254.G1Point memory registrationSignature = BLS.signMessage(privateKeys[i] % BN254.R, blsRegistrationMessage);

            ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(privateKeys[i], operators[i]);
            vm.prank(operators[i]);
            avs.registerOperatorToAVS(operators[i], operatorSignature, pubKeys[i]);

            // sign the task message
            BN254.G1Point memory signature = BLS.signMessage(privateKeys[i] % BN254.R, messageDigest);
            aggSignature = BN254.addG1(aggSignature, signature);
        }

        avs.submitData(taskId, messageData, operators, aggSignature);
    }

    // function testSubmitBatchData() external {
    //     uint256 taskId = 1;
    //     bytes memory message = bytes("hello, world!");
    //     bytes32 digest = MessageHashUtils.toEthSignedMessageHash(keccak256(abi.encodePacked(taskId, message)));
    //
    //     Vm.Wallet[] memory wallets = new Vm.Wallet[](3);
    //     ISkateAVS.SignatureTuple[] memory signatureTuples = new ISkateAVS.SignatureTuple[](wallets.length);
    //     for (uint256 i = 0; i < wallets.length; i++) {
    //         wallets[i] = vm.createWallet(uint256(keccak256(abi.encode("testname", uint256(i)))));
    //         vm.prank(0xE53FFF67f9f384d20Ebea36F43b93DC49Ed22753);
    //         strategyToken.transfer(wallets[i].addr, 100 ether);
    //         _stakeAndRegisterAsOperator(wallets[i]);
    //
    //         ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _prepareOperatorSignature(wallets[i]);
    //         vm.prank(wallets[i].addr);
    //         avs.registerOperatorToAVS(wallets[i].addr, operatorSignature);
    //
    //         (uint8 v, bytes32 r, bytes32 s) = vm.sign(wallets[i].privateKey, digest);
    //         signatureTuples[i].operator = wallets[i].addr;
    //         signatureTuples[i].signature = abi.encodePacked(r, s, v);
    //     }
    //     signatureTuples = _sortAddress(signatureTuples);
    //
    //     uint256[] memory taskIds = new uint256[](1);
    //     taskIds[0] = taskId;
    //
    //     bytes[] memory messages = new bytes[](1);
    //     messages[0] = message;
    //
    //     ISkateAVS.SignatureTuple[][] memory signaturesTuples = new ISkateAVS.SignatureTuple[][](1);
    //     signaturesTuples[0] = signatureTuples;
    //
    //     avs.batchSubmitData(taskIds, messages, signaturesTuples);
    // }

    function testUpgradeByNonOwner() external {
        vm.prank(address(0x2));
        (bool success, bytes memory data) = address(avs).call(abi.encodeWithSignature("upgradeToAndCall(address,bytes)", address(0x1), ""));
        assertEq(success, false);
        assertEq(data, abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(0x2)));
    }

    function testUpgrade() external {
        address newImpl = address(new SkateAVS(avsDirectory, delegationManager));
        (bool success, ) = address(avs).call(abi.encodeWithSignature("upgradeToAndCall(address,bytes)", newImpl, ""));
        assertEq(success, true);
    }

    function _fundAddress(address wallet) private {
        // Vm.Wallet memory wallet = vm.createWallet(uint256(keccak256(abi.encode("testname", uint256(1)))));
        vm.prank(0xE53FFF67f9f384d20Ebea36F43b93DC49Ed22753);
        strategyToken.transfer(wallet, 100 ether);
        // return wallet;
    }

    function _stakeAndRegisterAsOperator(address addr) private {
        vm.startPrank(addr);
        strategyToken.approve(address(strategyManager), 1 ether);
        strategyManager.depositIntoStrategy(strategyParam.strategy, strategyToken, 1 ether);

        IDelegationManager.OperatorDetails memory operatorDetails;
        operatorDetails.earningsReceiver = addr;
        operatorDetails.delegationApprover = addr;
        operatorDetails.stakerOptOutWindowBlocks = 100;
        delegationManager.registerAsOperator(operatorDetails, "");
        vm.stopPrank();
    }

    function _prepareOperatorSignature(
        uint256 privateKey,
        address addr
    ) private view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            privateKey,
            avsDirectory.calculateOperatorAVSRegistrationDigestHash(addr, address(avs), bytes32(uint256(1)), block.timestamp + 1 minutes)
        );
        operatorSignature.signature = abi.encodePacked(r, s, v);
        operatorSignature.salt = bytes32(uint256(1));
        operatorSignature.expiry = block.timestamp + 1 minutes;
    }

    function _prepareOperatorSignature(
        Vm.Wallet memory wallet
    ) private view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            wallet.privateKey,
            avsDirectory.calculateOperatorAVSRegistrationDigestHash(wallet.addr, address(avs), bytes32(uint256(1)), block.timestamp + 1 minutes)
        );
        operatorSignature.signature = abi.encodePacked(r, s, v);
        operatorSignature.salt = bytes32(uint256(1));
        operatorSignature.expiry = block.timestamp + 1 minutes;
    }

    function _sortAddress(address[] memory data) private pure returns (address[] memory) {
        // Perform bubble sort
        for (uint256 i = 0; i < data.length - 1; i++) {
            for (uint256 j = 0; j < data.length - i - 1; j++) {
                if (data[j] > data[j + 1]) {
                    // Swap elements
                    (data[j], data[j + 1]) = (data[j + 1], data[j]);
                }
            }
        }

        return data;
    }
}
