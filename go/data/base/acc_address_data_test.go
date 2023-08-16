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
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var (
	address    = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	accAddress = sdkTypes.MustAccAddressFromBech32(address)
)

func Test_NewAccAddressData(t *testing.T) {
	tests := []struct {
		name string
		args sdkTypes.AccAddress
		want data.Data
	}{
		{"+ve", accAddress, &AccAddressData{accAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAccAddressData(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccAddressData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want bool
	}{
		{"+ve", NewAccAddressData(accAddress), false},
		{"+ve", &AccAddressData{[]byte{}}, false},
		{"-ve", &AccAddressData{[]byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != (got != nil) {
				t.Errorf("AccAddressDataValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want bool
	}{
		{"+ve", NewAccAddressData(accAddress), true},
		{"+ve", &AccAddressData{[]byte{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(NewAccAddressData(accAddress))
			if (got == 0) != tt.want {
				t.Errorf("AccAddressData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want ids.ID
	}{
		{"+ve", NewAccAddressData(accAddress), baseIDs.GenerateHashID(accAddress)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataGet(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want sdkTypes.AccAddress
	}{
		{"+ve", NewAccAddressData(accAddress), accAddress},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressDataGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want ids.DataID
	}{
		{"+ve", NewAccAddressData(accAddress), &baseIDs.DataID{
			TypeID: idsConstants.AccAddressDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID(accAddress).(*baseIDs.HashID),
		}},
		{"+ve", &AccAddressData{[]byte{}}, &baseIDs.DataID{
			TypeID: idsConstants.AccAddressDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID().(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want ids.ID
	}{
		{"+ve", NewAccAddressData(accAddress), idsConstants.AccAddressDataTypeID},
		{"+ve", &AccAddressData{[]byte{}}, idsConstants.AccAddressDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want string
	}{
		{"+ve", NewAccAddressData(accAddress), address},
		{"+ve", &AccAddressData{[]byte{}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want data.Data
	}{
		{"+ve", NewAccAddressData(accAddress), &AccAddressData{sdkTypes.AccAddress{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.AccAddressData
		want []byte
	}{
		{"+ve", NewAccAddressData(accAddress), accAddress},
		{"+ve", &AccAddressData{[]byte{}}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AccAddressDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", address, &AccAddressData{accAddress}, false},
		{"-ve", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", &AccAddressData{[]byte{}}, true},
		{"+ve", "", &AccAddressData{[]byte{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeAccAddressData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccAddressDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("AccAddressDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("AccAddressDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
