// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"reflect"
	"testing"
)

func Test_OrderIDValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args ids.OrderID
		want error
	}{
		{"+ve", &OrderID{&HashID{testBytes}}, nil},
		{"-ve", &OrderID{&HashID{[]byte("test")}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("OrderIDValidateBasic() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_OrderIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.OrderID
		want ids.StringID
	}{
		{"+ve", &OrderID{&HashID{testBytes}}, NewStringID(constants.OrderIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_OrderIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", testValidBase64String, &OrderID{&HashID{testBytes}}, nil},
		{"+ve", "", &OrderID{&HashID{[]byte{}}}, nil},
		{"+ve", "test", &OrderID{&HashID{[]byte{}}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeOrderID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("OrderIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_OrderIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.OrderID
		want string
	}{
		{"+ve", &OrderID{&HashID{testBytes}}, testValidBase64URLString},
		{"+ve", &OrderID{&HashID{[]byte{}}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_OrderIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.OrderID
		want []byte
	}{
		{"+ve", &OrderID{&HashID{testBytes}}, testBytes},
		{"+ve", &OrderID{&HashID{[]byte{}}}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_OrderIDCompare(t *testing.T) {
	orderID := &OrderID{&HashID{testBytes}}
	tests := []struct {
		name string
		args ids.OrderID
		want bool
	}{
		{"+ve", &OrderID{&HashID{testBytes}}, true},
		{"-ve", &OrderID{&HashID{[]byte("test")}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(orderID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("OrderIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
