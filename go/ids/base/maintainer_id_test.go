// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"reflect"
	"testing"
)

func Test_MaintainerIDValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args ids.MaintainerID
		want error
	}{
		{"+ve", &MaintainerID{&HashID{testBytes}}, nil},
		{"-ve", &MaintainerID{&HashID{[]byte("test")}}, errorConstants.IncorrectFormat.Wrapf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("MaintainerIDValidateBasic() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_MaintainerIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.MaintainerID
		want ids.StringID
	}{
		{"+ve", &MaintainerID{&HashID{testBytes}}, NewStringID(constants.MaintainerIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaintainerIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaintainerIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", testValidBase64String, &MaintainerID{&HashID{testBytes}}, nil},
		{"+ve", "", &MaintainerID{&HashID{[]byte{}}}, nil},
		{"+ve", "test", &MaintainerID{&HashID{[]byte{}}}, errorConstants.IncorrectFormat.Wrapf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeMaintainerID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("MaintainerIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaintainerIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaintainerIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.MaintainerID
		want string
	}{
		{"+ve", &MaintainerID{&HashID{testBytes}}, testValidBase64URLString},
		{"+ve", &MaintainerID{&HashID{[]byte{}}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaintainerIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaintainerIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.MaintainerID
		want []byte
	}{
		{"+ve", &MaintainerID{&HashID{testBytes}}, testBytes},
		{"+ve", &MaintainerID{&HashID{[]byte{}}}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaintainerIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaintainerIDCompare(t *testing.T) {
	maintainerID := &MaintainerID{&HashID{testBytes}}
	tests := []struct {
		name string
		args ids.MaintainerID
		want bool
	}{
		{"+ve", &MaintainerID{&HashID{testBytes}}, true},
		{"-ve", &MaintainerID{&HashID{[]byte("test")}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(maintainerID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("MaintainerIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
