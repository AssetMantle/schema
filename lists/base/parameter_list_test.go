// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/parameters"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

var (
	randomParameterGenerator = func() parameters.Parameter {
		return baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(strconv.Itoa(rand.Intn(99999999999999999))), baseData.NewStringData(strconv.Itoa(rand.Intn(rand.Intn(99999999999999999))))))
	}
	randomParametersGenerator = func(n int) (unsorted []parameters.Parameter, sorted []parameters.Parameter) {
		unsorted = make([]parameters.Parameter, n)
		for i := 0; i < n; i++ {
			unsorted[i] = randomParameterGenerator()
		}

		sorted = make([]parameters.Parameter, n)
		copy(sorted, unsorted)

		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].Compare(sorted[j]) <= 0
		})

		return unsorted, sorted
	}

	maxParameters                                                    = 100
	randomParameterIndex                                             = rand.Intn(maxParameters)
	randomUnsortedParameters, randomParameters                       = randomParametersGenerator(maxParameters)
	testParameterGeneratorOne                                        = baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("Test 1"), baseData.NewStringData("Test 2")))
	testParameterGeneratorTwo                                        = baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("Test 3"), baseData.NewStringData("Test 4")))
	veryLargeNumberOfUnsortedParameters, veryLargeNumberOfParameters = randomParametersGenerator(2000)
)

func Test_parameterList_add(t *testing.T) {
	tests := []struct {
		name       string
		added      lists.ParameterList
		properties []parameters.Parameter
		want       lists.ParameterList
	}{
		{"add one property",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{randomParameters[2]},
			NewParameterList(randomParameters[0], randomParameters[1], randomParameters[2]),
		},
		{"add already added property",
			NewParameterList(randomParameters[randomParameterIndex]),
			[]parameters.Parameter{randomParameters[randomParameterIndex]},
			NewParameterList(randomParameters[randomParameterIndex]),
		},
		{"add two properties",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{randomParameters[2], randomParameters[3]},
			NewParameterList(randomParameters[0], randomParameters[1], randomParameters[2], randomParameters[3]),
		},
		{"add many properties",
			NewParameterList(),
			randomParameters,
			NewParameterList(randomParameters...),
		},
		{"add many unsorted properties",
			NewParameterList(),
			randomUnsortedParameters,
			NewParameterList(randomParameters...),
		},
		{"add duplicate properties",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{randomParameters[2], randomParameters[2]},
			NewParameterList(randomParameters[0], randomParameters[1], randomParameters[2]),
		},
		{"add nil properties",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{nil},
			NewParameterList(randomParameters[0], randomParameters[1]),
		},
		{"add nil property with other properties",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{randomParameters[2], nil, randomParameters[3]},
			NewParameterList(randomParameters[0], randomParameters[1], randomParameters[2], randomParameters[3]),
		},
		{
			"add properties out of order",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{randomParameters[2], randomParameters[3]},
			NewParameterList(randomParameters[0], randomParameters[1], randomParameters[2], randomParameters[3]),
		},
		{
			"add prototype property",
			NewParameterList(),
			[]parameters.Parameter{baseParameters.NewParameter(baseProperties.PrototypeMetaProperty())},
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty())),
		},
		{
			"add prototype property with other properties",
			NewParameterList(),
			[]parameters.Parameter{baseParameters.NewParameter(baseProperties.PrototypeMetaProperty()), testParameterGeneratorOne, testParameterGeneratorTwo},
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty()), testParameterGeneratorOne, testParameterGeneratorTwo),
		},
		{
			"add multiple nils",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{nil, nil, nil},
			NewParameterList(randomParameters[0], randomParameters[1]),
		},
		{
			"very large number of properties",
			NewParameterList(),
			veryLargeNumberOfParameters,
			NewParameterList(veryLargeNumberOfParameters...),
		},
		{
			"very large number of unsorted properties",
			NewParameterList(),
			veryLargeNumberOfUnsortedParameters,
			NewParameterList(veryLargeNumberOfParameters...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldProperties := make([]parameters.Parameter, len(tt.properties))
			copy(oldProperties, tt.properties)

			if got := tt.added.(*ParameterList).add(tt.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = \n %v \n want \n %v", got, tt.want)
			}

			// check if input list was modified
			if !reflect.DeepEqual(tt.properties, oldProperties) {
				t.Errorf("input modified")
			}
		})
	}
}

