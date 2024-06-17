package base

import (
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
)

func Test_new_CoinAsset(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("denom"), baseData.NewStringData("denom1"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList())
	coinAsset := asset{Document: NewDocument(coinAssetClassificationID, immutables, mutables)}
	tests := []struct {
		name      string
		fields    string
		want      documents.CoinAsset
		wantError bool
	}{
		{"+ve", "denom1", coinAsset, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantError {
					t.Errorf("error = %v, wantErr %v", r, tt.wantError)
				}
			}()
			if got := NewCoinAsset(tt.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClassification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinAsset_GetDenom(t *testing.T) {
	tests := []struct {
		name   string
		fields documents.CoinAsset
		want   string
	}{
		{"+ve", NewCoinAsset("denom1"), "denom1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				fmt.Println(r)
			}()
			if got := tt.fields.GetDenom(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("coinAsset_GetDenom() = %v, want %v", got, tt.want)
			}
		})
	}
}
