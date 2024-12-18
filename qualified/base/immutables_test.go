// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/qualified"
)

func TestNewImmutables(t *testing.T) {
	testImmutablePropertyList := base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("MutableData1")))
	type args struct {
		propertyList lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want qualified.Immutables
	}{
		{"nil", args{}, &Immutables{PropertyList: &base.PropertyList{}}},
		{"+ve empty list", args{base.NewPropertyList()}, &Immutables{PropertyList: &base.PropertyList{}}},
		{"+ve", args{testImmutablePropertyList}, &Immutables{testImmutablePropertyList.(*base.PropertyList)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImmutables(tt.args.propertyList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImmutables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_immutables_GenerateHashID(t *testing.T) {
	testMesaProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("ImmutableData"))
	testImmutablePropertyList := base.NewPropertyList(testMesaProperty)
	metaList2 := make([][]byte, len(testImmutablePropertyList.Get()))
	metaList2[0] = testMesaProperty.GetDataID().GetHashID().Bytes()
	type fields struct {
		PropertyList lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		{"+ve", fields{base.NewPropertyList()}, baseIDs.GenerateHashID([][]byte{}...)},
		{"+ve", fields{testImmutablePropertyList}, baseIDs.GenerateHashID(metaList2...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testImmutables := Immutables{
				PropertyList: tt.fields.PropertyList.(*base.PropertyList),
			}
			if got := testImmutables.GenerateHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_immutables_GetImmutablePropertyList(t *testing.T) {
	testImmutablePropertyList := base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("MutableData1")))
	type fields struct {
		PropertyList lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		{"+ve with nil", fields{base.NewPropertyList()}, base.NewPropertyList()},
		{"+ve", fields{testImmutablePropertyList}, Immutables{testImmutablePropertyList.(*base.PropertyList)}.PropertyList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testImmutables := Immutables{
				PropertyList: tt.fields.PropertyList.(*base.PropertyList),
			}
			if got := testImmutables.GetImmutablePropertyList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImmutablePropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}
