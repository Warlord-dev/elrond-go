package systemSmartContracts

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/parsers"
	"github.com/ElrondNetwork/elrond-go/core/vmcommon"
	"github.com/ElrondNetwork/elrond-go/process/smartContract/hooks"
	"github.com/ElrondNetwork/elrond-go/vm"
	"github.com/ElrondNetwork/elrond-go/vm/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var configChangeAddress = []byte("config change address")

func createMockArgumentsForDelegationManager() ArgsNewDelegationManager {
	return ArgsNewDelegationManager{
		DelegationSCConfig: config.DelegationSystemSCConfig{
			MinServiceFee: 5,
			MaxServiceFee: 150,
		},
		DelegationMgrSCConfig: config.DelegationManagerSystemSCConfig{
			MinCreationDeposit: "10",
			MinStakeAmount:     "10",
		},
		Eei:                    &mock.SystemEIStub{},
		DelegationMgrSCAddress: vm.DelegationManagerSCAddress,
		StakingSCAddress:       vm.StakingSCAddress,
		ValidatorSCAddress:     vm.ValidatorSCAddress,
		ConfigChangeAddress:    configChangeAddress,
		GasCost:                vm.GasCost{MetaChainSystemSCsCost: vm.MetaChainSystemSCsCost{ESDTIssue: 10}},
		Marshalizer:            &mock.MarshalizerMock{},
		EpochNotifier:          &mock.EpochNotifierStub{},
	}
}

func getDefaultVmInputForDelegationManager(funcName string, args [][]byte) *vmcommon.ContractCallInput {
	return &vmcommon.ContractCallInput{
		VMInput: vmcommon.VMInput{
			CallerAddr:     []byte("addr"),
			Arguments:      args,
			CallValue:      big.NewInt(0),
			CallType:       0,
			GasPrice:       0,
			GasProvided:    0,
			OriginalTxHash: nil,
			CurrentTxHash:  nil,
		},
		RecipientAddr: []byte("addr"),
		Function:      funcName,
	}
}

func TestNewDelegationManagerSystemSC_NilEeiShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.Eei = nil

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	assert.Equal(t, vm.ErrNilSystemEnvironmentInterface, err)
}

func TestNewDelegationManagerSystemSC_InvalidStakingSCAddressShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.StakingSCAddress = nil

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	expectedErr := fmt.Errorf("%w for staking sc address", vm.ErrInvalidAddress)
	assert.Equal(t, expectedErr, err)
}

func TestNewDelegationManagerSystemSC_InvalidValidatorSCAddressShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.ValidatorSCAddress = nil

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	expectedErr := fmt.Errorf("%w for validator sc address", vm.ErrInvalidAddress)
	assert.Equal(t, expectedErr, err)
}

func TestNewDelegationManagerSystemSC_InvalidDelegationManagerSCAddressShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.DelegationMgrSCAddress = nil

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	expectedErr := fmt.Errorf("%w for delegation sc address", vm.ErrInvalidAddress)
	assert.Equal(t, expectedErr, err)
}

func TestNewDelegationManagerSystemSC_InvalidConfigChangeAddressShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.ConfigChangeAddress = nil

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	expectedErr := fmt.Errorf("%w for config change address", vm.ErrInvalidAddress)
	assert.Equal(t, expectedErr, err)
}

func TestNewDelegationManagerSystemSC_NilMarshalizerShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.Marshalizer = nil

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	assert.Equal(t, vm.ErrNilMarshalizer, err)
}

func TestNewDelegationManagerSystemSC_NilEpochNotifierShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.EpochNotifier = nil

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	assert.Equal(t, vm.ErrNilEpochNotifier, err)
}

func TestNewDelegationManagerSystemSC_InvalidMinCreationDepositShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.DelegationMgrSCConfig.MinCreationDeposit = "-10"

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	assert.Equal(t, vm.ErrInvalidMinCreationDeposit, err)
}

