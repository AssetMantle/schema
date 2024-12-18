// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/types"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

var _ data.HeightData = (*HeightData)(nil)

func (heightData *HeightData) ValidateBasic() error {
	return heightData.Value.ValidateBasic()
}
func (heightData *HeightData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(heightData)
}
func (heightData *HeightData) GetBondWeight() math.Int {
	return dataConstants.HeightDataWeight
}
func (heightData *HeightData) AsString() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData *HeightData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeHeightData(), nil
	}

	Height, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return PrototypeHeightData(), err
	}

	heightData.Value = baseTypes.NewHeight(Height).(*baseTypes.Height)
	if err = heightData.ValidateBasic(); err != nil {
		return PrototypeHeightData(), err
	}

	return heightData, nil
}
func (heightData *HeightData) Compare(listableData data.ListableData) int {
	return heightData.Get().Compare(listableData.ToAnyListableData().Get().(*HeightData).Get())
}
func (heightData *HeightData) Bytes() []byte {
	Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(Bytes, uint64(heightData.Get().Get()))

	return Bytes
}
func (heightData *HeightData) GetTypeID() ids.StringID {
	return dataConstants.HeightDataTypeID
}
func (heightData *HeightData) ZeroValue() data.Data {
	return NewHeightData(baseTypes.NewHeight(0))
}
func (heightData *HeightData) GenerateHashID() ids.HashID {
	if heightData.Compare(heightData.ZeroValue().(data.ListableData)) == 0 {
		return baseIDs.GenerateHashID()
	}
	return baseIDs.GenerateHashID(heightData.Bytes())
}
func (heightData *HeightData) Get() types.Height {
	return heightData.Value
}
func (heightData *HeightData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_HeightData{
			HeightData: heightData,
		},
	}
}
func (heightData *HeightData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_HeightData{
			HeightData: heightData,
		},
	}
}

func PrototypeHeightData() data.HeightData {
	return NewHeightData(baseTypes.NewHeight(0)).ZeroValue().(data.HeightData)
}

func NewHeightData(value types.Height) data.HeightData {
	return &HeightData{
		Value: value.(*baseTypes.Height),
	}
}
