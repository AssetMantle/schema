// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
)

var (
	idData                   = (&baseIDs.HashID{[]byte("a")}).ToAnyID().(*baseIDs.AnyID)
	testPropertyList         = &PropertyList{propertiesToAnyProperties([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewNumberData(sdkTypes.NewInt(1)))}...)}
	invalidTestPropertyList1 = &PropertyList{propertiesToAnyProperties([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("a"))}...)}
	invalidTestPropertyList2 = &PropertyList{propertiesToAnyProperties([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), &baseData.IDData{idData}), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("a"))}...)}
)

type fields struct {
	List []properties.Property
}

func TestNewPropertyList(t *testing.T) {
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name string
		args args
		want lists.PropertyList
	}{
		{"+ve", args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, &PropertyList{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")).ToAnyProperty().(*baseProperties.AnyProperty), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPropertyList(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Add(t *testing.T) {
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{
			[]properties.Property{
				baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))}}, args{[]properties.Property{

			baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, NewPropertyList([]properties.Property{

			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")),
			baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)},
		{"+ve", fields{[]properties.Property{
			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("1"))}}, args{[]properties.Property{

			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("2"))}}, NewPropertyList([]properties.Property{

			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("2"))}...)},
		{"-ve", fields{[]properties.Property{
			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))}}, args{[]properties.Property{

			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, NewPropertyList([]properties.Property{

			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")),
			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)},
		{"+ve with nil", fields{[]properties.Property{
			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))}}, args{

			nil}, NewPropertyList([]properties.Property{

			baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))}...)}, // TODO: Should fail as add method should not be able to mutate/add property with existing key
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.fields.List...)}
			if got := propertyList.Add(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			} else if !reflect.DeepEqual(propertyList, PropertyList{propertiesToAnyProperties(tt.fields.List...)}) { // check if the original list is mutated
				t.Errorf("Receiver = %v, Receiver after calling method %v", propertyList, PropertyList{propertiesToAnyProperties(tt.fields.List...)})
			}
		})
	}
}

func Test_propertyList_GetList(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   []properties.Property
	}{
		{"+ve", fields{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.fields.List...)}
			for i, got := range propertyList.Get() {
				if got.Compare(tt.want[i]) != 0 {
					t.Errorf("GetList() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}

func Test_propertyList_GetProperty(t *testing.T) {

	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    properties.AnyProperty
		wantErr bool
	}{
		{"+ve Meta", fields{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, args{baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))}, baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), false},
		{"+ve Mesa", fields{[]properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, args{baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))}, baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), false},
		{"panic nil propertyID", fields{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, args{propertyID: nil}, baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), true}, // TODO: panics if propertyID is nil
		{"-ve", fields{[]properties.Property{baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D")))}}, args{baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))}, baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))).ToAnyProperty(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.fields.List...)}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("&StringIDFromInterface() error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := propertyList.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetPropertyIDList(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   lists.IDList
	}{
		{"+ve", fields{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, NewIDList().Add(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).GetID())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.fields.List...)}
			if got := propertyList.GetPropertyIDList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPropertyIDList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Mutate(t *testing.T) {

	propertyList1 := &PropertyList{propertiesToAnyProperties([]properties.Property{
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewNumberData(sdkTypes.OneInt())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID3"), baseData.NewHeightData(baseTypes.NewHeight(10))),
	}...)}
	propertyListMutated := &PropertyList{propertiesToAnyProperties([]properties.Property{
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("DataUpdated")),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewNumberData(sdkTypes.OneInt())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID3"), baseData.NewHeightData(baseTypes.NewHeight(10))),
	}...)}
	tests := []struct {
		name   string
		fields lists.PropertyList
		args   []properties.Property
		want   lists.PropertyList
	}{
		{"+ve", propertyList1, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("DataUpdated"))}, propertyListMutated},
		{"+ve", propertyList1, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("DataUpdated"))}, propertyList1},
		{"+ve", propertyList1, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewNumberData(sdkTypes.OneInt()))}, propertyList1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Mutate(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("propertyList_Mutate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Remove(t *testing.T) {
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply1"), baseData.NewDecData(sdkTypes.NewDec(2)))}}, args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply1"), baseData.NewDecData(sdkTypes.NewDec(2)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.fields.List...)}
			if got := propertyList.Remove(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_ScrubData(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		{"+ve", fields{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ScrubData()}...)},
		{"+ve", fields{[]properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Test"))}}, NewPropertyList([]properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Test")).ScrubData()}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.fields.List...)}
			if got := propertyList.ScrubData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrubData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PropertyListValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args lists.PropertyList
		want bool
	}{
		{"+ve", testPropertyList, false},
		{"-ve", invalidTestPropertyList1, true},
		{"-ve", invalidTestPropertyList2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err != nil {
				fmt.Println(err.Error())
			}
			if err == nil && tt.want {
				t.Errorf("PropertyListValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("PropertyListValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_propertyList_FromMetaPropertiesString(t *testing.T) {

	propertyList1 := &PropertyList{propertiesToAnyProperties([]properties.Property{
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewNumberData(sdkTypes.OneInt())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID3"), baseData.NewHeightData(baseTypes.NewHeight(10))),
	}...)}
	propertyList2 := &PropertyList{propertiesToAnyProperties([]properties.Property{
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewNumberData(sdkTypes.OneInt())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID3"), baseData.NewHeightData(baseTypes.NewHeight(10))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID4"), baseData.NewListData(baseData.NewStringData("a"))),
	}...)}
	propertyList3 := &PropertyList{propertiesToAnyProperties([]properties.Property{
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewNumberData(sdkTypes.OneInt())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID3"), baseData.NewHeightData(baseTypes.NewHeight(10))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID4"), baseData.NewListData(baseData.NewStringData("a"), baseData.NewStringData("b"))),
	}...)}
	tests := []struct {
		name    string
		args    string
		want    lists.PropertyList
		wantErr bool
	}{
		{"-ve", "ID1:S|Data1,,ID2:N|1,,ID3:H10", nil, true},
		{"+ve", "ID1:S|Data1,,ID2:N|1,,ID3:H|10", propertyList1, false},
		{"-ve", "ID1:S|Data1,,ID2:N|1,,ID3:H|10,,ID3:H|12", nil, true},
		{"+ve", "ID1:S|Data1,,ID2:N|1,,ID3:H|10,,ID4:L|S|a", propertyList2, false},
		{"+ve", "ID1:S|Data1,,ID2:N|1,,ID3:H|10,,ID4:L|S|a,,S|b", propertyList3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypePropertyList().FromMetaPropertiesString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromMetaPropertiesString() got = %v, want %v", got, tt.want)
			}
			if err != nil {
				fmt.Println(err.Error())
			}
			if err == nil && tt.wantErr {
				t.Errorf("FromMetaPropertiesString() Error got = %v, want %v", err, tt.wantErr)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("FromMetaPropertiesString() Error got = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
