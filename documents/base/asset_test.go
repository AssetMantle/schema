package base

import (
	"github.com/AssetMantle/schema/types"
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/schema/data/base"
	documentsSchema "github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

func TestNewAsset(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want documentsSchema.Asset
	}{
		{"+ve", args{classificationID: classificationID, immutables: immutables, mutables: mutables}, asset{Document: NewDocument(classificationID, immutables, mutables)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAsset(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetBurnHeight(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithBurn := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMetaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(99))))))

	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Height
	}{
		{"+ve", fields{Document: testDocumentWithBurn}, baseTypes.NewHeight(99)},
		{"+ve", fields{Document: testDocument}, baseTypes.NewHeight(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetBurnHeight(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBurnHeight() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_asset_GetLockHeight(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithLock := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMetaProperty(constants.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(100))))))

	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Height
	}{
		{"+ve with default lock", fields{testDocument}, baseTypes.NewHeight(0)},
		{"+ve with mutated", fields{testDocumentWithLock}, baseTypes.NewHeight(100)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetLockHeight(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLockHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetSupply(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithSupply := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMetaProperty(constants.SupplyProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(77))))))
	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   math.Int
	}{
		{"+ve", fields{testDocument}, sdkTypes.OneInt()},
		{"+ve with supply", fields{testDocumentWithSupply}, sdkTypes.NewInt(77)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := NewAssetFromDocument(tt.fields.Document)
			if got := asset.GetSupply(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}