func Test_parameterList_Mutate(t *testing.T) {
	tests := []struct {
		name       string
		added      lists.ParameterList
		parameters []parameters.Parameter
		want       lists.ParameterList
	}{
		{
			"empty list",
			NewParameterList(),
			[]parameters.Parameter{randomParameters[randomParameterIndex]},
			NewParameterList(),
		},
		{
			"one property",
			NewParameterList(randomParameters[0]),
			[]parameters.Parameter{randomParameters[randomParameterIndex]},
			NewParameterList(randomParameters[0]),
		},
		{
			"two properties",
			NewParameterList(randomParameters[0], randomParameters[1]),
			[]parameters.Parameter{randomParameters[2], randomParameters[3]},
			NewParameterList(randomParameters[0], randomParameters[1]),
		},
		{
			"many properties",
			NewParameterList(randomParameters...),
			[]parameters.Parameter{randomParameters[2], randomParameters[3]},
			NewParameterList(randomParameters...),
		},
		{
			"mutate one property in one property list",
			NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(0))))),
			[]parameters.Parameter{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))))},
			NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))))),
		},
		{
			"mutate one property in two property list",
			NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(0)))), randomParameters[randomParameterIndex]),
			[]parameters.Parameter{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))))},
			NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))), randomParameters[randomParameterIndex]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldParameters := make([]parameters.Parameter, len(tt.parameters))
			copy(oldParameters, tt.parameters)

			if got := tt.added.(*ParameterList).Mutate(tt.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}

			// check if input list was modified
			if !reflect.DeepEqual(tt.parameters, oldParameters) {
				t.Errorf("input modified")
			}
		})
	}
}

