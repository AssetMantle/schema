// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/schema/go/data"
	dataConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/traits"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ data.AccAddressData = (*AccAddressData)(nil)

func (accAddressData *AccAddressData) ValidateBasic() error {
	return sdkTypes.VerifyAddressFormat(accAddressData.Value)
}
func (accAddressData *AccAddressData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(accAddressData)
}
func (accAddressData *AccAddressData) GetBondWeight() sdkTypes.Int {
	return dataConstants.AccAddressDataWeight
}
func (accAddressData *AccAddressData) Compare(listable traits.Listable) int {
	compareAccAddressData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(accAddressData.Bytes(), compareAccAddressData.Bytes())
}
func (accAddressData *AccAddressData) AsString() string {
	return sdkTypes.AccAddress(accAddressData.Value).String()
}
func (accAddressData *AccAddressData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeAccAddressData(), nil
	}

	accAddress, err := sdkTypes.AccAddressFromBech32(dataString)
	if err != nil {
		return PrototypeAccAddressData(), err
	}

	return NewAccAddressData(accAddress), nil
}
func (accAddressData *AccAddressData) Bytes() []byte {
	return sdkTypes.AccAddress(accAddressData.Value).Bytes()
}
func (accAddressData *AccAddressData) GetTypeID() ids.StringID {
	return dataConstants.AccAddressDataTypeID
}
func (accAddressData *AccAddressData) ZeroValue() data.Data {
	return PrototypeAccAddressData()
}
func (accAddressData *AccAddressData) GenerateHashID() ids.HashID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		// TODO test
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(accAddressData.Bytes())
}
func (accAddressData *AccAddressData) Get() sdkTypes.AccAddress {
	return accAddressData.Value
}
func (accAddressData *AccAddressData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_AccAddressData{
			AccAddressData: accAddressData,
		}}
}
func (accAddressData *AccAddressData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_AccAddressData{
			AccAddressData: accAddressData,
		}}
}

func PrototypeAccAddressData() data.Data {
	return NewAccAddressData(sdkTypes.AccAddress{})
}

func NewAccAddressData(value sdkTypes.AccAddress) data.AccAddressData {
	return &AccAddressData{
		Value: value,
	}
}
