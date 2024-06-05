// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/AssetMantle/schema/data"
	idsConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

func Test_NewBooleanData(t *testing.T) {
	tests := []struct {
		name string
		args bool
		want data.Data
	}{
		{"+ve", true, &BooleanData{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBooleanData(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBooleanData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want bool
	}{
		{"+ve", NewBooleanData(true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != (got != nil) {
				t.Errorf("BooleanDataValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want bool
	}{
		{"+ve", NewBooleanData(true), true},
		{"-ve", NewBooleanData(false), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(NewBooleanData(true))
			if (got == 0) != tt.want {
				t.Errorf("BooleanData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want ids.ID
	}{
		{"+ve", NewBooleanData(true), baseIDs.GenerateHashID([]byte{0x1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataGet(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want bool
	}{
		{"+ve", NewBooleanData(true), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanDataGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want ids.DataID
	}{
		{"+ve", NewBooleanData(true), &baseIDs.DataID{
			TypeID: idsConstants.BooleanDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID([]byte{0x1}).(*baseIDs.HashID),
		}},
		{"+ve", &BooleanData{false}, &baseIDs.DataID{
			TypeID: idsConstants.BooleanDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID().(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want ids.ID
	}{
		{"+ve", NewBooleanData(true), idsConstants.BooleanDataTypeID},
		{"+ve", &BooleanData{false}, idsConstants.BooleanDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want string
	}{
		{"+ve", NewBooleanData(true), strconv.FormatBool(true)},
		{"+ve", &BooleanData{false}, strconv.FormatBool(false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want data.Data
	}{
		{"+ve", NewBooleanData(true), &BooleanData{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.BooleanData
		want []byte
	}{
		{"+ve", NewBooleanData(true), []byte{0x1}},
		{"+ve", &BooleanData{false}, []byte{0x0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BooleanDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", "true", &BooleanData{true}, false},
		{"+ve", "false", &BooleanData{false}, false},
		{"-ve", "abc", &BooleanData{false}, true},
		{"+ve", "", &BooleanData{false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeBooleanData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("BooleanDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("BooleanDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
