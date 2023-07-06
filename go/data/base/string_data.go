// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	dataConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/data/utilities"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

var _ data.StringData = (*StringData)(nil)

func (stringData *StringData) ValidateBasic() error {
	if !utilities.IsValidStringData(stringData.Value) {
		return errorConstants.IncorrectFormat
	}

	return nil
}

func (stringData *StringData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(stringData)
}
func (stringData *StringData) GetBondWeight() sdkTypes.Int {
	return dataConstants.StringDataWeight
}
func (stringData *StringData) Compare(listableData data.ListableData) int {
	return bytes.Compare(stringData.Bytes(), listableData.Bytes())
}
func (stringData *StringData) Bytes() []byte {
	return []byte(stringData.Value)
}
func (stringData *StringData) GetTypeID() ids.StringID {
	return dataConstants.StringDataTypeID
}
func (stringData *StringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData *StringData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(stringData.Bytes())
}
func (stringData *StringData) AsString() string {
	return stringData.Value
}
func (stringData *StringData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeStringData(), nil
	}

	stringData.Value = dataString
	if err := stringData.ValidateBasic(); err != nil {
		return PrototypeStringData(), err
	}

	return stringData, nil
}
func (stringData *StringData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_StringData{
			StringData: stringData,
		},
	}
}
func (stringData *StringData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_StringData{
			StringData: stringData,
		},
	}
}

func PrototypeStringData() data.StringData {
	return NewStringData("")
}

func NewStringData(value string) data.StringData {
	return &StringData{
		Value: value,
	}
}
