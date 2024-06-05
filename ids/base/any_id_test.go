// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/base64"
	"fmt"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/constants"
	"github.com/AssetMantle/schema/ids/utilities"
	"reflect"
	"strings"
	"testing"
)

var (
	testValidBase64String    = "nBuFnhfmBVznCR9vS5/V1mqZim8aclm1jBlqR8NGU94="
	testValidBase64URLString = "nBuFnhfmBVznCR9vS5/V1mqZim8aclm1jBlqR8NGU94="
	testBytes, _             = base64.StdEncoding.DecodeString(testValidBase64String)
)

func Test_AnyIDGetTypeID(t *testing.T) {
	tests := []struct {
		name string
		args ids.AnyID
		want ids.StringID
	}{
		{"+ve", (&AssetID{&HashID{testBytes}}).ToAnyID(), NewStringID(constants.AssetIDType)},
		{"+ve", (&ClassificationID{&HashID{testBytes}}).ToAnyID(), NewStringID(constants.ClassificationIDType)},
		{"+ve", (&DataID{&StringID{"typeID"}, &HashID{testBytes}}).ToAnyID(), NewStringID(constants.DataIDType)},
		{"+ve", (&HashID{testBytes}).ToAnyID(), NewStringID(constants.HashIDType)},
		{"+ve", (&IdentityID{&HashID{testBytes}}).ToAnyID(), NewStringID(constants.IdentityIDType)},
		{"+ve", (&MaintainerID{&HashID{testBytes}}).ToAnyID(), NewStringID(constants.MaintainerIDType)},
		{"+ve", (&OrderID{&HashID{testBytes}}).ToAnyID(), NewStringID(constants.OrderIDType)},
		{"+ve", (&PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}).ToAnyID(), NewStringID(constants.PropertyIDType)},
		{"+ve", (&StringID{"test"}).ToAnyID(), NewStringID(constants.StringIDType)},
		{"+ve", (&SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}).ToAnyID(), NewStringID(constants.SplitIDType)},
		{"+ve", &AnyID{}, NewStringID(constants.AnyIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnyIDGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AnyIDFromString(t *testing.T) {
	type arg struct {
		typeID string
		value  string
	}
	tests := []struct {
		name    string
		arg     arg
		want    ids.ID
		wantErr error
	}{
		{"+ve", arg{constants.AssetIDType, testValidBase64String}, (&AssetID{&HashID{testBytes}}).ToAnyID(), nil},
		{"+ve", arg{constants.ClassificationIDType, testValidBase64String}, (&ClassificationID{&HashID{testBytes}}).ToAnyID(), nil},
		{"+ve", arg{constants.DataIDType, strings.Join([]string{"typeID", testValidBase64String}, utilities.IDSeparator)}, (&DataID{&StringID{"typeID"}, &HashID{testBytes}}).ToAnyID(), nil},
		{"+ve", arg{constants.HashIDType, testValidBase64String}, (&HashID{testBytes}).ToAnyID(), nil},
		{"+ve", arg{constants.IdentityIDType, testValidBase64String}, (&IdentityID{&HashID{testBytes}}).ToAnyID(), nil},
		{"+ve", arg{constants.MaintainerIDType, testValidBase64String}, (&MaintainerID{&HashID{testBytes}}).ToAnyID(), nil},
		{"+ve", arg{constants.OrderIDType, testValidBase64String}, (&OrderID{&HashID{testBytes}}).ToAnyID(), nil},
		{"+ve", arg{constants.PropertyIDType, strings.Join([]string{"keyID", "typeID"}, utilities.IDSeparator)}, (&PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}).ToAnyID(), nil},
		{"+ve", arg{constants.SplitIDType, strings.Join([]string{testValidBase64String, testValidBase64String}, utilities.IDSeparator)}, (&SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}).ToAnyID(), nil},
		{"+ve", arg{constants.StringIDType, testValidBase64String}, (&StringID{testValidBase64String}).ToAnyID(), nil},
		{"+ve", arg{"", ""}, &AnyID{}, nil},
		{"-ve", arg{"unknown", ""}, nil, fmt.Errorf("unknown type identifier is not recognised")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := strings.Join([]string{tt.arg.typeID, tt.arg.value}, idTypeAndValueSeparator)
			got, err := PrototypeAnyID().FromString(v)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.wantErr.Error()) {
					t.Errorf("AnyIDFromString() got error = %v, want error = %v", err.Error(), tt.wantErr.Error())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnyIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AnyIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.AnyID
		want string
	}{
		{"+ve", (&AssetID{&HashID{testBytes}}).ToAnyID(), strings.Join([]string{constants.AssetIDType, testValidBase64String}, idTypeAndValueSeparator)},
		{"+ve", (&ClassificationID{&HashID{testBytes}}).ToAnyID(), strings.Join([]string{constants.ClassificationIDType, testValidBase64String}, idTypeAndValueSeparator)},
		{"+ve", (&DataID{&StringID{"typeID"}, &HashID{testBytes}}).ToAnyID(), strings.Join([]string{constants.DataIDType, strings.Join([]string{"typeID", testValidBase64String}, utilities.IDSeparator)}, idTypeAndValueSeparator)},
		{"+ve", (&HashID{testBytes}).ToAnyID(), strings.Join([]string{constants.HashIDType, testValidBase64String}, idTypeAndValueSeparator)},
		{"+ve", (&IdentityID{&HashID{testBytes}}).ToAnyID(), strings.Join([]string{constants.IdentityIDType, testValidBase64String}, idTypeAndValueSeparator)},
		{"+ve", (&MaintainerID{&HashID{testBytes}}).ToAnyID(), strings.Join([]string{constants.MaintainerIDType, testValidBase64String}, idTypeAndValueSeparator)},
		{"+ve", (&OrderID{&HashID{testBytes}}).ToAnyID(), strings.Join([]string{constants.OrderIDType, testValidBase64String}, idTypeAndValueSeparator)},
		{"+ve", (&PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}).ToAnyID(), strings.Join([]string{constants.PropertyIDType, strings.Join([]string{"keyID", "typeID"}, utilities.IDSeparator)}, idTypeAndValueSeparator)},
		{"+ve", (&StringID{"test"}).ToAnyID(), strings.Join([]string{constants.StringIDType, "test"}, idTypeAndValueSeparator)},
		{"+ve", (&SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}).ToAnyID(), strings.Join([]string{constants.SplitIDType, strings.Join([]string{testValidBase64String, testValidBase64String}, utilities.IDSeparator)}, idTypeAndValueSeparator)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnyIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AnyIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.AnyID
		want []byte
	}{
		{"+ve", (&AssetID{&HashID{testBytes}}).ToAnyID(), testBytes},
		{"+ve", (&ClassificationID{&HashID{testBytes}}).ToAnyID(), testBytes},
		{"+ve", (&DataID{&StringID{"typeID"}, &HashID{testBytes}}).ToAnyID(), append([]byte("typeID"), testBytes...)},
		{"+ve", (&HashID{testBytes}).ToAnyID(), testBytes},
		{"+ve", (&IdentityID{&HashID{testBytes}}).ToAnyID(), testBytes},
		{"+ve", (&MaintainerID{&HashID{testBytes}}).ToAnyID(), testBytes},
		{"+ve", (&OrderID{&HashID{testBytes}}).ToAnyID(), testBytes},
		{"+ve", (&PropertyID{&StringID{"keyID"}, &StringID{"typeID"}}).ToAnyID(), append([]byte("keyID"), []byte("typeID")...)},
		{"+ve", (&StringID{"test"}).ToAnyID(), []byte("test")},
		{"+ve", (&SplitID{&AssetID{&HashID{testBytes}}, &IdentityID{&HashID{testBytes}}}).ToAnyID(), append(testBytes, testBytes...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnyIDBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
