package base

import (
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/parameters"
	"github.com/AssetMantle/schema/parameters/base"
)

var _ lists.ParameterList = (*ParameterList)(nil)

func (parameterList *ParameterList) Get() []parameters.Parameter {
	Parameters := make([]parameters.Parameter, len(parameterList.Parameters))
	for i, parameter := range parameterList.Parameters {
		Parameters[i] = parameter
	}
	return Parameters
}

func NewParameterList(parameters ...parameters.Parameter) lists.ParameterList {
	newParameters := make([]*base.Parameter, len(parameters))
	for i, parameter := range parameters {
		newParameters[i] = parameter.(*base.Parameter)
	}

	return &ParameterList{
		Parameters: newParameters,
	}
}