func Test_parameterList_GetList(t *testing.T) {
	tests := []struct {
		name       string
		properties []properties.Property
		want       []properties.Property
	}{
		{
			name:       "+ve",
			properties: []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))},
			want:       []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))},
		},
		{
			name:       "empty list",
			properties: []properties.Property{},
			want:       []properties.Property{},
		},
		{
			name:       "multiple properties",
			properties: []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("demand"), baseData.NewDecData(sdkTypes.NewDec(2)))},
			want:       []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("demand"), baseData.NewDecData(sdkTypes.NewDec(2)))},
		},
		{
			name:       "nil property",
			properties: []properties.Property{nil},
			want:       []properties.Property{},
		},
		{
			name:       "nil property with other properties",
			properties: []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), nil},
			want:       []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.properties...)}
			for i, got := range propertyList.Get() {
				if got.Compare(tt.want[i]) != 0 {
					t.Errorf("GetList() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}

func Test_parameterList_GetProperty(t *testing.T) {

	tests := []struct {
		name       string
		properties []properties.Property
		propertyID ids.PropertyID
		want       properties.AnyProperty
		wantErr    bool
	}{
		{"+ve Meta", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}, baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), false},
		{"+ve Mesa", []properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}, baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D")), baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), false},
		{"panic nil propertyID", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}, nil, baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), true}, // TODO: panics if propertyID is nil
		{"-ve", []properties.Property{baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D")))}, baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D")), baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))).ToAnyProperty(), false},
		{"empty properties", []properties.Property{}, baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D")), nil, false},
		{"nil properties", nil, baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D")), nil, false},
		{"multiple properties", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("demand"), baseData.NewDecData(sdkTypes.NewDec(2)))}, baseIDs.NewPropertyID(baseIDs.NewStringID("demand"), baseIDs.NewStringID("D")), baseProperties.NewMetaProperty(baseIDs.NewStringID("demand"), baseData.NewDecData(sdkTypes.NewDec(2))).ToAnyProperty(), false},
		{"property not found", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}, baseIDs.NewPropertyID(baseIDs.NewStringID("demand"), baseIDs.NewStringID("D")), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.properties...)}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("got error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := propertyList.GetProperty(tt.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameterList_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		added   lists.ParameterList
		wantErr bool
	}{
		{
			"one valid property",
			NewParameterList(randomParameters[randomParameterIndex]),
			false,
		},
		{
			"two valid properties",
			NewParameterList(randomParameters[0], randomParameters[1]),
			false,
		},
		{
			"many valid properties",
			NewParameterList(randomParameters...),
			false,
		},
		{
			"valid properties with duplicate",
			&ParameterList{[]*baseParameters.Parameter{randomParameters[randomParameterIndex].(*baseParameters.Parameter), randomParameters[randomParameterIndex].(*baseParameters.Parameter)}},
			true,
		},
		{
			"valid properties with nil",
			&ParameterList{[]*baseParameters.Parameter{randomParameters[0].(*baseParameters.Parameter), nil, randomParameters[1].(*baseParameters.Parameter)}},
			false, //nil is removed while sorting
		},
		{
			"empty property list",
			NewParameterList(),
			false,
		},
		{
			"invalid property",
			NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")))),
			true,
		},
		{
			"invalid property with valid property",
			NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a"))), randomParameters[0]),
			true,
		},
		{
			"validate a large number of properties",
			NewParameterList(veryLargeNumberOfParameters...),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.added.ValidateBasic()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterList_GetProperty(t *testing.T) {
	tests := []struct {
		name       string
		added      lists.ParameterList
		propertyID ids.PropertyID
		want       parameters.Parameter
	}{
		{
			"one property",
			NewParameterList(randomParameters[randomParameterIndex]),
			randomParameters[randomParameterIndex].GetMetaProperty().GetID(),
			randomParameters[randomParameterIndex],
		},
		{
			"two properties",
			NewParameterList(randomParameters[0], randomParameters[1]),
			randomParameters[1].GetMetaProperty().GetID(),
			randomParameters[1],
		},
		{
			"many properties",
			NewParameterList(randomParameters...),
			randomParameters[randomParameterIndex].GetMetaProperty().GetID(),
			randomParameters[randomParameterIndex],
		},
		{
			"id not in list",
			NewParameterList(randomParameters[0], randomParameters[1]),
			randomParameters[3].GetMetaProperty().GetID(),
			nil,
		},
		{
			"empty list",
			NewParameterList(),
			randomParameters[randomParameterIndex].GetMetaProperty().GetID(),
			nil,
		},
		{
			"very large number of properties",
			NewParameterList(veryLargeNumberOfParameters...),
			veryLargeNumberOfParameters[randomParameterIndex].GetMetaProperty().GetID(),
			veryLargeNumberOfParameters[randomParameterIndex],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameterList := tt.added
			if got := parameterList.GetParameter(tt.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameterList_search(t *testing.T) {
	tests := []struct {
		name      string
		added     lists.ParameterList
		parameter parameters.Parameter
		wantIndex int
		wantFound bool
	}{
		{
			"empty list",
			NewParameterList(),
			randomParameters[randomParameterIndex],
			0,
			false,
		},
		{
			"one parameter",
			NewParameterList(randomParameters[randomParameterIndex]),
			randomParameters[randomParameterIndex],
			0,
			true,
		},
		{
			"two parameters",
			NewParameterList(randomParameters[0], randomParameters[1]),
			randomParameters[1],
			1,
			true,
		},
		{
			"many parameters",
			NewParameterList(randomParameters...),
			randomParameters[randomParameterIndex],
			randomParameterIndex,
			true,
		},
		{
			"unsorted parameters",
			NewParameterList(randomUnsortedParameters...),
			randomParameters[randomParameterIndex],
			randomParameterIndex,
			true,
		},
		{
			"nil parameter",
			NewParameterList(randomParameters[0], randomParameters[1]),
			nil,
			-1,
			false,
		},
		{
			"parameter not in list",
			NewParameterList(randomParameters[0], randomParameters[1]),
			randomParameters[2],
			2,
			false,
		},
		{
			"very large number of parameters",
			NewParameterList(veryLargeNumberOfParameters...),
			veryLargeNumberOfParameters[randomParameterIndex],
			randomParameterIndex,
			true,
		},
		{
			"very large number of unsorted parameters",
			NewParameterList(veryLargeNumberOfUnsortedParameters...),
			veryLargeNumberOfParameters[randomParameterIndex],
			randomParameterIndex,
			true,
		},
		{
			"prototype parameter",
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty())),
			baseParameters.NewParameter(baseProperties.PrototypeMetaProperty()),
			0,
			true,
		},
		{
			"prototype parameter with other parameters",
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty()), randomParameters[0], randomParameters[1]),
			baseParameters.NewParameter(baseProperties.PrototypeMetaProperty()),
			2,
			true,
		},
		{
			"multiple nils",
			NewParameterList(nil, nil, nil),
			nil,
			-1,
			false,
		},
		{
			"multiple nils with other parameters",
			NewParameterList(randomParameters[0], nil, randomParameters[1]),
			nil,
			-1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameterList := tt.added
			oldParameters := make([]parameters.Parameter, len(tt.added.Get()))
			copy(oldParameters, tt.added.Get())

			gotIndex, gotFound := parameterList.(*ParameterList).search(tt.parameter)

			if gotIndex != tt.wantIndex {
				t.Errorf("search() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}

			if gotFound != tt.wantFound {
				t.Errorf("search() gotFound = %v, want %v", gotFound, tt.wantFound)
			}

			if !reflect.DeepEqual(tt.added.Get(), oldParameters) {
				t.Errorf("parameter list modified")
			}
		})
	}
}

func TestParameterList_sort(t *testing.T) {
	tests := []struct {
		name  string
		added lists.ParameterList
		want  lists.ParameterList
	}{
		{
			"empty list",
			NewParameterList(),
			NewParameterList(),
		},
		{
			"one parameter",
			NewParameterList(randomParameters[randomParameterIndex]),
			NewParameterList(randomParameters[randomParameterIndex]),
		},
		{
			"two parameters",
			NewParameterList(randomParameters[1], randomParameters[0]),
			NewParameterList(randomParameters[0], randomParameters[1]),
		},
		{
			"three parameters",
			NewParameterList(randomParameters[1], randomParameters[0], randomParameters[2]),
			NewParameterList(randomParameters[0], randomParameters[1], randomParameters[2]),
		},
		{
			"many parameters",
			NewParameterList(randomUnsortedParameters...),
			NewParameterList(randomParameters...),
		},
		{
			"many parameters already sorted",
			NewParameterList(randomParameters...),
			NewParameterList(randomParameters...),
		},
		{
			"very large number of parameters",
			NewParameterList(veryLargeNumberOfUnsortedParameters...),
			NewParameterList(veryLargeNumberOfParameters...),
		},
		{
			"very large number of parameters already sorted",
			NewParameterList(veryLargeNumberOfParameters...),
			NewParameterList(veryLargeNumberOfParameters...),
		},
		{
			"prototype parameter",
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty())),
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty())),
		},
		{
			"prototype parameter with other parameters",
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty()), testParameterGeneratorOne, testParameterGeneratorTwo),
			NewParameterList(baseParameters.NewParameter(baseProperties.PrototypeMetaProperty()), testParameterGeneratorOne, testParameterGeneratorTwo),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameterList := tt.added.(*ParameterList)

			parameterList.sort()

			if !reflect.DeepEqual(parameterList.Get(), tt.want.Get()) {
				t.Errorf("sort() got = %v, want %v", parameterList.Get(), tt.want.Get())
			}

			if reflect.TypeOf(parameterList.Get()) != reflect.TypeOf(tt.want.Get()) {
				t.Errorf("\n got type: \n %v \n want type: \n %v", reflect.TypeOf(parameterList.Get()), reflect.TypeOf(tt.want.Get()))
			}
		})
	}
}