func TestNewDelegationManagerSystemSC_InvalidMinStakeAmountShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	args.DelegationMgrSCConfig.MinStakeAmount = "-10"

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, dm)
	assert.Equal(t, vm.ErrInvalidMinStakeValue, err)
}

func TestNewDelegationManagerSystemSC(t *testing.T) {
	t.Parallel()

	registerNotifyHandlerCalled := false
	args := createMockArgumentsForDelegationManager()
	args.EpochNotifier = &mock.EpochNotifierStub{
		RegisterNotifyHandlerCalled: func(handler core.EpochSubscriberHandler) {
			registerNotifyHandlerCalled = true
		}}

	dm, err := NewDelegationManagerSystemSC(args)
	assert.Nil(t, err)
	assert.NotNil(t, dm)
	assert.True(t, registerNotifyHandlerCalled)
}

func TestDelegationManagerSystemSC_ExecuteWithNilArgsShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	dm, _ := NewDelegationManagerSystemSC(args)

	output := dm.Execute(nil)
	assert.Equal(t, vmcommon.UserError, output)
}

func TestDelegationManagerSystemSC_ExecuteWithDelegationManagerDisabled(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	dm.delegationMgrEnabled.Unset()
	vmInput := getDefaultVmInputForDelegationManager("createNewDelegationContract", [][]byte{})

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, "delegation manager contract is not enabled"))
}

func TestDelegationManagerSystemSC_ExecuteInvalidFunction(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("func", [][]byte{})

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, "invalid function to call"))
}

func TestDelegationManagerSystemSC_ExecuteInit(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager(core.SCDeployInitFunctionName, [][]byte{})
	vmInput.CallValue = big.NewInt(15)

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrCallValueMustBeZero.Error()))

	vmInput.CallValue = big.NewInt(0)
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.Ok, output)

	dManagementData, _ := dm.getDelegationManagementData()
	assert.Equal(t, uint32(0), dManagementData.NumOfContracts)
	assert.Equal(t, vm.FirstDelegationSCAddress, dManagementData.LastAddress)
	assert.Equal(t, dm.minFee, dManagementData.MinServiceFee)
	assert.Equal(t, dm.maxFee, dManagementData.MaxServiceFee)
	assert.Equal(t, dm.minCreationDeposit, dManagementData.MinDeposit)

	dContractList, _ := dm.getDelegationContractList()
	assert.Equal(t, 1, len(dContractList.Addresses))
}

func TestDelegationManagerSystemSC_ExecuteCreateNewDelegationContractUserErrors(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("createNewDelegationContract", [][]byte{})
	dm.gasCost.MetaChainSystemSCsCost.DelegationMgrOps = 10

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.FunctionWrongSignature, output)
	assert.True(t, strings.Contains(eei.returnMessage, "wrong number of arguments"))

	vmInput.Arguments = [][]byte{{10}, {150}}
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.OutOfGas, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrNotEnoughGas.Error()))

	dm.gasCost.MetaChainSystemSCsCost.DelegationMgrOps = 0
	delegationsMap := map[string][]byte{}
	delegationsMap[string(vmInput.CallerAddr)] = []byte("deployed contract")
	eei.storageUpdate[string(eei.scAddress)] = delegationsMap
	vmInput.CallValue = big.NewInt(0)
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, "caller already deployed a delegation sc"))

	delete(delegationsMap, string(vmInput.CallerAddr))
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	expectedErr := fmt.Errorf("%w getDelegationManagementData", vm.ErrDataNotFoundUnderKey)
	assert.True(t, strings.Contains(eei.returnMessage, expectedErr.Error()))

	_ = dm.saveDelegationManagementData(&DelegationManagement{
		MinDeposit: big.NewInt(10),
	})
	vmInput.CallValue = big.NewInt(9)
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, "not enough call value"))

	vmInput.CallValue = big.NewInt(20)
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	expectedErr = fmt.Errorf("%w getDelegationContractList", vm.ErrDataNotFoundUnderKey)
	assert.True(t, strings.Contains(eei.returnMessage, expectedErr.Error()))
}

