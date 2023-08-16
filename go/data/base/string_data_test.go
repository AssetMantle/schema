// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data"
	idsConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

func Test_NewStringData(t *testing.T) {
	tests := []struct {
		name string
		args string
		want data.Data
	}{
		{"+ve", "test", &StringData{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStringData(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataValidateBasic(t *testing.T) {
	longString := ""
	for i := 1; i <= 300; i++ {
		longString = longString + "a"
	}
	tests := []struct {
		name string
		args data.StringData
		want bool
	}{
		{"+ve", NewStringData("test"), false},
		{"-ve", &StringData{longString}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err == nil && tt.want {
				t.Errorf("StringDataValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("StringDataValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_StringData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want int
	}{
		{"+ve", NewStringData("b"), 0},
		{"+ve", NewStringData("a"), -1},
		{"+ve", NewStringData("c"), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(NewStringData("b"))
			if got != tt.want {
				t.Errorf("StringData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want ids.HashID
	}{
		{"+ve", NewStringData("test"), baseIDs.GenerateHashID([]byte("test")).(*baseIDs.HashID)},
		{"+ve", NewStringData(""), &baseIDs.HashID{[]byte{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataGet(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want string
	}{
		{"+ve", NewStringData("test"), "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDataGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want ids.DataID
	}{
		{"+ve", NewStringData("test"), &baseIDs.DataID{
			TypeID: idsConstants.StringDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID([]byte("test")).(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want ids.ID
	}{
		{"+ve", NewStringData("test"), idsConstants.StringDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want string
	}{
		{"+ve", NewStringData("test"), "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want data.Data
	}{
		{"+ve", NewStringData("test"), &StringData{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.StringData
		want []byte
	}{
		{name: "+ve", args: NewStringData("test"), want: []byte("test")},
		{name: "+ve", args: NewStringData("0"), want: []byte("0")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", "1", &StringData{"1"}, false},
		{"+ve", "", PrototypeStringData(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeStringData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("StringDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("StringDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
