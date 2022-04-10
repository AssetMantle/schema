// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type TestKeepers struct {
	ClassificationsKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		ClassificationsKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers
}

func Test_Auxiliary_Keeper_Help(t *testing.T) {
	context, keepers := CreateTestInput(t)

	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("Data2")))
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")))

	classificationID := key.NewClassificationID(baseIDs.NewID(context.ChainID()), immutableProperties, mutableProperties)

	testClassificationID := key.NewClassificationID(baseIDs.NewID(context.ChainID()), baseTypes.NewProperties(), baseTypes.NewProperties())

	keepers.ClassificationsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewClassification(testClassificationID, baseTypes.NewProperties(), baseTypes.NewProperties()))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(baseIDs.NewID(classificationID.String()), nil)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(immutableProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Classification already present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(baseIDs.NewID(testClassificationID.String()), errors.EntityAlreadyExists)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseTypes.NewProperties(), baseTypes.NewProperties())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Max Property Count", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")), baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("Data2")), baseTypes.NewProperty(baseIDs.NewID("ID3"), baseData.NewStringData("Data3")), baseTypes.NewProperty(baseIDs.NewID("ID4"), baseData.NewStringData("Data4")), baseTypes.NewProperty(baseIDs.NewID("ID5"), baseData.NewStringData("Data5")), baseTypes.NewProperty(baseIDs.NewID("ID6"), baseData.NewStringData("Data6")), baseTypes.NewProperty(baseIDs.NewID("ID7"), baseData.NewStringData("Data7")), baseTypes.NewProperty(baseIDs.NewID("ID8"), baseData.NewStringData("Data8")), baseTypes.NewProperty(baseIDs.NewID("ID9"), baseData.NewStringData("Data9")), baseTypes.NewProperty(baseIDs.NewID("ID10"), baseData.NewStringData("Data10")), baseTypes.NewProperty(baseIDs.NewID("ID9"), baseData.NewStringData("Data9")), baseTypes.NewProperty(baseIDs.NewID("ID10"), baseData.NewStringData("Data10"))), baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")), baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("Data2")), baseTypes.NewProperty(baseIDs.NewID("ID3"), baseData.NewStringData("Data3")), baseTypes.NewProperty(baseIDs.NewID("ID4"), baseData.NewStringData("Data4")), baseTypes.NewProperty(baseIDs.NewID("ID5"), baseData.NewStringData("Data5")), baseTypes.NewProperty(baseIDs.NewID("ID6"), baseData.NewStringData("Data6")), baseTypes.NewProperty(baseIDs.NewID("ID7"), baseData.NewStringData("Data7")), baseTypes.NewProperty(baseIDs.NewID("ID8"), baseData.NewStringData("Data8")), baseTypes.NewProperty(baseIDs.NewID("ID9"), baseData.NewStringData("Data9")), baseTypes.NewProperty(baseIDs.NewID("ID10"), baseData.NewStringData("Data10")), baseTypes.NewProperty(baseIDs.NewID("ID9"), baseData.NewStringData("Data9")), baseTypes.NewProperty(baseIDs.NewID("ID10"), baseData.NewStringData("Data10"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")), baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data2")), baseTypes.NewProperty(baseIDs.NewID("ID3"), baseData.NewStringData("Data3")), baseTypes.NewProperty(baseIDs.NewID("ID4"), baseData.NewStringData("Data4")), baseTypes.NewProperty(baseIDs.NewID("ID5"), baseData.NewStringData("Data5")), baseTypes.NewProperty(baseIDs.NewID("ID6"), baseData.NewStringData("Data6")), baseTypes.NewProperty(baseIDs.NewID("ID7"), baseData.NewStringData("Data7")), baseTypes.NewProperty(baseIDs.NewID("ID8"), baseData.NewStringData("Data8")), baseTypes.NewProperty(baseIDs.NewID("ID9"), baseData.NewStringData("Data9")), baseTypes.NewProperty(baseIDs.NewID("ID10"), baseData.NewStringData("Data10"))), baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID11"), baseData.NewStringData("Data11")), baseTypes.NewProperty(baseIDs.NewID("ID12"), baseData.NewStringData("Data12")), baseTypes.NewProperty(baseIDs.NewID("ID13"), baseData.NewStringData("Data13")), baseTypes.NewProperty(baseIDs.NewID("ID14"), baseData.NewStringData("Data14")), baseTypes.NewProperty(baseIDs.NewID("ID15"), baseData.NewStringData("Data15")), baseTypes.NewProperty(baseIDs.NewID("ID16"), baseData.NewStringData("Data16")), baseTypes.NewProperty(baseIDs.NewID("ID17"), baseData.NewStringData("Data17")), baseTypes.NewProperty(baseIDs.NewID("ID18"), baseData.NewStringData("Data18")), baseTypes.NewProperty(baseIDs.NewID("ID19"), baseData.NewStringData("Data19")), baseTypes.NewProperty(baseIDs.NewID("ID20"), baseData.NewStringData("Data20"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable and Mutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")), baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("Data2")), baseTypes.NewProperty(baseIDs.NewID("ID3"), baseData.NewStringData("Data3")), baseTypes.NewProperty(baseIDs.NewID("ID4"), baseData.NewStringData("Data4")), baseTypes.NewProperty(baseIDs.NewID("ID5"), baseData.NewStringData("Data5")), baseTypes.NewProperty(baseIDs.NewID("ID6"), baseData.NewStringData("Data6")), baseTypes.NewProperty(baseIDs.NewID("ID7"), baseData.NewStringData("Data7")), baseTypes.NewProperty(baseIDs.NewID("ID8"), baseData.NewStringData("Data8")), baseTypes.NewProperty(baseIDs.NewID("ID9"), baseData.NewStringData("Data9")), baseTypes.NewProperty(baseIDs.NewID("ID10"), baseData.NewStringData("Data10"))), baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data11")), baseTypes.NewProperty(baseIDs.NewID("ID12"), baseData.NewStringData("Data12")), baseTypes.NewProperty(baseIDs.NewID("ID13"), baseData.NewStringData("Data13")), baseTypes.NewProperty(baseIDs.NewID("ID14"), baseData.NewStringData("Data14")), baseTypes.NewProperty(baseIDs.NewID("ID15"), baseData.NewStringData("Data15")), baseTypes.NewProperty(baseIDs.NewID("ID16"), baseData.NewStringData("Data16")), baseTypes.NewProperty(baseIDs.NewID("ID17"), baseData.NewStringData("Data17")), baseTypes.NewProperty(baseIDs.NewID("ID18"), baseData.NewStringData("Data18")), baseTypes.NewProperty(baseIDs.NewID("ID19"), baseData.NewStringData("Data19")), baseTypes.NewProperty(baseIDs.NewID("ID20"), baseData.NewStringData("Data20"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
