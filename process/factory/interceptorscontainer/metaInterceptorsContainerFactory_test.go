package interceptorscontainer_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/p2p"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/process/factory"
	"github.com/ElrondNetwork/elrond-go/process/factory/interceptorscontainer"
	"github.com/ElrondNetwork/elrond-go/process/mock"
	"github.com/ElrondNetwork/elrond-go/storage"
	"github.com/stretchr/testify/assert"
)

const maxTxNonceDeltaAllowed = 100

var chainID = []byte("chain ID")
var errExpected = errors.New("expected error")

func createMetaStubTopicHandler(matchStrToErrOnCreate string, matchStrToErrOnRegister string) process.TopicHandler {
	return &mock.TopicHandlerStub{
		CreateTopicCalled: func(name string, createChannelForTopic bool) error {
			if matchStrToErrOnCreate == "" {
				return nil
			}

			if strings.Contains(name, matchStrToErrOnCreate) {
				return errExpected
			}

			return nil
		},
		RegisterMessageProcessorCalled: func(topic string, handler p2p.MessageProcessor) error {
			if matchStrToErrOnRegister == "" {
				return nil
			}

			if strings.Contains(topic, matchStrToErrOnRegister) {
				return errExpected
			}

			return nil
		},
	}
}

func createMetaDataPools() dataRetriever.PoolsHolder {
	pools := &mock.PoolsHolderStub{
		HeadersCalled: func() dataRetriever.HeadersPool {
			return &mock.HeadersCacherStub{}
		},
		MiniBlocksCalled: func() storage.Cacher {
			return &mock.CacherStub{}
		},
		TransactionsCalled: func() dataRetriever.ShardedDataCacherNotifier {
			return &mock.ShardedDataStub{}
		},
		UnsignedTransactionsCalled: func() dataRetriever.ShardedDataCacherNotifier {
			return &mock.ShardedDataStub{}
		},
		TrieNodesCalled: func() storage.Cacher {
			return &mock.CacherStub{}
		},
		RewardTransactionsCalled: func() dataRetriever.ShardedDataCacherNotifier {
			return &mock.ShardedDataStub{}
		},
	}

	return pools
}

func createMetaStore() *mock.ChainStorerMock {
	return &mock.ChainStorerMock{
		GetStorerCalled: func(unitType dataRetriever.UnitType) storage.Storer {
			return &mock.StorerStub{}
		},
	}
}

//------- NewInterceptorsContainerFactory

func TestNewMetaInterceptorsContainerFactory_NilShardCoordinatorShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.ShardCoordinator = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilShardCoordinator, err)
}

func TestNewMetaInterceptorsContainerFactory_NilNodesCoordinatorShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.NodesCoordinator = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilNodesCoordinator, err)
}

func TestNewMetaInterceptorsContainerFactory_NilTopicHandlerShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Messenger = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilMessenger, err)
}

func TestNewMetaInterceptorsContainerFactory_NilStoreShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Store = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilStore, err)
}

func TestNewMetaInterceptorsContainerFactory_NilMarshalizerShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.ProtoMarshalizer = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilMarshalizer, err)
}

func TestNewMetaInterceptorsContainerFactory_NilMarshalizerAndSizeCheckShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.ProtoMarshalizer = nil
	args.SizeCheckDelta = 1
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilMarshalizer, err)
}

func TestNewMetaInterceptorsContainerFactory_NilHasherShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Hasher = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilHasher, err)
}

func TestNewMetaInterceptorsContainerFactory_NilMultiSignerShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.MultiSigner = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilMultiSigVerifier, err)
}

func TestNewMetaInterceptorsContainerFactory_NilDataPoolShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.DataPool = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilDataPoolHolder, err)
}

func TestNewMetaInterceptorsContainerFactory_NilAccountsShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Accounts = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilAccountsAdapter, err)
}

func TestNewMetaInterceptorsContainerFactory_NilAddrConvShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.AddrConverter = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilAddressConverter, err)
}

func TestNewMetaInterceptorsContainerFactory_NilSingleSignerShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.SingleSigner = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilSingleSigner, err)
}

func TestNewMetaInterceptorsContainerFactory_NilKeyGenShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.KeyGen = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilKeyGen, err)
}

func TestNewMetaInterceptorsContainerFactory_NilFeeHandlerShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.TxFeeHandler = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilEconomicsFeeHandler, err)
}

func TestNewMetaInterceptorsContainerFactory_NilBlackListHandlerShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.BlackList = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilBlackListHandler, err)
}

func TestNewMetaInterceptorsContainerFactory_EmptyCahinIDShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.ChainID = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrInvalidChainID, err)
}

func TestNewMetaInterceptorsContainerFactory_NilValidityAttesterShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.ValidityAttester = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilValidityAttester, err)
}

func TestNewMetaInterceptorsContainerFactory_EpochStartTriggerShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.EpochStartTrigger = nil
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.Nil(t, icf)
	assert.Equal(t, process.ErrNilEpochStartTrigger, err)
}
func TestNewMetaInterceptorsContainerFactory_ShouldWork(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.NotNil(t, icf)
	assert.Nil(t, err)
}

func TestNewMetaInterceptorsContainerFactory_ShouldWorkWithSizeCheck(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.SizeCheckDelta = 1
	icf, err := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	assert.NotNil(t, icf)
	assert.Nil(t, err)
	assert.False(t, icf.IsInterfaceNil())
}

//------- Create

