// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/binary"
	"github.com/AssetMantle/schema/go/types"
	"github.com/AssetMantle/schema/go/types/base"
	"math"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data"
	idsConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

func Test_NewHeightData(t *testing.T) {
	tests := []struct {
		name string
		args types.Height
		want data.Data
	}{
		{name: "+ve", args: base.NewHeight(10), want: &HeightData{&base.Height{10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHeightData(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeightData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want bool
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), false},
		{"+ve", NewHeightData(base.NewHeight(-10)), false},
		{"-ve", &HeightData{&base.Height{-10}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err == nil && tt.want {
				t.Errorf("HeightDataValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("HeightDataValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_HeightData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want int
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), 0},
		{"+ve", NewHeightData(base.NewHeight(11)), 1},
		{"+ve", NewHeightData(base.NewHeight(9)), -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(NewHeightData(base.NewHeight(10)))
			if got != tt.want {
				t.Errorf("HeightData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want ids.ID
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), baseIDs.GenerateHashID(NewHeightData(base.NewHeight(10)).Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataGet(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want types.Height
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), base.NewHeight(10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightDataGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want ids.DataID
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), &baseIDs.DataID{
			TypeID: idsConstants.HeightDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID(NewHeightData(base.NewHeight(10)).Bytes()).(*baseIDs.HashID),
		}},
		{"+ve", &HeightData{&base.Height{-1}}, &baseIDs.DataID{
			TypeID: idsConstants.HeightDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID().(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want ids.ID
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), idsConstants.HeightDataTypeID},
		{"+ve", &HeightData{&base.Height{-1}}, idsConstants.HeightDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want string
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), "10"},
		{"+ve", &HeightData{&base.Height{-1}}, "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.HeightData
		want data.Data
	}{
		{"+ve", NewHeightData(base.NewHeight(10)), &HeightData{&base.Height{-1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataBytes(t *testing.T) {
	want1 := make([]byte, 8)
	binary.LittleEndian.PutUint64(want1, uint64(10))
	want2 := make([]byte, 8)
	binary.LittleEndian.PutUint64(want2, uint64(math.MaxUint64))

	tests := []struct {
		name string
		args data.HeightData
		want []byte
	}{
		{name: "+ve", args: NewHeightData(base.NewHeight(10)), want: want1},
		{name: "+ve", args: NewHeightData(base.NewHeight(-1)), want: want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HeightDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", "1", NewHeightData(base.NewHeight(1)), false},
		{"-ve", "1.0", PrototypeHeightData(), true},
		{"-ve", "18446744073709551615", PrototypeHeightData(), true},
		{"-ve", "abc", PrototypeHeightData(), true},
		{"+ve", "", PrototypeHeightData(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeHeightData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("HeightDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("HeightDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
