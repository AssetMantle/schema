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

func Test_PropertyIDValidateBasic(t *testing.T) {
	tests := []struct {
		name      string
		args      ids.PropertyID
		wantError bool
	}{
		{"+ve", &PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}, false},
		// TODO: Regex always throwing true in testing
		{"-ve", &PropertyID{&StringID{"keyID"}, &StringID{"typ%?$%^eID"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err != nil && !tt.wantError {
				t.Errorf("PropertyIDValidateBasic() got = %v, want %v", err, tt.wantError)
			}
			if err == nil && tt.wantError {
				t.Errorf("PropertyIDValidateBasic() got = %v, want %v", err, tt.wantError)
			}
		})
	}
}

func Test_PropertyIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.PropertyID
		want ids.StringID
	}{
		{"+ve", &PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}, NewStringID(constants.PropertyIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PropertyIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PropertyIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", strings.Join([]string{"keyID", "typeID"}, utilities.IDSeparator), &PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}, nil},
		{"+ve", "", &PropertyID{&StringID{}, &StringID{}}, nil},
		{"+ve", "test", &PropertyID{&StringID{}, &StringID{}}, fmt.Errorf("invalid propertyID string test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypePropertyID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("PropertyIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PropertyIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PropertyIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.PropertyID
		want string
	}{
		{"+ve", &PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}, strings.Join([]string{"keyID", "typeID"}, utilities.IDSeparator)},
		{"+ve", &PropertyID{&StringID{}, &StringID{}}, "."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PropertyIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PropertyIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.PropertyID
		want []byte
	}{
		{"+ve", &PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}, append([]byte("keyID"), []byte("typeID")...)},
		{"+ve", &PropertyID{&StringID{}, &StringID{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PropertyIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PropertyIDCompare(t *testing.T) {
	propertyID := &PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}
	tests := []struct {
		name string
		args ids.PropertyID
		want bool
	}{
		{"+ve", &PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}, true},
		{"-ve", &PropertyID{&StringID{}, &StringID{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(propertyID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("PropertyIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
