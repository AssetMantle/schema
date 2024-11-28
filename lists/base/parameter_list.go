package base

import (
	"fmt"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/parameters"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"sort"
)

var _ lists.ParameterList = (*ParameterList)(nil)

func (parameterList *ParameterList) ValidateBasic() error {
	parameterList.sort()

	parameterIDMap := map[string]bool{}

	for i, parameter := range parameterList.Get() {
		if parameter == nil {
			return fmt.Errorf("nil parameter in list at index %d", i)
		}

		if _, found := parameterIDMap[parameter.GetMetaProperty().GetID().AsString()]; found {
			return fmt.Errorf("duplicate parameter with ID: %s", parameter.GetMetaProperty().GetID().AsString())
		} else {
			parameterIDMap[parameter.GetMetaProperty().GetID().AsString()] = true
		}

		if err := parameter.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (parameterList *ParameterList) Get() []parameters.Parameter {
	parameterList.sort()
	return baseParametersToParameters(parameterList.Parameters...)
}
func baseParametersToParameters(baseParameters ...*baseParameters.Parameter) []parameters.Parameter {
	Parameters := make([]parameters.Parameter, len(baseParameters))
	for i, parameter := range baseParameters {
		Parameters[i] = parameter
	}

	return Parameters
}
func (parameterList *ParameterList) GetParameter(propertyID ids.PropertyID) parameters.Parameter {
	if index, found := parameterList.search(baseParameters.NewEmptyParameterFromID(propertyID)); found {
		return parameterList.Parameters[index]
	}

	return nil
}
func (parameterList *ParameterList) Mutate(parameters ...parameters.Parameter) lists.ParameterList {
	for _, property := range parameters {
		if property != nil {
			if index, found := parameterList.search(property); found {
				parameterList.Parameters[index] = parameterList.Parameters[index].Mutate(property.GetMetaProperty().GetData()).(*baseParameters.Parameter)
			}
		}
	}

	return parameterList
}
func (parameterList *ParameterList) sort() lists.ParameterList {
	var filteredParameters []*baseParameters.Parameter
	for _, parameter := range parameterList.Parameters {
		if parameter != nil {
			filteredParameters = append(filteredParameters, parameter)
		}
	}
	sort.Slice(filteredParameters, func(i, j int) bool {
		return filteredParameters[i].Compare(filteredParameters[j]) < 0
	})
	parameterList.Parameters = filteredParameters
	return parameterList
}
func (parameterList *ParameterList) search(parameter parameters.Parameter) (index int, found bool) {
	if parameter == nil {
		return -1, false
	}

	index = sort.Search(
		len(parameterList.Parameters),
		func(i int) bool {
			return parameterList.Parameters[i].Compare(parameter) >= 0
		},
	)

	if index < len(parameterList.Parameters) && parameterList.Parameters[index].Compare(parameter) == 0 {
		return index, true
	}

	return index, false
}
func (parameterList *ParameterList) add(parameters ...parameters.Parameter) lists.ParameterList {
	for _, parameter := range parameters {
		if parameter != nil {
			if index, found := parameterList.search(parameter); !found {
				parameterList.Parameters = append(parameterList.Parameters, parameter.(*baseParameters.Parameter))
				copy(parameterList.Parameters[index+1:], parameterList.Parameters[index:])
				parameterList.Parameters[index] = parameter.(*baseParameters.Parameter)
			} else {
				parameterList.Parameters[index] = parameter.(*baseParameters.Parameter)
			}
		}
	}
	return parameterList
}

func NewParameterList(parameters ...parameters.Parameter) lists.ParameterList {
	return (&ParameterList{}).add(parameters...)
}
