// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/ids"
)

type listableDataGetter interface {
	get() data.ListableData
}

func (x *AnyListableData_AccAddressData) get() data.ListableData {
	return x.AccAddressData
}
func (x *AnyListableData_BooleanData) get() data.ListableData {
	return x.BooleanData
}
func (x *AnyListableData_DecData) get() data.ListableData {
	return x.DecData
}
func (x *AnyListableData_HeightData) get() data.ListableData {
	return x.HeightData
}
func (x *AnyListableData_IDData) get() data.ListableData {
	return x.IDData
}
func (x *AnyListableData_NumberData) get() data.ListableData {
	return x.NumberData
}
func (x *AnyListableData_StringData) get() data.ListableData {
	return x.StringData
}
func (x *AnyListableData_LinkedData) get() data.ListableData {
	return x.LinkedData
}

var _ data.AnyListableData = (*AnyListableData)(nil)

func (anyListableData *AnyListableData) IsAnyListableData() {
}
func (anyListableData *AnyListableData) Get() data.Data {
	return anyListableData.Impl.(listableDataGetter).get()
}
func (anyListableData *AnyListableData) GetID() ids.DataID {
	return anyListableData.Impl.(listableDataGetter).get().GetID()
}
func (anyListableData *AnyListableData) GetBondWeight() math.Int {
	return anyListableData.Impl.(listableDataGetter).get().GetBondWeight()
}
func (anyListableData *AnyListableData) ValidateBasic() error {
	return anyListableData.Impl.(listableDataGetter).get().ValidateBasic()
}
func (anyListableData *AnyListableData) AsString() string {
	return anyListableData.Impl.(listableDataGetter).get().AsString()
}
func (anyListableData *AnyListableData) FromString(dataString string) (data.Data, error) {
	return anyListableData.Impl.(listableDataGetter).get().FromString(dataString)
}
func (anyListableData *AnyListableData) Bytes() []byte {
	return anyListableData.Impl.(listableDataGetter).get().Bytes()
}
func (anyListableData *AnyListableData) GetTypeID() ids.StringID {
	return anyListableData.Impl.(listableDataGetter).get().GetTypeID()
}
func (anyListableData *AnyListableData) ZeroValue() data.Data {
	return anyListableData.Impl.(listableDataGetter).get().ZeroValue()
}
func (anyListableData *AnyListableData) GenerateHashID() ids.HashID {
	return anyListableData.Impl.(listableDataGetter).get().GenerateHashID()
}
func (anyListableData *AnyListableData) ToAnyData() data.AnyData {
	return anyListableData.Impl.(listableDataGetter).get().ToAnyData()
}
func (anyListableData *AnyListableData) ToAnyListableData() data.AnyListableData {
	return anyListableData.Impl.(listableDataGetter).get().ToAnyListableData()
}
func (anyListableData *AnyListableData) Compare(listableData data.ListableData) int {
	return anyListableData.Impl.(listableDataGetter).get().Compare(listableData)
}
