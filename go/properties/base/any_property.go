package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/properties"
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

func (m *AnyProperty) IsAnyProperty() {}
func (m *AnyProperty) Get() properties.Property {
	return m.Impl.(getter).get()
}
func (m *AnyProperty) GetID() ids.PropertyID {
	return m.Impl.(getter).get().GetID()
}
func (m *AnyProperty) GetDataID() ids.DataID {
	return m.Impl.(getter).get().GetDataID()
}
func (m *AnyProperty) GetKey() ids.StringID {
	return m.Impl.(getter).get().GetKey()
}
func (m *AnyProperty) GetDataTypeID() ids.StringID {
	return m.Impl.(getter).get().GetDataTypeID()
}
func (m *AnyProperty) GetBondWeight() sdkTypes.Int {
	return m.Impl.(getter).get().GetBondWeight()
}
func (m *AnyProperty) IsMeta() bool {
	return m.Impl.(getter).get().IsMeta()
}
func (m *AnyProperty) IsMesa() bool {
	return m.Impl.(getter).get().IsMesa()
}
func (m *AnyProperty) ToAnyProperty() properties.AnyProperty {
	return m.Impl.(getter).get().ToAnyProperty()
}
func (m *AnyProperty) Compare(property properties.Property) int {
	return m.Impl.(getter).get().Compare(property)
}
func (m *AnyProperty) ValidateBasic() error {
	return m.Impl.(getter).get().ValidateBasic()
}
func (m *AnyProperty) Mutate(data data.Data) properties.Property {
	return m.Impl.(getter).get().Mutate(data)
}
