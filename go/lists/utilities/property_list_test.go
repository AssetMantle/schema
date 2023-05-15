// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/lists/base"
	propertiesSchema "github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
)

func TestReadMetaProperties(t *testing.T) {
	type args struct {
		metaPropertiesString string
	}
	tests := []struct {
		name    string
		args    args
		want    lists.PropertyList
		wantErr bool
	}{
		{"+ve with empty string", args{""}, base.NewPropertyList([]propertiesSchema.Property{}...), false},
		{"+ve", args{"ID:S|Data,ID1:S|Data1,ID2:S|Data2"}, base.NewPropertyList([]propertiesSchema.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Data")), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))}...), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadMetaPropertyList(tt.args.metaPropertiesString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMetaProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range got.Get() {
				if got.Get()[i].Compare(tt.want.Get()[i]) != 0 {
					t.Errorf("ReadMetaProperties() got = %v, want %v", got.Get()[i], tt.want.Get()[i])
				}
			}
		})
	}
}
