// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data"
	dataConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

var fromAddress = "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"

var accAddress = NewAccAddressData(sdkTypes.AccAddress(fromAddress)).AsString()

func TestListDataValidateBasic(t *testing.T) {
	type args struct {
		value data.ListData
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"+ve", args{NewListData(NewStringData("Data"))}, true},
	}
	for _, tt := range tests {
		if err := tt.args.value.ValidateBasic(); (err != nil) != tt.want {
			t.Errorf("got = %v, want = %v", err, tt.want)
		}
	}
}
func TestListDataPrototype(t *testing.T) {
	type args struct {
		value data.ListData
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve", args{}, &ListData{[]*AnyListableData(nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeListData(), "Prototype(%v)", tt.args.value)
		})
	}
}

func TestNewListData(t *testing.T) {
	type args struct {
		value data.ListableData
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve for some id", args{NewStringData("Data")}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}},
		{"+ve for empty String", args{NewStringData("")}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}},

		// {"+ve empty datalist", args{data.Data()}, (&ListData{}).ZeroValue()},
		{"+ve address string", args{NewStringData(fromAddress)}, &ListData{[]*AnyListableData{NewStringData(fromAddress).ToAnyListableData().(*AnyListableData)}}},
		// TODO: Check address format
		// {"-ve wrong address string format", args{NewListData(NewStringData(fromAddress))}, &ListData{}.ZeroValue()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListData(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_Add(t *testing.T) {
	type fields struct {
		Value []data.ListableData
	}

	tests := []struct {
		name        string
		fields      fields
		want        data.ListData
		wantFailure bool
	}{
		{"+ve for multiple ids", fields{[]data.ListableData{NewStringData("Data"), NewStringData("Data"), NewStringData("Data")}}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}, false},
		{"+ve for multiple ids/nils", fields{[]data.ListableData{NewStringData("Data"), NewStringData(""), NewStringData("Data")}}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData), NewStringData("Data").ToAnyListableData().(*AnyListableData)}}, false},
		{"+ve for some id", fields{[]data.ListableData{NewStringData("Data")}}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}, false},
		{"+ve for empty String", fields{[]data.ListableData{NewStringData("")}}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}, false},
		{"-ve for value inequality", fields{[]data.ListableData{NewStringData("Data")}}, &ListData{[]*AnyListableData{NewStringData("Data1").ToAnyListableData().(*AnyListableData)}}, true},
		{"-ve for occurrence inequality", fields{[]data.ListableData{NewStringData("Data"), NewStringData("Data"), NewStringData("Data")}}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData), NewStringData("Data").ToAnyListableData().(*AnyListableData)}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := &ListData{}

			if got := listData.Add(tt.fields.Value...); reflect.DeepEqual(got, tt.want) != !tt.wantFailure {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_Bytes(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, NewStringData("Data").Bytes()}, // for a single data no loop iteration is required, so directly it's byte should match
		{"+ve for multiple ids", fields{NewListData(NewStringData("Data"), NewStringData("Data1"))}, bytes.Join([][]byte{NewStringData("Data").Bytes(), NewStringData("Data1").Bytes()}, nil)},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, []byte(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.Bytes(), "Bytes()")
		})
	}
}

func Test_listData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, baseIDs.GenerateHashID((&ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}).Bytes()).AsString()},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, baseIDs.GenerateHashID((&ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}).Bytes()).AsString()},
		{"empty string", fields{NewListData()}, baseIDs.NewStringID("").AsString()},
		{"+ve case", fields{NewListData(NewStringData(accAddress))}, baseIDs.NewStringID("xrHmURH4R458qdPeDW8kU9eO3a3bvQRE0W6CAoZ8yCw=").AsString()},
		{"-ve case", fields{NewListData(NewStringData(""))}, baseIDs.NewStringID("").AsString()},
		{"-ve case with empty datalist", fields{NewListData([]data.ListableData{}...)}, baseIDs.NewStringID("").AsString()},
		{"-ve case with nil data", fields{NewListData()}, baseIDs.NewStringID("").AsString()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.GenerateHashID().AsString(), "GenerateHashID()")
		})
	}
}

func Test_listData_Get(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   []data.ListableData
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, (&ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}).Get()},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, (&ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}).Get()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.Get(), "Get()")
		})
	}
}

func Test_listData_GetID(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, baseIDs.GenerateDataID(&ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}})},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, baseIDs.GenerateDataID(&ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.GetID(), "GetID()")
		})
	}
}

func Test_listData_GetType(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, dataConstants.ListDataTypeID},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, dataConstants.ListDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.GetTypeID(), "GetTypeID()")
		})
	}
}

func Test_listData_Remove(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	type args struct {
		listableData []data.ListableData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   data.ListData
	}{
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, args{[]data.ListableData{}}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}},
		{"+ve for empty String & removing it", fields{NewListData(NewStringData(""))}, args{[]data.ListableData{NewStringData("")}}, &ListData{[]*AnyListableData{}}},
		{"+ve ", fields{NewListData(NewStringData("data"))}, args{[]data.ListableData{NewStringData("data")}}, &ListData{[]*AnyListableData{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.Remove(tt.args.listableData...), "Remove(%v)", tt.args.listableData)
		})
	}
}

func Test_listData_Search(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	type args struct {
		listableData data.ListableData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, args{NewStringData("Data")}, 0, true},
		{"+ve for empty String", fields{NewListData([]data.ListableData{NewStringData("Data"), NewStringData("")}...)}, args{NewStringData("")}, 0, true},
		{"-ve", fields{NewListData([]data.ListableData{NewStringData("Data"), NewStringData("")}...)}, args{NewStringData("test")}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			got, got1 := listData.Search(tt.args.listableData)
			assert.Equalf(t, tt.want, got, "search(%v)", tt.args.listableData)
			assert.Equalf(t, tt.want1, got1, "search(%v)", tt.args.listableData)
		})
	}
}

func Test_listData_AsString(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, "Data"},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.AsString(), "String()")
		})
	}
}

func Test_listData_ZeroValue(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, NewListData([]data.ListableData{}...)},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, NewListData([]data.ListableData{}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.ZeroValue(), "ZeroValue()")
		})
	}
}

func TestListData_ValidateWithType(t *testing.T) {
	type fields struct {
		AnyListableDataList []*AnyListableData
	}
	type args struct {
		expectedTypeID ids.StringID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve for some id", fields{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}, args{baseIDs.NewStringID("string")}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := &ListData{
				tt.fields.AnyListableDataList,
			}
			tt.wantErr(t, listData.ValidateWithType(tt.args.expectedTypeID), fmt.Sprintf("ValidateWithType(%v)", tt.args.expectedTypeID))
		})
	}
}
