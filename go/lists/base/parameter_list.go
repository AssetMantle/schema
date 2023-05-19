package base

import (
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/types"
	"github.com/AssetMantle/schema/go/types/base"
)

var _ lists.ParameterList = (*ParameterList)(nil)

func (parameterList *ParameterList) Get() []types.Parameter {
	Parameters := make([]types.Parameter, len(parameterList.Parameters))
	for i, parameter := range parameterList.Parameters {
		Parameters[i] = parameter
	}
	return Parameters
}

func NewParameterList(parameters ...types.Parameter) lists.ParameterList {
	newParameters := make([]*base.Parameter, len(parameters))
	for i, parameter := range parameters {
		newParameters[i] = parameter.(*base.Parameter)
	}

	return &ParameterList{
		Parameters: newParameters,
	}
}