func createSystemSCContainer(eei *vmContext) vm.SystemSCContainer {
	argsStaking := createMockStakingScArguments()
	argsStaking.Eei = eei
	stakingSc, _ := NewStakingSmartContract(argsStaking)

	argsValidator := createMockArgumentsForValidatorSC()
	argsValidator.Eei = eei
	validatorSc, _ := NewValidatorSmartContract(argsValidator)

	delegationSCArgs := createMockArgumentsForDelegation()
	delegationSCArgs.Eei = eei
	delegationSc, _ := NewDelegationSystemSC(delegationSCArgs)

	systemSCContainer := &mock.SystemSCContainerStub{
		GetCalled: func(key []byte) (vm.SystemSmartContract, error) {
			switch string(key) {
			case string(vm.StakingSCAddress):
				return stakingSc, nil
			case string(vm.ValidatorSCAddress):
				return validatorSc, nil
			case string(vm.FirstDelegationSCAddress):
				return delegationSc, nil
			}
			return nil, nil
		},
	}

	return systemSCContainer
}

func TestDelegationManagerSystemSC_ExecuteCreateNewDelegationContract(t *testing.T) {
	t.Parallel()

	maxDelegationCap := []byte{250}
	serviceFee := []byte{10}
	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		parsers.NewCallArgsParser(),
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	_ = eei.SetSystemSCContainer(
		createSystemSCContainer(eei),
	)

	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("createNewDelegationContract", [][]byte{maxDelegationCap, serviceFee})

	_ = dm.saveDelegationContractList(&DelegationContractList{Addresses: make([][]byte, 0)})
	_ = dm.saveDelegationManagementData(&DelegationManagement{
		MinDeposit:  big.NewInt(10),
		LastAddress: vm.FirstDelegationSCAddress,
	})
	vmInput.CallValue = big.NewInt(20)

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.Ok, output)

	dManagement, _ := dm.getDelegationManagementData()
	assert.Equal(t, uint32(1), dManagement.NumOfContracts)
	expectedAddress := createNewAddress(vm.FirstDelegationSCAddress)
	assert.Equal(t, expectedAddress, dManagement.LastAddress)

	dList, _ := dm.getDelegationContractList()
	assert.Equal(t, 1, len(dList.Addresses))
	assert.Equal(t, expectedAddress, dList.Addresses[0])

	assert.Equal(t, 1, len(eei.output))
	assert.Equal(t, expectedAddress, eei.output[0])

	outAcc := eei.outputAccounts[string(expectedAddress)]
	assert.Equal(t, vm.FirstDelegationSCAddress, outAcc.Code)
	assert.Equal(t, vmInput.CallerAddr, outAcc.CodeDeployerAddress)

	codeMetaData := &vmcommon.CodeMetadata{
		Upgradeable: false,
		Payable:     false,
		Readable:    true,
	}
	expectedMetaData := codeMetaData.ToBytes()
	assert.Equal(t, expectedMetaData, outAcc.CodeMetadata)

	systemSc, _ := eei.systemContracts.Get(vm.FirstDelegationSCAddress)
	delegationSc := systemSc.(*delegation)
	eei.scAddress = createNewAddress(vm.FirstDelegationSCAddress)
	dContractConfig, _ := delegationSc.getDelegationContractConfig()
	retrievedOwnerAddress := eei.GetStorage([]byte(ownerKey))
	retrievedServiceFee := eei.GetStorage([]byte(serviceFeeKey))
	assert.Equal(t, vmInput.CallerAddr, retrievedOwnerAddress)
	assert.Equal(t, []byte{10}, retrievedServiceFee)
	assert.Equal(t, big.NewInt(250), dContractConfig.MaxDelegationCap)

	marshalledData := eei.GetStorageFromAddress(vm.ValidatorSCAddress, eei.scAddress)
	stakedData := &StakedDataV2_0{}
	err := args.Marshalizer.Unmarshal(stakedData, marshalledData)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(stakedData.RewardAddress, eei.scAddress))
}

