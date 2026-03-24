// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
		"strings"
)

var _ data.DecData = (*DecData)(nil)

func (decData *DecData) ValidateBasic() error {
	if dec, err := math.LegacyNewDecFromStr(decData.Value); err != nil {
		return fmt.Errorf("dec data value %s is not a valid decimal", decData.Value)
	} else if !math.LegacyValidSortableDec(dec) {
		return fmt.Errorf("dec value %s out of allowed range of -%s to %s", decData.Value, math.LegacyMaxSortableDec.String(), math.LegacyMaxSortableDec.String())
	}

	return nil
}
func (decData *DecData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(decData)
}
func (decData *DecData) GetBondWeight() math.Int {
	return dataConstants.DecDataWeight
}
func (decData *DecData) Compare(listableData data.ListableData) int {
	if difference := decData.Get().Sub(listableData.ToAnyListableData().Get().(*DecData).Get()); difference.IsZero() {
		return 0
	} else if difference.IsPositive() {
		return 1
	}
	return -1
}
func (decData *DecData) Bytes() []byte {
	return math.LegacySortableDecBytes(math.LegacyMustNewDecFromStr(decData.Value))
}
func (decData *DecData) GetTypeID() ids.StringID {
	return dataConstants.DecDataTypeID
}
func (decData *DecData) ZeroValue() data.Data {
	return NewDecData(math.LegacyZeroDec())
}
func (decData *DecData) GenerateHashID() ids.HashID {
	if decData.Compare(decData.ZeroValue().(data.ListableData)) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(decData.Bytes())
}
func (decData *DecData) AsString() string {
	return decData.Value
}
func (decData *DecData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeDecData(), nil
	}

	dec, err := math.LegacyNewDecFromStr(dataString)
	if err != nil {
		return PrototypeDecData(), err
	}

	decData.Value = dec.String()
	if err = decData.ValidateBasic(); err != nil {
		return PrototypeDecData(), err
	}

	return decData, nil
}
func (decData *DecData) Get() math.LegacyDec {
	if decData.Value == "<nil>" {
		return math.LegacyDec{}
	}

	if value, err := math.LegacyNewDecFromStr(decData.Value); err != nil {
		panic(err)
	} else {
		return value
	}
}
func (decData *DecData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_DecData{
			DecData: decData,
		}}
}
func (decData *DecData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_DecData{
			DecData: decData,
		}}
}

func PrototypeDecData() data.DecData {
	return NewDecData(math.LegacyZeroDec()).ZeroValue().(data.DecData)
}

func NewDecData(value math.LegacyDec) data.DecData {
	return &DecData{
		Value: value.String(),
	}
}
