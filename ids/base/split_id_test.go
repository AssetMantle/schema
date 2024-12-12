// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/constants"
	"github.com/AssetMantle/schema/ids/utilities"
	"reflect"
	"strings"
	"testing"
)

func Test_SplitIDValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args ids.SplitID
		want error
	}{
		{"+ve", &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}, nil},
		{"-ve", &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{[]byte("asda")}}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("SplitIDValidateBasic() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_SplitIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.SplitID
		want ids.StringID
	}{
		{"+ve", &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}, NewStringID(constants.SplitIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SplitIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", strings.Join([]string{testValidBase64URLString, testValidBase64URLString}, utilities.IDSeparator), &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}, nil},
		{"+ve", ".", PrototypeSplitID(), nil},
		{"+ve", testValidBase64URLString, &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{[]byte{}}}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeSplitID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("SplitIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SplitIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.SplitID
		want string
	}{
		{"+ve", &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}, strings.Join([]string{testValidBase64URLString, testValidBase64URLString}, utilities.IDSeparator)},
		{"+ve", &SplitID{&AssetID{&HashID{}}, &IdentityID{&HashID{}}}, "."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SplitIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.SplitID
		want []byte
	}{
		{"+ve", &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}, append(testBytes, testBytes...)},
		{"+ve", &SplitID{&AssetID{&HashID{}}, &IdentityID{&HashID{}}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SplitIDCompare(t *testing.T) {
	splitID := &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}
	tests := []struct {
		name string
		args ids.SplitID
		want bool
	}{
		{"+ve", &SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}, true},
		{"-ve", &SplitID{&AssetID{&HashID{}}, &IdentityID{&HashID{}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(splitID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("SplitIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
