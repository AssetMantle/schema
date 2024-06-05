// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"encoding/base64"
	"fmt"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/data"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var (
	testValidBase64String    = "nBuFnhfmBVznCR9vS5/V1mqZim8aclm1jBlqR8NGU94="
	testValidBase64URLString = "nBuFnhfmBVznCR9vS5/V1mqZim8aclm1jBlqR8NGU94="
	testBytes, _             = base64.StdEncoding.DecodeString(testValidBase64String)
	testService              = "https://www.africau.edu/images/default/sample.jpg"
	testLinkedData           = NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("jpg"), testService)
)

func Test_NewLinkedData(t *testing.T) {
	type testArgLinkedData struct {
		resource  []byte
		service   string
		extension string
	}
	tests := []struct {
		name string
		args testArgLinkedData
		want data.Data
	}{
		{"+ve", testArgLinkedData{testBytes, testService, "jpg"}, &LinkedData{&baseIDs.HashID{testBytes}, &baseIDs.StringID{"jpg"}, testService}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLinkedData(&baseIDs.HashID{tt.args.resource}, baseIDs.NewStringID(tt.args.extension), tt.args.service)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want bool
	}{
		{"+ve", testLinkedData, false},
		{"-ve", NewLinkedData(&baseIDs.HashID{[]byte("abc")}, baseIDs.NewStringID("jpg"), testService), true},
		{"-ve", NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("//%jpg"), testService), true},
		{"-ve", NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("jpg"), "www.africau.edu/images/default/sample.jpg"), true},
		{"-ve", NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("jpg"), "https:///images/default/sample.jpg"), true},
		{"+ve", NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("jpg"), "https://images/default/sample.jpg"), false},
		{"-ve", NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("jpg"), "images/default/sample.jpg"), true},
		{"-ve", NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("jpg"), "ht tps://images/test"), true},
		{"-ve", NewLinkedData(&baseIDs.HashID{testBytes}, baseIDs.NewStringID("abc"), "images/default/sample.jpg"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err == nil && tt.want {
				t.Errorf("LinkedDataValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("LinkedDataValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_LinkedData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want int
	}{
		{"+ve", testLinkedData, 0},
		{"+ve", NewLinkedData(&baseIDs.HashID{[]byte{uint8(160)}}, baseIDs.NewStringID("jpg"), testService), 1},
		{"+ve", NewLinkedData(&baseIDs.HashID{[]byte{uint8(150)}}, baseIDs.NewStringID("jpg"), testService), -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(testLinkedData)
			if got != tt.want {
				t.Errorf("LinkedData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want ids.HashID
	}{
		{"+ve", testLinkedData, baseIDs.GenerateHashID(bytes.Join([][]byte{testBytes, []byte("jpg"), []byte(testService)}, dataConstants.ListBytesSeparator))},
		{"+ve", PrototypeLinkedData(), baseIDs.GenerateHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want ids.DataID
	}{
		{"+ve", testLinkedData, &baseIDs.DataID{
			TypeID: dataConstants.LinkedDataTypeID.(*baseIDs.StringID),
			HashID: baseIDs.GenerateHashID(bytes.Join([][]byte{testBytes, []byte("jpg"), []byte(testService)}, dataConstants.ListBytesSeparator)).(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want ids.ID
	}{
		{"+ve", testLinkedData, dataConstants.LinkedDataTypeID},
		{"+ve", PrototypeLinkedData(), dataConstants.LinkedDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want string
	}{
		{"+ve", testLinkedData, testValidBase64URLString + ",jpg," + testService},
		{"+ve", PrototypeLinkedData(), ",,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want data.Data
	}{
		{"+ve", testLinkedData, &LinkedData{
			ResourceID:      baseIDs.GenerateHashID().(*baseIDs.HashID),
			ExtensionID:     baseIDs.NewStringID("").(*baseIDs.StringID),
			ServiceEndpoint: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.LinkedData
		want []byte
	}{
		{name: "+ve", args: testLinkedData, want: bytes.Join([][]byte{testBytes, []byte("jpg"), []byte(testService)}, dataConstants.ListBytesSeparator)},
		{name: "+ve", args: PrototypeLinkedData(), want: bytes.Join([][]byte{dataConstants.ListBytesSeparator, dataConstants.ListBytesSeparator}, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LinkedDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", "", PrototypeLinkedData(), false},
		{"-ve", testValidBase64URLString, PrototypeLinkedData(), true},
		{"-ve", "1,a,b", PrototypeLinkedData(), true},
		{"-ve", testValidBase64URLString + ",a,b", PrototypeLinkedData(), true},
		{"-ve", testValidBase64URLString + ",a///,b", PrototypeLinkedData(), true},
		{"-ve", testValidBase64URLString + ",jpg,b", PrototypeLinkedData(), true},
		{"-ve", testValidBase64URLString + ",jpg," + testService, testLinkedData, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeLinkedData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil {
				fmt.Println(err.Error())
			}
			if err != nil && !tt.wantErr {
				t.Errorf("LinkedDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("LinkedDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
