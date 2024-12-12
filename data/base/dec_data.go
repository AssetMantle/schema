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
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

var _ data.DecData = (*DecData)(nil)

func (decData *DecData) ValidateBasic() error {
	if dec, err := sdkTypes.NewDecFromStr(decData.Value); err != nil {
		return fmt.Errorf("dec data value %s is not a valid decimal", decData.Value)
	} else if !sdkTypes.ValidSortableDec(dec) {
		return fmt.Errorf("dec value %s out of allowed range of -%s to %s", decData.Value, sdkTypes.MaxSortableDec.String(), sdkTypes.MaxSortableDec.String())
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
	return sdkTypes.SortableDecBytes(sdkTypes.MustNewDecFromStr(decData.Value))
}
func (decData *DecData) GetTypeID() ids.StringID {
	return dataConstants.DecDataTypeID
}
func (decData *DecData) ZeroValue() data.Data {
	return NewDecData(sdkTypes.ZeroDec())
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

	dec, err := sdkTypes.NewDecFromStr(dataString)
	if err != nil {
		return PrototypeDecData(), err
	}

	decData.Value = dec.String()
	if err = decData.ValidateBasic(); err != nil {
		return PrototypeDecData(), err
	}

	return decData, nil
}
func (decData *DecData) Get() sdkTypes.Dec {
	if decData.Value == "<nil>" {
		return sdkTypes.Dec{}
	}

	if value, err := sdkTypes.NewDecFromStr(decData.Value); err != nil {
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
	return NewDecData(sdkTypes.ZeroDec()).ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.DecData {
	return &DecData{
		Value: value.String(),
	}
}
