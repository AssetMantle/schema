// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified"
	"github.com/AssetMantle/schema/types"
)

func TestNewIdentity(t *testing.T) {
	classificationID, immutables, mutables, testDocument := createTestInput()
	testIdentity := identity{testDocument}
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want documents.Identity
	}{
		{"+ve", args{classificationID, immutables, mutables}, testIdentity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdentity(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_IsProvisioned(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromAccAddress2, err := sdkTypes.AccAddressFromBech32("cosmos1u6xn6rv07p2yzzj2rm8st04x54xe5ur0t9nl5j")
	require.Nil(t, err)
	testIdentity := NewIdentity(classificationID, immutables, mutables)
	testIdentity.ProvisionAddress(fromAccAddress)

	type fields struct {
		Document documents.Identity
	}
	type args struct {
		accAddress sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"+ve", fields{testIdentity}, args{fromAccAddress}, true},
		{"-ve", fields{testIdentity}, args{fromAccAddress2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := tt.fields.Document
			if got := identity.IsProvisioned(tt.args.accAddress); got != tt.want {
				t.Errorf("IsProvisioned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_GetExpiry(t *testing.T) {
	testImmutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	testMutables1 := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData().ZeroValue()), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiry"), baseData.NewHeightData(baseTypes.NewHeight(10))), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiryHeight"), baseData.NewHeightData(baseTypes.NewHeight(100)))))
	testMutables2 := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData().ZeroValue()), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiry"), baseData.NewHeightData(baseTypes.NewHeight(10)))))
	classificationID := baseIDs.NewClassificationID(testImmutables, testMutables1)
	testDocument1 := NewDocument(classificationID, testImmutables, testMutables1)
	testIdentity1 := identity{testDocument1}
	testDocument2 := NewDocument(classificationID, testImmutables, testMutables2)
	testIdentity2 := identity{testDocument2}
	type fields struct {
		Document documents.Identity
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Height
	}{
		{"+ve", fields{testIdentity1}, baseTypes.NewHeight(100)},
		{"+ve", fields{testIdentity2}, baseTypes.NewHeight(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := tt.fields.Document

			if got := identity.GetExpiry(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpiry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_ProvisionAddress(t *testing.T) {
	classificationID, immutables, mutables, testDocument := createTestInput()
	testIdentity := identity{testDocument}
	testIdentity2 := testIdentity
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromAccAddress2, err := sdkTypes.AccAddressFromBech32("cosmos1u6xn6rv07p2yzzj2rm8st04x54xe5ur0t9nl5j")
	require.Nil(t, err)
	testIdentity.Document = testIdentity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), testIdentity.GetAuthentication().Add(baseData.NewAccAddressData(fromAccAddress))))
	fmt.Println("TEST:	", testIdentity.IsProvisioned(fromAccAddress))
	type fields struct {
		Document documents.Identity
	}
	type args struct {
		accAddresses []sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{[]sdkTypes.AccAddress{fromAccAddress}}, testIdentity.IsProvisioned(fromAccAddress)},
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{[]sdkTypes.AccAddress{fromAccAddress2}}, testIdentity2.IsProvisioned(fromAccAddress)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.ProvisionAddress(tt.args.accAddresses...).IsProvisioned(tt.args.accAddresses[0]); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProvisionAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_GetAuthentication(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	tests := []struct {
		name   string
		fields documents.Identity
		want   data.ListData
	}{
		{"+ve", NewIdentityFromDocument(NewDocument(classificationID, immutables, mutables)).ProvisionAddress(fromAccAddress), baseData.NewListData(baseData.NewAccAddressData(fromAccAddress))},
		{"+ve", identity{NewDocument(classificationID, immutables, baseQualified.NewMutables(baseLists.NewPropertyList()))}, constants.AuthenticationProperty.GetData().Get().(data.ListData)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.fields.GetAuthentication(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_UnprovisionAddress(t *testing.T) {
	classificationID, immutables, mutables, testDocument := createTestInput()
	testIdentity := identity{testDocument}
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	type fields struct {
		Document documents.Identity
	}
	type args struct {
		accAddresses []sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   documents.Identity
	}{
		// TODO: panic: MetaDataError fix it after
		// https://github.com/AssetMantle/schema/issues/59
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{[]sdkTypes.AccAddress{fromAccAddress}}, testIdentity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.UnprovisionAddress(tt.args.accAddresses...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnprovisionAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IdentityValidateBasic(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	testMutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseData.NewStringData("a")))))
	testIdentity1 := NewIdentity(classificationID, immutables, mutables).ProvisionAddress(fromAccAddress)
	testIdentity2 := NewIdentity(classificationID, immutables, testMutables)
	tests := []struct {
		name string
		args documents.Identity
		want bool
	}{
		{"+ve 1", testIdentity1, false},
		{"-ve 2", testIdentity2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err != nil {
				fmt.Println(err.Error())
			}
			if err == nil && tt.want {
				t.Errorf("IdentityValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("IdentityValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}
