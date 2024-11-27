package base

import (
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/parameters"
	"github.com/AssetMantle/schema/properties"
	"github.com/AssetMantle/schema/properties/base"
)

var _ parameters.Parameter = (*Parameter)(nil)

func (parameter *Parameter) ValidateBasic() error {
	return parameter.MetaProperty.ValidateBasic()
}
func (parameter *Parameter) GetMetaProperty() properties.MetaProperty {
	return parameter.MetaProperty
}
func (parameter *Parameter) Compare(otherParameter parameters.Parameter) int {
	return parameter.MetaProperty.Compare(otherParameter.GetMetaProperty())
}
func (parameter *Parameter) Mutate(data data.Data) parameters.Parameter {
	if parameter.MetaProperty.GetData().GetTypeID().Compare(data.GetTypeID()) == 0 {
		parameter.MetaProperty = base.NewMetaProperty(parameter.MetaProperty.GetKey(), data).(*base.MetaProperty)
	}

	return parameter
}

func NewParameter(metaProperty properties.MetaProperty) *Parameter {
	return &Parameter{
		MetaProperty: metaProperty.(*base.MetaProperty),
	}
}
