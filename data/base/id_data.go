// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"strings"

	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var _ data.IDData = (*IDData)(nil)

func (idData *IDData) ValidateBasic() error {
	return idData.Value.ValidateBasic()
}
func (idData *IDData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(idData)
}
func (idData *IDData) GetBondWeight() math.Int {
	return dataConstants.IDDataWeight
}
func (idData *IDData) AsString() string {
	return idData.Value.AsString()
}
func (idData *IDData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeIDData(), nil
	}

	id, err := baseIDs.PrototypeAnyID().FromString(dataString)
	if err != nil {
		return PrototypeIDData(), err
	}

	idData.Value = id.ToAnyID().(*baseIDs.AnyID)
	if err = idData.ValidateBasic(); err != nil {
		return PrototypeIDData(), err
	}

	return idData, nil
}
func (idData *IDData) Compare(listableData data.ListableData) int {
	return idData.Get().Compare(listableData.ToAnyListableData().Get().(*IDData).Get())
}
func (idData *IDData) Bytes() []byte {
	return idData.Value.Bytes()
}
func (idData *IDData) GetTypeID() ids.StringID {
	return idData.Value.GetTypeID()
}
func (idData *IDData) ZeroValue() data.Data {
	return PrototypeIDData()
}
func (idData *IDData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(idData.Bytes())
}
func (idData *IDData) Get() ids.AnyID {
	return idData.Value
}
func (idData *IDData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_IDData{
			IDData: idData,
		},
	}
}
func (idData *IDData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_IDData{
			IDData: idData,
		},
	}
}

func PrototypeIDData() data.IDData {
	return NewIDData(baseIDs.PrototypeAnyID())
}

func NewIDData(value ids.ID) data.IDData {
	return &IDData{
		Value: value.ToAnyID().(*baseIDs.AnyID),
	}
}
