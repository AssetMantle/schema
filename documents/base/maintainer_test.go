package base

import (
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	baseProperties "github.com/AssetMantle/schema/properties/base"
)

func TestNewMaintainer(t *testing.T) {
	classificationID, _, _, testDocument := createTestInput()
	testIdentity := NewIdentityFromDocument(testDocument)
	identityID := baseIDs.NewIdentityID(testIdentity.GetClassificationID(), testIdentity.GetImmutables())
	maintainerProperties := baseLists.NewIDList(baseIDs.NewStringID("a"), baseIDs.NewStringID("b"))
	permissions := baseLists.NewIDList(baseIDs.NewStringID("c"), baseIDs.NewStringID("d"))
	maintainer := NewDocument(maintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(identityID)),
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(classificationID)),
		)),
		baseQualified.NewMutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedPropertiesProperty.GetKey(), idListToListData(maintainerProperties)),
			baseProperties.NewMetaProperty(constantProperties.PermissionsProperty.GetKey(), idListToListData(permissions)),
		)),
	)
	type args struct {
		identityID           ids.IdentityID
		classificationID     ids.ClassificationID
		maintainerProperties lists.IDList
		permissions          lists.IDList
	}
	tests := []struct {
		name string
		args args
		want documents.Maintainer
	}{
		{"+ve", args{identityID, classificationID, maintainerProperties, permissions}, NewMaintainerFromDocument(maintainer)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaintainer(tt.args.identityID, tt.args.classificationID, tt.args.maintainerProperties, tt.args.permissions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_identity_GetExpiry(t *testing.T) {
//	testImmutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
//	testMutables1 := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData().ZeroValue()), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiry"), baseData.NewHeightData(baseTypes.NewHeight(10))), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiryHeight"), baseData.NewHeightData(baseTypes.NewHeight(100)))))
//	testMutables2 := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData().ZeroValue()), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiry"), baseData.NewHeightData(baseTypes.NewHeight(10)))))
//	classificationID := baseIDs.NewClassificationID(testImmutables, testMutables1)
//	testDocument1 := NewDocument(classificationID, testImmutables, testMutables1)
//	testIdentity1 := identity{testDocument1}
//	testDocument2 := NewDocument(classificationID, testImmutables, testMutables2)
//	testIdentity2 := identity{testDocument2}
//	type fields struct {
//		Document documents.Identity
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   types.Height
//	}{
//		{"+ve", fields{testIdentity1}, baseTypes.NewHeight(100)},
//		{"+ve", fields{testIdentity2}, baseTypes.NewHeight(-1)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			identity := tt.fields.Document
//
//			if got := identity.GetExpiry(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetExpiry() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_identity_GetAuthentication(t *testing.T) {
//	classificationID, immutables, mutables, _ := createTestInput()
//	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
//	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
//	require.Nil(t, err)
//	type fields struct {
//		Document documents.Identity
//	}
//
//	tests := []struct {
//		name   string
//		fields documents.Identity
//		want   data.ListData
//	}{
//		{"+ve", NewIdentityFromDocument(NewDocument(classificationID, immutables, mutables)).ProvisionAddress(fromAccAddress), baseData.NewListData(baseData.NewAccAddressData(fromAccAddress))},
//		{"+ve", identity{NewDocument(classificationID, immutables, baseQualified.NewMutables(baseLists.NewPropertyList()))}, constants.AuthenticationProperty.GetData().Get().(data.ListData)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			if got := tt.fields.GetAuthentication(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAuthentication() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
