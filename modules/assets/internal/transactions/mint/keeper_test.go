// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	conformAuxiliary           helpers.Auxiliary
	mintAuxiliary              helpers.Auxiliary
	authenticateAuxiliary      helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
)

type TestKeepers struct {
	MintKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	ParamsKeeper := paramsKeeper.NewKeeper(
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	mintAuxiliary = mint.AuxiliaryMock.Initialize(Mapper, Parameters)
	conformAuxiliary = conform.AuxiliaryMock.Initialize(Mapper, Parameters)
	maintainersVerifyAuxiliary = verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)

	keepers := TestKeepers{
		MintKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, Parameters
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		{"+ve", transactionKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	_, _, Mapper, Parameters := createTestInput(t)
	type fields struct {
		mapper                     helpers.Mapper
		parameters                 helpers.Parameters
		conformAuxiliary           helpers.Auxiliary
		mintAuxiliary              helpers.Auxiliary
		authenticateAuxiliary      helpers.Auxiliary
		maintainersVerifyAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		parameters  helpers.Parameters
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve", fields{Mapper, Parameters, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, Parameters, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                     tt.fields.mapper,
				parameters:                 tt.fields.parameters,
				conformAuxiliary:           tt.fields.conformAuxiliary,
				mintAuxiliary:              tt.fields.mintAuxiliary,
				authenticateAuxiliary:      tt.fields.authenticateAuxiliary,
				maintainersVerifyAuxiliary: tt.fields.maintainersVerifyAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameters, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper, Parameters := createTestInput(t)
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testAsset := base.NewAsset(classificationID, immutables, mutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromID := baseIDs.NewIdentityID(classificationID, immutables)
	keepers.MintKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(testAsset))

	type fields struct {
		mapper                     helpers.Mapper
		parameters                 helpers.Parameters
		conformAuxiliary           helpers.Auxiliary
		mintAuxiliary              helpers.Auxiliary
		authenticateAuxiliary      helpers.Auxiliary
		maintainersVerifyAuxiliary helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		msg     sdkTypes.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		{"+ve", fields{Mapper, Parameters, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, fromID, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, newTransactionResponse(nil)},
		{"+ve Entity Already Exists", fields{Mapper, Parameters, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, fromID, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, newTransactionResponse(errorConstants.EntityAlreadyExists)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                     tt.fields.mapper,
				parameters:                 tt.fields.parameters,
				conformAuxiliary:           tt.fields.conformAuxiliary,
				mintAuxiliary:              tt.fields.mintAuxiliary,
				authenticateAuxiliary:      tt.fields.authenticateAuxiliary,
				maintainersVerifyAuxiliary: tt.fields.maintainersVerifyAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
