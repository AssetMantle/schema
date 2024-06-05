// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
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

func Test_DataIDValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args ids.DataID
		want error
	}{
		{"+ve", &DataID{&StringID{"S"}, &HashID{testBytes}}, nil},
		{"-ve", &DataID{&StringID{"S"}, &HashID{[]byte("jhavsjh")}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("DataIDValidateBasic() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_DataIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.DataID
		want ids.StringID
	}{
		{"+ve", &DataID{&StringID{"S"}, &HashID{testBytes}}, NewStringID(constants.DataIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DataIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", strings.Join([]string{"S", testValidBase64URLString}, utilities.IDSeparator), &DataID{&StringID{"S"}, &HashID{testBytes}}, nil},
		{"+ve", ".", PrototypeDataID(), nil},
		{"+ve", "test", PrototypeDataID(), fmt.Errorf("expected composite id")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeDataID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("DataIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DataIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.DataID
		want string
	}{
		{"+ve", &DataID{&StringID{"S"}, &HashID{testBytes}}, strings.Join([]string{"S", testValidBase64URLString}, utilities.IDSeparator)},
		{"+ve", &DataID{&StringID{}, &HashID{}}, "."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DataIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.DataID
		want []byte
	}{
		{"+ve", &DataID{&StringID{"S"}, &HashID{testBytes}}, append([]byte("S"), testBytes...)},
		{"+ve", &DataID{&StringID{}, &HashID{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DataIDCompare(t *testing.T) {
	dataID := &DataID{&StringID{"S"}, &HashID{testBytes}}
	tests := []struct {
		name string
		args ids.DataID
		want bool
	}{
		{"+ve", &DataID{&StringID{"S"}, &HashID{testBytes}}, true},
		{"-ve", &DataID{&StringID{}, &HashID{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(dataID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("DataIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