func TestDelegationManagerSystemSC_ExecuteGetAllContractAddresses(t *testing.T) {
	t.Parallel()

	addr1 := []byte("addr1")
	addr2 := []byte("addr2")
	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("getAllContractAddresses", [][]byte{})

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrInvalidCaller.Error()))

	vmInput.CallerAddr = dm.delegationMgrSCAddress
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	expectedErr := fmt.Errorf("%w getDelegationContractList", vm.ErrDataNotFoundUnderKey)
	assert.True(t, strings.Contains(eei.returnMessage, expectedErr.Error()))

	_ = dm.saveDelegationContractList(&DelegationContractList{Addresses: [][]byte{addr1, addr2}})
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.Ok, output)
	assert.Equal(t, 1, len(eei.output))
	assert.Equal(t, addr2, eei.output[0])
}

func TestDelegationManagerSystemSC_ExecuteChangeMinDepositUserErrors(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("changeMinDeposit", [][]byte{})
	vmInput.CallValue = big.NewInt(10)

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrCallValueMustBeZero.Error()))

	vmInput.CallValue = big.NewInt(0)
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrInvalidNumOfArguments.Error()))

	vmInput.Arguments = [][]byte{{25}}
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrInvalidCaller.Error()))

	vmInput.CallerAddr = configChangeAddress
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	expectedErr := fmt.Errorf("%w getDelegationManagementData", vm.ErrDataNotFoundUnderKey)
	assert.True(t, strings.Contains(eei.returnMessage, expectedErr.Error()))
}

func TestDelegationManagerSystemSC_ExecuteChangeMinDeposit(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("changeMinDeposit", [][]byte{{25}})
	vmInput.CallerAddr = configChangeAddress
	_ = dm.saveDelegationManagementData(&DelegationManagement{MinDeposit: big.NewInt(0)})

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.Ok, output)

	dManagementData, _ := dm.getDelegationManagementData()
	assert.Equal(t, big.NewInt(25), dManagementData.MinDeposit)
}

func TestDelegationManager_ChangeMinDelegationAmountInvalidCallerShouldError(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("changeMinDelegationAmount", [][]byte{})
	vmInput.CallValue = big.NewInt(10)

	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrCallValueMustBeZero.Error()))

	vmInput.CallValue = big.NewInt(0)
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrInvalidNumOfArguments.Error()))

	vmInput.Arguments = [][]byte{{25}}
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrInvalidCaller.Error()))

	vmInput.CallerAddr = configChangeAddress
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	expectedErr := fmt.Errorf("%w getDelegationManagementData", vm.ErrDataNotFoundUnderKey)
	assert.True(t, strings.Contains(eei.returnMessage, expectedErr.Error()))

	_ = dm.saveDelegationManagementData(&DelegationManagement{MinDelegationAmount: big.NewInt(25)})
	vmInput.Arguments = [][]byte{{0}}
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, "invalid min delegation amount"))
}

func TestDelegationManager_ChangeMinDelegationMarhalizingFailsShouldError(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("expected error")
	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("changeMinDelegationAmount", [][]byte{})
	vmInput.Arguments = [][]byte{{25}}
	vmInput.CallerAddr = configChangeAddress

	_ = dm.saveDelegationManagementData(&DelegationManagement{MinDelegationAmount: big.NewInt(25)})
	dm.marshalizer = &mock.MarshalizerStub{
		MarshalCalled: func(obj interface{}) ([]byte, error) {
			return nil, expectedErr
		},
		UnmarshalCalled: func(obj interface{}, buff []byte) error {
			return nil
		},
	}
	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, expectedErr.Error()))
}

func TestDelegationManager_ChangeMinDelegationShouldWork(t *testing.T) {
	t.Parallel()

	newMinDelegationAmount := big.NewInt(224)
	existingMinDelegationAmount := big.NewInt(25)
	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("changeMinDelegationAmount", [][]byte{})
	vmInput.Arguments = [][]byte{newMinDelegationAmount.Bytes()}
	vmInput.CallerAddr = configChangeAddress

	_ = dm.saveDelegationManagementData(&DelegationManagement{MinDelegationAmount: existingMinDelegationAmount})
	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.Ok, output)

	recovered, err := dm.getDelegationManagementData()
	require.Nil(t, err)
	assert.Equal(t, newMinDelegationAmount, recovered.MinDelegationAmount)
}

