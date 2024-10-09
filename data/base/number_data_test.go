// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/data"
	idsConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

func Test_NewNumberData(t *testing.T) {
	tests := []struct {
		name string
		args math.Int
		want data.Data
	}{
		{"+ve", sdkTypes.NewInt(10), &NumberData{"10"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNumberData(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNumberData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want bool
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), false},
		{"-ve", &NumberData{"abc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err == nil && tt.want {
				t.Errorf("NumberDataValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("NumberDataValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_NumberData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want int
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), 0},
		{"+ve", NewNumberData(sdkTypes.NewInt(11)), 1},
		{"+ve", NewNumberData(sdkTypes.NewInt(9)), -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(NewNumberData(sdkTypes.NewInt(10)))
			if got != tt.want {
				t.Errorf("NumberData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want ids.HashID
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), baseIDs.GenerateHashID([]byte("10")).(*baseIDs.HashID)},
		{name: "+ve", args: NewNumberData(sdkTypes.NewInt(0)), want: &baseIDs.HashID{IDBytes: []byte{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataGet(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want math.Int
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), sdkTypes.NewInt(10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberDataGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want ids.DataID
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), &baseIDs.DataID{
			TypeID: idsConstants.NumberDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID([]byte("10")).(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want ids.ID
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), idsConstants.NumberDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want string
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), "10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want data.Data
	}{
		{"+ve", NewNumberData(sdkTypes.NewInt(10)), &NumberData{"0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.NumberData
		want []byte
	}{
		{name: "+ve", args: NewNumberData(sdkTypes.NewInt(10)), want: []byte("10")},
		{name: "+ve", args: NewNumberData(sdkTypes.NewInt(0)), want: []byte("0")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NumberDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", "1", &NumberData{"1"}, false},
		{"-ve", "1.0", PrototypeNumberData(), true},
		{"-ve", "abc", PrototypeNumberData(), true},
		{"+ve", "", PrototypeNumberData(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeNumberData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("NumberDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("NumberDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