func TestMetaInterceptorsContainerFactory_CreateTopicMetablocksFailsShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Messenger = createMetaStubTopicHandler(factory.MetachainBlocksTopic, "")
	icf, _ := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	container, err := icf.Create()

	assert.Nil(t, container)
	assert.Equal(t, errExpected, err)
}

func TestMetaInterceptorsContainerFactory_CreateTopicShardHeadersForMetachainFailsShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Messenger = createMetaStubTopicHandler(factory.ShardBlocksTopic, "")
	icf, _ := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	container, err := icf.Create()

	assert.Nil(t, container)
	assert.Equal(t, errExpected, err)
}

func TestMetaInterceptorsContainerFactory_CreateRegisterForMetablocksFailsShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Messenger = createMetaStubTopicHandler("", factory.MetachainBlocksTopic)
	icf, _ := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	container, err := icf.Create()

	assert.Nil(t, container)
	assert.Equal(t, errExpected, err)
}

func TestMetaInterceptorsContainerFactory_CreateRegisterShardHeadersForMetachainFailsShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Messenger = createMetaStubTopicHandler("", factory.MetachainBlocksTopic)
	icf, _ := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	container, err := icf.Create()

	assert.Nil(t, container)
	assert.Equal(t, errExpected, err)
}

func TestMetaInterceptorsContainerFactory_CreateRegisterTrieNodesFailsShouldErr(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Messenger = createMetaStubTopicHandler("", factory.AccountTrieNodesTopic)
	icf, _ := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	container, err := icf.Create()

	assert.Nil(t, container)
	assert.Equal(t, errExpected, err)
}

func TestMetaInterceptorsContainerFactory_CreateShouldWork(t *testing.T) {
	t.Parallel()

	args := getArgumentsMeta()
	args.Messenger = &mock.TopicHandlerStub{
		CreateTopicCalled: func(name string, createChannelForTopic bool) error {
			return nil
		},
		RegisterMessageProcessorCalled: func(topic string, handler p2p.MessageProcessor) error {
			return nil
		},
	}
	icf, _ := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	container, err := icf.Create()

	assert.NotNil(t, container)
	assert.Nil(t, err)
}

func TestMetaInterceptorsContainerFactory_With4ShardsShouldWork(t *testing.T) {
	t.Parallel()

	noOfShards := 4

	shardCoordinator := mock.NewMultipleShardsCoordinatorMock()
	shardCoordinator.SetNoShards(uint32(noOfShards))
	shardCoordinator.CurrentShard = 1

	nodesCoordinator := &mock.NodesCoordinatorMock{
		ShardConsensusSize: 1,
		MetaConsensusSize:  1,
		NbShards:           uint32(noOfShards),
		ShardId:            1,
	}

	args := getArgumentsMeta()
	args.ShardCoordinator = shardCoordinator
	args.NodesCoordinator = nodesCoordinator
	args.Messenger = &mock.TopicHandlerStub{
		CreateTopicCalled: func(name string, createChannelForTopic bool) error {
			return nil
		},
		RegisterMessageProcessorCalled: func(topic string, handler p2p.MessageProcessor) error {
			return nil
		},
	}
	icf, _ := interceptorscontainer.NewMetaInterceptorsContainerFactory(args)

	container, err := icf.Create()

	numInterceptorsMetablock := 1
	numInterceptorsShardHeadersForMetachain := noOfShards
	numInterceptorsTransactionsForMetachain := noOfShards + 1
	numInterceptorsMiniBlocksForMetachain := noOfShards + 1 + 1
	numInterceptorsUnsignedTxsForMetachain := noOfShards
	numInterceptorsRewardsTxsForMetachain := noOfShards
	numInterceptorsTrieNodes := (noOfShards + 1) * 2
	totalInterceptors := numInterceptorsMetablock + numInterceptorsShardHeadersForMetachain + numInterceptorsTrieNodes +
		numInterceptorsTransactionsForMetachain + numInterceptorsUnsignedTxsForMetachain + numInterceptorsMiniBlocksForMetachain +
		numInterceptorsRewardsTxsForMetachain

	assert.Nil(t, err)
	assert.Equal(t, totalInterceptors, container.Len())
}

func getArgumentsMeta() interceptorscontainer.MetaInterceptorsContainerFactoryArgs {
	return interceptorscontainer.MetaInterceptorsContainerFactoryArgs{
		ShardCoordinator:       mock.NewOneShardCoordinatorMock(),
		NodesCoordinator:       mock.NewNodesCoordinatorMock(),
		Messenger:              &mock.TopicHandlerStub{},
		Store:                  createMetaStore(),
		ProtoMarshalizer:       &mock.MarshalizerMock{},
		TxSignMarshalizer:      &mock.MarshalizerMock{},
		Hasher:                 &mock.HasherMock{},
		MultiSigner:            mock.NewMultiSigner(),
		DataPool:               createMetaDataPools(),
		Accounts:               &mock.AccountsStub{},
		AddrConverter:          &mock.AddressConverterMock{},
		SingleSigner:           &mock.SignerMock{},
		BlockSingleSigner:      &mock.SignerMock{},
		KeyGen:                 &mock.SingleSignKeyGenMock{},
		BlockKeyGen:            &mock.SingleSignKeyGenMock{},
		MaxTxNonceDeltaAllowed: maxTxNonceDeltaAllowed,
		TxFeeHandler:           &mock.FeeHandlerStub{},
		BlackList:              &mock.BlackListHandlerStub{},
		HeaderSigVerifier:      &mock.HeaderSigVerifierStub{},
		ChainID:                chainID,
		SizeCheckDelta:         0,
		ValidityAttester:       &mock.ValidityAttesterStub{},
		EpochStartTrigger:      &mock.EpochStartTriggerStub{},
	}
}