func TestCreateNewAddress_NextAddressShouldWork(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		lastAddress         []byte
		expectedNextAddress []byte
	}

	tests := []*testStruct{
		{
			lastAddress:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 255, 255, 255},
			expectedNextAddress: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 255, 255, 255},
		},
		{
			lastAddress:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255},
			expectedNextAddress: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255},
		},
		{
			lastAddress:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255},
			expectedNextAddress: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 255, 255, 255},
		},
		{
			lastAddress:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 255, 255},
			expectedNextAddress: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 255, 255, 255},
		},
		{
			lastAddress:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 34, 23, 255, 255, 255, 255, 255, 255, 255, 255},
			expectedNextAddress: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 34, 24, 0, 0, 0, 0, 0, 255, 255, 255},
		},
	}

	for _, test := range tests {
		nextAddress := createNewAddress(test.lastAddress)
		assert.Equal(t, test.expectedNextAddress, nextAddress,
			fmt.Sprintf("expected: %v, got %d", test.expectedNextAddress, nextAddress))
	}
}

func TestDelegationManager_GetContractConfigErrors(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("getContractConfig", [][]byte{})
	vmInput.CallerAddr = []byte("not the correct caller")
	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	assert.True(t, strings.Contains(eei.returnMessage, vm.ErrInvalidCaller.Error()))

	//missing data
	vmInput.CallerAddr = vm.DelegationManagerSCAddress
	output = dm.Execute(vmInput)
	assert.Equal(t, vmcommon.UserError, output)
	expectedErr := fmt.Errorf("%w getDelegationManagementData", vm.ErrDataNotFoundUnderKey)
	assert.True(t, strings.Contains(eei.returnMessage, expectedErr.Error()))
}

func TestDelegationManager_GetContractConfigShouldWork(t *testing.T) {
	t.Parallel()

	args := createMockArgumentsForDelegationManager()
	eei, _ := NewVMContext(
		&mock.BlockChainHookStub{},
		hooks.NewVMCryptoHook(),
		&mock.ArgumentParserMock{},
		&mock.AccountsStub{},
		&mock.RaterMock{},
	)
	args.Eei = eei

	dm, _ := NewDelegationManagerSystemSC(args)
	vmInput := getDefaultVmInputForDelegationManager("getContractConfig", [][]byte{})

	delegationManagement := &DelegationManagement{
		NumOfContracts:      123,
		LastAddress:         []byte("last address"),
		MinServiceFee:       456,
		MaxServiceFee:       789,
		MinDeposit:          big.NewInt(112233),
		MinDelegationAmount: big.NewInt(445566),
	}

	_ = dm.saveDelegationManagementData(delegationManagement)

	vmInput.CallerAddr = vm.DelegationManagerSCAddress
	output := dm.Execute(vmInput)
	assert.Equal(t, vmcommon.Ok, output)

	results := eei.CreateVMOutput()
	//this test also verify the position in results.ReturnData
	assert.Equal(t, big.NewInt(int64(delegationManagement.NumOfContracts)).Bytes(), results.ReturnData[0])
	assert.Equal(t, delegationManagement.LastAddress, results.ReturnData[1])
	assert.Equal(t, big.NewInt(int64(delegationManagement.MinServiceFee)).Bytes(), results.ReturnData[2])
	assert.Equal(t, big.NewInt(int64(delegationManagement.MaxServiceFee)).Bytes(), results.ReturnData[3])
	assert.Equal(t, delegationManagement.MinDeposit.Bytes(), results.ReturnData[4])
	assert.Equal(t, delegationManagement.MinDelegationAmount.Bytes(), results.ReturnData[5])
}
