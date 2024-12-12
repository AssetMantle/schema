// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/properties"
)

type getter interface {
	get() properties.Property
}

var _ getter = (*AnyProperty_MetaProperty)(nil)
var _ getter = (*AnyProperty_MesaProperty)(nil)

func (m *AnyProperty_MetaProperty) get() properties.Property {
	return m.MetaProperty
}
func (m *AnyProperty_MesaProperty) get() properties.Property {
	return m.MesaProperty
}

var _ properties.AnyProperty = (*AnyProperty)(nil)

func (anyProperty *AnyProperty) IsAnyProperty() {}
func (anyProperty *AnyProperty) Get() properties.Property {
	return anyProperty.Impl.(getter).get()
}
func (anyProperty *AnyProperty) GetID() ids.PropertyID {
	return anyProperty.Impl.(getter).get().GetID()
}
func (anyProperty *AnyProperty) GetDataID() ids.DataID {
	return anyProperty.Impl.(getter).get().GetDataID()
}
func (anyProperty *AnyProperty) GetKey() ids.StringID {
	return anyProperty.Impl.(getter).get().GetKey()
}
func (anyProperty *AnyProperty) GetDataTypeID() ids.StringID {
	return anyProperty.Impl.(getter).get().GetDataTypeID()
}
func (anyProperty *AnyProperty) GetBondWeight() math.Int {
	return anyProperty.Impl.(getter).get().GetBondWeight()
}
func (anyProperty *AnyProperty) IsMeta() bool {
	return anyProperty.Impl.(getter).get().IsMeta()
}
func (anyProperty *AnyProperty) ToAnyProperty() properties.AnyProperty {
	return anyProperty.Impl.(getter).get().ToAnyProperty()
}
func (anyProperty *AnyProperty) Compare(property properties.Property) int {
	return anyProperty.Impl.(getter).get().Compare(property)
}
func (anyProperty *AnyProperty) ValidateBasic() error {
	return anyProperty.Impl.(getter).get().ValidateBasic()
}
func (anyProperty *AnyProperty) Mutate(data data.Data) (properties.Property, error) {
	return anyProperty.Impl.(getter).get().Mutate(data)
}
