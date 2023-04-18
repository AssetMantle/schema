package base

import (
	"github.com/AssetMantle/schema/go/parameters"
)

var _ parameters.ParameterList = (*ParameterList)(nil)

func (parameterList *ParameterList) Get() []parameters.Parameter {
	Parameters := make([]parameters.Parameter, len(parameterList.Parameters))
	for i, parameter := range parameterList.Parameters {
		Parameters[i] = parameter
	}
	return Parameters
}

func NewParameterList(parameters ...parameters.Parameter) parameters.ParameterList {
	newParameters := make([]*Parameter, len(parameters))
	for i, parameter := range parameters {
		newParameters[i] = parameter.(*Parameter)
	}

	return &ParameterList{
		Parameters: newParameters,
	}
}
