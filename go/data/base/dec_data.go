// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	dataConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/traits"
)

var _ data.DecData = (*DecData)(nil)

func (decData *DecData) ValidateBasic() error {
	if _, err := sdkTypes.NewDecFromStr(decData.Value); err != nil {
		return errorConstants.IncorrectFormat.Wrapf("dec data value %s is not a valid decimal", decData.Value)
	}

	return nil
}
func (decData *DecData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(decData)
}
func (decData *DecData) GetBondWeight() sdkTypes.Int {
	return dataConstants.DecDataWeight
}
func (decData *DecData) Compare(listable traits.Listable) int {
	compareDecData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(decData.Bytes(), compareDecData.Bytes())
}
func (decData *DecData) Bytes() []byte {
	return []byte(decData.Value)
}
func (decData *DecData) GetTypeID() ids.StringID {
	return dataConstants.DecDataTypeID
}
func (decData *DecData) ZeroValue() data.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData *DecData) GenerateHashID() ids.HashID {
	if decData.Compare(decData.ZeroValue()) == 0 {
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

	return NewDecData(dec), nil
}
func (decData *DecData) Get() sdkTypes.Dec {
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

func PrototypeDecData() data.DecData {
	return NewDecData(sdkTypes.ZeroDec()).ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.DecData {
	return &DecData{
		Value: value.String(),
	}
}
