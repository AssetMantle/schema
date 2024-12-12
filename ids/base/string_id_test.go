// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/constants"
	"reflect"
	"testing"
)

func Test_StringIDValidateBasic(t *testing.T) {
	tests := []struct {
		name      string
		args      ids.StringID
		wantError bool
	}{
		{"+ve", &StringID{"test"}, false},
		{"+ve", &StringID{""}, false},
		{"-ve", &StringID{"*&^%$#"}, true},
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

func Test_StringIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.StringID
		want ids.StringID
	}{
		{"+ve", &StringID{"test"}, NewStringID(constants.StringIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
	}{
		{"+ve", "", &StringID{""}},
		{"+ve", "test", &StringID{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStringID(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.StringID
		want string
	}{
		{"+ve", &StringID{"test"}, "test"},
		{"+ve", &StringID{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.StringID
		want []byte
	}{
		{"+ve", &StringID{"test"}, []byte("test")},
		{"+ve", &StringID{""}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringIDCompare(t *testing.T) {
	stringID := &StringID{"test"}
	tests := []struct {
		name string
		args ids.StringID
		want bool
	}{
		{"+ve", &StringID{"test"}, true},
		{"-ve", &StringID{"test2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(stringID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("StringIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
