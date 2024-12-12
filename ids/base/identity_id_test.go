// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/constants"
	"reflect"
	"testing"
)

func Test_IdentityIDValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args ids.IdentityID
		want error
	}{
		{"+ve", &IdentityID{&HashID{testBytes}}, nil},
		{"-ve", &IdentityID{&HashID{[]byte("test")}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("IdentityIDValidateBasic() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_IdentityIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.IdentityID
		want ids.StringID
	}{
		{"+ve", &IdentityID{&HashID{testBytes}}, NewStringID(constants.IdentityIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentityIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IdentityIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", testValidBase64String, &IdentityID{&HashID{testBytes}}, nil},
		{"+ve", "", &IdentityID{&HashID{[]byte{}}}, nil},
		{"+ve", "test", &IdentityID{&HashID{[]byte{}}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeIdentityID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("IdentityIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentityIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IdentityIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.IdentityID
		want string
	}{
		{"+ve", &IdentityID{&HashID{testBytes}}, testValidBase64URLString},
		{"+ve", &IdentityID{&HashID{[]byte{}}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentityIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IdentityIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.IdentityID
		want []byte
	}{
		{"+ve", &IdentityID{&HashID{testBytes}}, testBytes},
		{"+ve", &IdentityID{&HashID{[]byte{}}}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentityIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IdentityIDCompare(t *testing.T) {
	identityID := &IdentityID{&HashID{testBytes}}
	tests := []struct {
		name string
		args ids.IdentityID
		want bool
	}{
		{"+ve", &IdentityID{&HashID{testBytes}}, true},
		{"-ve", &IdentityID{&HashID{[]byte("test")}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(identityID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("IdentityIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
