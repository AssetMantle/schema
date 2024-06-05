// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/data"
	idsConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

func Test_NewDecData(t *testing.T) {
	tests := []struct {
		name string
		args sdkTypes.Dec
		want data.Data
	}{
		{"+ve", sdkTypes.MustNewDecFromStr("1.0"), &DecData{sdkTypes.MustNewDecFromStr("1").String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDecData(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want bool
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), false},
		{"+ve", NewDecData(sdkTypes.MaxSortableDec.Add(sdkTypes.MustNewDecFromStr("1.0"))), true},
		{"-ve", &DecData{"abc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err == nil && tt.want {
				t.Errorf("DecDataValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("DecDataValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_DecData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want bool
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), true},
		{"-ve", NewDecData(sdkTypes.MustNewDecFromStr("2.0")), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(NewDecData(sdkTypes.MustNewDecFromStr("1.0")))
			if (got == 0) != tt.want {
				t.Errorf("DecData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want ids.ID
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), baseIDs.GenerateHashID(sdkTypes.SortableDecBytes(sdkTypes.MustNewDecFromStr("1.0")))},
		{"+ve", &DecData{"0.0"}, &baseIDs.HashID{[]byte{}}},
		{"+ve", &DecData{"0.0000"}, &baseIDs.HashID{[]byte{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataGet(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want sdkTypes.Dec
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), sdkTypes.MustNewDecFromStr("1.0")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecDataGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want ids.DataID
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), &baseIDs.DataID{
			TypeID: idsConstants.DecDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID(sdkTypes.SortableDecBytes(sdkTypes.MustNewDecFromStr("1.0"))).(*baseIDs.HashID),
		}},
		{"+ve", &DecData{"0"}, &baseIDs.DataID{
			TypeID: idsConstants.DecDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID().(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want ids.ID
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), idsConstants.DecDataTypeID},
		{"+ve", &DecData{}, idsConstants.DecDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want string
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), sdkTypes.MustNewDecFromStr("1.0").String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want data.Data
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), &DecData{sdkTypes.MustNewDecFromStr("0.0").String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.DecData
		want []byte
	}{
		{"+ve", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), sdkTypes.SortableDecBytes(sdkTypes.MustNewDecFromStr("1.0"))},
		{"+ve", &DecData{"0.0"}, sdkTypes.SortableDecBytes(sdkTypes.MustNewDecFromStr("0.0"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DecDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", "1.0", NewDecData(sdkTypes.MustNewDecFromStr("1.0")), false},
		{"-ve", sdkTypes.MaxSortableDec.Add(sdkTypes.MustNewDecFromStr("1.0")).String(), PrototypeDecData(), true},
		{"-ve", "abc", PrototypeDecData(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeDecData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("DecDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("DecDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
