// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/constants"
	"reflect"
	"testing"
)

func Test_AssetIDValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args ids.AssetID
		want error
	}{
		{"+ve", &AssetID{&HashID{testBytes}}, nil},
		{"-ve", &AssetID{&HashID{[]byte("test")}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("AssetIDValidateBasic() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_AssetIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.AssetID
		want ids.StringID
	}{
		{"+ve", &AssetID{&HashID{testBytes}}, NewStringID(constants.AssetIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AssetIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", testValidBase64String, &AssetID{&HashID{testBytes}}, nil},
		{"+ve", "", &AssetID{&HashID{[]byte{}}}, nil},
		{"+ve", "test", &AssetID{&HashID{[]byte{}}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeAssetID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("AssetIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AssetIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.AssetID
		want string
	}{
		{"+ve", &AssetID{&HashID{testBytes}}, testValidBase64URLString},
		{"+ve", &AssetID{&HashID{[]byte{}}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AssetIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.AssetID
		want []byte
	}{
		{"+ve", &AssetID{&HashID{testBytes}}, testBytes},
		{"+ve", &AssetID{&HashID{[]byte{}}}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AssetIDCompare(t *testing.T) {
	assetID := &AssetID{&HashID{testBytes}}
	tests := []struct {
		name string
		args ids.AssetID
		want bool
	}{
		{"+ve", &AssetID{&HashID{testBytes}}, true},
		{"-ve", &AssetID{&HashID{[]byte("test")}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(assetID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("AssetIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
