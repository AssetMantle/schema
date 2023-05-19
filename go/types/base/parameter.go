package base

import (
	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/types"
)

var _ types.Parameter = (*Parameter)(nil)

func (m *Parameter) ValidateBasic() error {
	return m.MetaProperty.ValidateBasic()
}
func (m *Parameter) GetMetaProperty() properties.MetaProperty {
	return m.MetaProperty
}
func (m *Parameter) Mutate(data data.Data) types.Parameter {
	if m.MetaProperty.GetData().GetTypeID().Compare(data.GetTypeID()) == 0 {
		m.MetaProperty = base.NewMetaProperty(m.MetaProperty.GetKey(), data).(*base.MetaProperty)
	}
	return m
}

func ParametersToInterface(inParameters []*Parameter) []types.Parameter {
	returnParameters := make([]types.Parameter, len(inParameters))

	for i, parameter := range inParameters {
		returnParameters[i] = parameter
	}

	return returnParameters
}

func ParametersFromInterface(parameters []types.Parameter) []*Parameter {
	returnParameters := make([]*Parameter, len(parameters))

	for i, parameter := range parameters {
		returnParameters[i] = parameter.(*Parameter)
	}

	return returnParameters
}

func NewParameter(metaProperty properties.MetaProperty) *Parameter {
	return &Parameter{
		MetaProperty: metaProperty.(*base.MetaProperty),
	}
}
