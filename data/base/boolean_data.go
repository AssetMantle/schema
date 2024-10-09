// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"strconv"
	"strings"

	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var _ data.BooleanData = (*BooleanData)(nil)

func (booleanData *BooleanData) ValidateBasic() error {
	return nil
}
func (booleanData *BooleanData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(booleanData)
}
func (booleanData *BooleanData) GetBondWeight() math.Int {
	return dataConstants.BooleanDataWeight
}
func (booleanData *BooleanData) Compare(listableData data.ListableData) int {
	compareBooleanData := listableData.ToAnyListableData().Get().(*BooleanData)
	if booleanData.Get() == compareBooleanData.Get() {
		return 0
	} else if booleanData.Get() && !compareBooleanData.Get() {
		return 1
	}
	return -1
}
func (booleanData *BooleanData) AsString() string {
	return strconv.FormatBool(booleanData.Value)
}
func (booleanData *BooleanData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeBooleanData(), nil
	}

	Bool, err := strconv.ParseBool(dataString)
	if err != nil {
		return PrototypeBooleanData(), err
	}

	booleanData.Value = Bool
	if err = booleanData.ValidateBasic(); err != nil {
		return PrototypeBooleanData(), err
	}

	return booleanData, nil
}
func (booleanData *BooleanData) Bytes() []byte {
	if booleanData.Get() {
		return []byte{0x1}
	}
	return []byte{0x0}
}
func (booleanData *BooleanData) GetTypeID() ids.StringID {
	return dataConstants.BooleanDataTypeID
}
func (booleanData *BooleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData *BooleanData) GenerateHashID() ids.HashID {
	if booleanData.Compare(booleanData.ZeroValue().(data.ListableData)) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(booleanData.Bytes())
}
func (booleanData *BooleanData) Get() bool {
	return booleanData.Value
}
func (booleanData *BooleanData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_BooleanData{
			BooleanData: booleanData,
		},
	}
}
func (booleanData *BooleanData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_BooleanData{
			BooleanData: booleanData,
		},
	}
}

func PrototypeBooleanData() data.BooleanData {
	return &BooleanData{}
}

func NewBooleanData(value bool) data.BooleanData {
	return &BooleanData{
		Value: value,
	}
}
