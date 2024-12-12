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

func Test_ClassificationIDValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args ids.ClassificationID
		want error
	}{
		{"+ve", &ClassificationID{&HashID{testBytes}}, nil},
		{"-ve", &ClassificationID{&HashID{[]byte("test")}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ValidateBasic()
			if tt.want != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("ClassificationIDValidateBasic() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_ClassificationIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.ClassificationID
		want ids.StringID
	}{
		{"+ve", &ClassificationID{&HashID{testBytes}}, NewStringID(constants.ClassificationIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClassificationIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ClassificationIDFromString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ids.ID
		err  error
	}{
		{"+ve", testValidBase64String, &ClassificationID{&HashID{testBytes}}, nil},
		{"+ve", "", &ClassificationID{&HashID{[]byte{}}}, nil},
		{"+ve", "test", &ClassificationID{&HashID{[]byte{}}}, fmt.Errorf("invalid hashID length")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeClassificationID().FromString(tt.args)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("ClassificationIDFromString() got error = %v, want error %v", err.Error(), tt.err.Error())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClassificationIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ClassificationIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.ClassificationID
		want string
	}{
		{"+ve", &ClassificationID{&HashID{testBytes}}, testValidBase64URLString},
		{"+ve", &ClassificationID{&HashID{[]byte{}}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClassificationIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ClassificationIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.ClassificationID
		want []byte
	}{
		{"+ve", &ClassificationID{&HashID{testBytes}}, testBytes},
		{"+ve", &ClassificationID{&HashID{[]byte{}}}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClassificationIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ClassificationIDCompare(t *testing.T) {
	classificationID := &ClassificationID{&HashID{testBytes}}
	tests := []struct {
		name string
		args ids.ClassificationID
		want bool
	}{
		{"+ve", &ClassificationID{&HashID{testBytes}}, true},
		{"-ve", &ClassificationID{&HashID{[]byte("test")}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(classificationID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("ClassificationIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
