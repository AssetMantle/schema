package base

import (
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/schema/go/data/base"
	documentsSchema "github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/qualified"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
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

func Test_asset_GetBurn(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithBurn := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))))

	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		{"+ve", fields{Document: testDocumentWithBurn}, baseProperties.NewMesaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))).ToAnyProperty()},
		{"+ve", fields{Document: testDocument}, constants.BurnHeightProperty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetBurn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBurn() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_asset_GetLock(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithLock := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))))

	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		{"+ve with default lock", fields{testDocument}, constants.LockProperty},
		{"+ve with mutated", fields{testDocumentWithLock}, baseProperties.NewMesaProperty(constants.LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))).ToAnyProperty()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetLock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetSupply(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithSupply := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.SupplyProperty.GetKey(), baseData.NewDecData(sdkTypes.NewDec(1))))))
	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.AnyProperty
	}{
		{"+ve", fields{testDocument}, constants.SupplyProperty.ToAnyProperty()},
		{"+ve with supply", fields{testDocumentWithSupply}, baseProperties.NewMesaProperty(constants.SupplyProperty.GetKey(), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetSupply().ToAnyProperty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}