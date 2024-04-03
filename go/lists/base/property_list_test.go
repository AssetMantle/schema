// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/errors"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/lists/constants"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var (
	randomPropertyGenerator = func() *baseProperties.AnyProperty {
		return baseProperties.NewMetaProperty(baseIDs.NewStringID(strconv.Itoa(rand.Intn(99999999999999999))), baseData.NewStringData(strconv.Itoa(rand.Intn(rand.Intn(99999999999999999))))).ToAnyProperty().(*baseProperties.AnyProperty)
	}
	randomPropertiesGenerator = func(n int) []*baseProperties.AnyProperty {
		properties := make([]*baseProperties.AnyProperty, n)
		for i := 0; i < n; i++ {
			properties[i] = randomPropertyGenerator()
		}

		sort.Slice(properties, func(i, j int) bool {
			return properties[i].Compare(properties[j]) <= 0
		})

		return properties
	}
	maxProperties               = 100
	randomIndex                 = rand.Intn(maxProperties)
	randomProperties            = randomPropertiesGenerator(maxProperties)
	veryLargeNumber             = 5000
	veryLargeNumberOfProperties = randomPropertiesGenerator(veryLargeNumber)
)

func Test_propertyList_Add(t *testing.T) {
	tests := []struct {
		name       string
		properties []properties.Property
		want       lists.PropertyList
	}{
		{"add one property",
			anyPropertiesToProperties(randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex]}},
		},
		{"add two properties",
			anyPropertiesToProperties(randomProperties[0], randomProperties[1]),
			&PropertyList{randomProperties[0:2]},
		},
		{"add many properties",
			anyPropertiesToProperties(randomProperties...),
			&PropertyList{randomProperties},
		},
		{"add duplicate properties",
			anyPropertiesToProperties(randomProperties[randomIndex], randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex]}},
		},
		{"add nil properties",
			anyPropertiesToProperties(nil),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{"add nil property with other properties",
			anyPropertiesToProperties(randomProperties[0], nil, randomProperties[1]),
			&PropertyList{randomProperties[0:2]},
		},
		{
			"add properties out of order",
			anyPropertiesToProperties(randomProperties[1], randomProperties[0]),
			&PropertyList{randomProperties[0:2]},
		},
		{
			"add prototype property",
			[]properties.Property{baseProperties.PrototypeMetaProperty()},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty)}},
		},
		{
			"add prototype property with other properties",
			[]properties.Property{baseProperties.PrototypeMetaProperty(), randomProperties[0], randomProperties[1]},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty), randomProperties[0], randomProperties[1]}},
		},
		{
			"add multiple nils",
			[]properties.Property{nil, nil, nil},
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"very large number of properties",
			anyPropertiesToProperties(veryLargeNumberOfProperties...),
			&PropertyList{veryLargeNumberOfProperties},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldProperties := make([]properties.Property, len(tt.properties))
			copy(oldProperties, tt.properties)

			if got := NewPropertyList().Add(tt.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = \n %v \n want \n %v", got, tt.want)
			}

			// check if input list was modified
			if !reflect.DeepEqual(tt.properties, oldProperties) {
				t.Errorf("input modified")
			}
		})
	}
}

func Test_propertyList_GetList(t *testing.T) {

	tests := []struct {
		name       string
		properties []properties.Property
		want       []properties.Property
	}{
		{"+ve", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}},
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

func Test_propertyList_GetProperty(t *testing.T) {

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.properties...)}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("&StringIDFromInterface() error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := propertyList.GetProperty(tt.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetPropertyIDList(t *testing.T) {

	tests := []struct {
		name       string
		properties []properties.Property
		want       lists.IDList
	}{
		{"+ve", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}, NewIDList().Add(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).GetID())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.properties...)}
			if got := propertyList.GetPropertyIDList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPropertyIDList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Mutate(t *testing.T) {

	propertyList1 := &PropertyList{propertiesToAnyProperties([]properties.Property{
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewNumberData(sdkTypes.OneInt())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID3"), baseData.NewHeightData(baseTypes.NewHeight(10))),
	}...)}
	propertyListMutated := &PropertyList{propertiesToAnyProperties([]properties.Property{
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("DataUpdated")),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewNumberData(sdkTypes.OneInt())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("ID3"), baseData.NewHeightData(baseTypes.NewHeight(10))),
	}...)}
	tests := []struct {
		name   string
		fields lists.PropertyList
		args   []properties.Property
		want   lists.PropertyList
	}{
		{"+ve", propertyList1, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("DataUpdated"))}, propertyListMutated},
		{"+ve", propertyList1, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("DataUpdated"))}, propertyList1},
		{"+ve", propertyList1, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewNumberData(sdkTypes.OneInt()))}, propertyList1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Mutate(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("propertyList_Mutate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Remove(t *testing.T) {
	tests := []struct {
		name       string
		properties []properties.Property
		remove     []properties.Property
		want       lists.PropertyList
	}{
		{"+ve", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply1"), baseData.NewDecData(sdkTypes.NewDec(2)))}, []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply1"), baseData.NewDecData(sdkTypes.NewDec(2)))}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.properties...)}
			if got := propertyList.Remove(tt.remove...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_ScrubData(t *testing.T) {

	tests := []struct {
		name       string
		properties []properties.Property
		want       lists.PropertyList
	}{
		{"+ve", []properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ScrubData()}...)},
		{"+ve", []properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Test"))}, NewPropertyList([]properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Test")).ScrubData()}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := PropertyList{propertiesToAnyProperties(tt.properties...)}
			if got := propertyList.ScrubData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrubData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPropertyList(t *testing.T) {
	tests := []struct {
		name          string
		addProperties []properties.Property
		want          lists.PropertyList
	}{
		{"add one property",
			anyPropertiesToProperties(randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex]}},
		},
		{"add two properties",
			anyPropertiesToProperties(randomProperties[0], randomProperties[1]),
			&PropertyList{randomProperties[0:2]},
		},
		{"add many properties",
			anyPropertiesToProperties(randomProperties...),
			&PropertyList{randomProperties},
		},
		{"add duplicate properties",
			anyPropertiesToProperties(randomProperties[randomIndex], randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex]}},
		},
		{"add nil properties",
			anyPropertiesToProperties(nil),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{"add nil property with other properties",
			anyPropertiesToProperties(randomProperties[0], nil, randomProperties[1]),
			&PropertyList{randomProperties[0:2]},
		},
		{
			"add properties out of order",
			anyPropertiesToProperties(randomProperties[1], randomProperties[0]),
			&PropertyList{randomProperties[0:2]},
		},
		{
			"add prototype property",
			[]properties.Property{baseProperties.PrototypeMetaProperty()},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty)}},
		},
		{
			"add prototype property with other properties",
			[]properties.Property{baseProperties.PrototypeMetaProperty(), randomProperties[0], randomProperties[1]},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty), randomProperties[0], randomProperties[1]}},
		},
		{
			"add multiple nils",
			[]properties.Property{nil, nil, nil},
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"very large number of properties",
			anyPropertiesToProperties(veryLargeNumberOfProperties...),
			&PropertyList{veryLargeNumberOfProperties},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPropertyList(tt.addProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyList_Remove(t *testing.T) {
	tests := []struct {
		name             string
		removeProperties []properties.Property
		added            lists.PropertyList
		want             lists.PropertyList
	}{
		{"remove one property",
			anyPropertiesToProperties(randomProperties[randomIndex]),
			NewPropertyList(randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{"remove two properties",
			anyPropertiesToProperties(randomProperties[0], randomProperties[1]),
			NewPropertyList(randomProperties[0], randomProperties[1]),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{"remove many properties",
			anyPropertiesToProperties(randomProperties...),
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{"remove duplicate properties",
			anyPropertiesToProperties(randomProperties[randomIndex], randomProperties[randomIndex]),
			NewPropertyList(randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{"remove nil properties",
			anyPropertiesToProperties(nil),
			NewPropertyList(),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{"remove nil property with other properties",
			anyPropertiesToProperties(randomProperties[0], nil, randomProperties[1]),
			NewPropertyList(anyPropertiesToProperties(randomProperties[0:2]...)...),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"remove properties out of order",
			anyPropertiesToProperties(randomProperties[1], randomProperties[0]),
			NewPropertyList(anyPropertiesToProperties(randomProperties[0:2]...)...),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"remove prototype property",
			[]properties.Property{baseProperties.PrototypeMetaProperty()},
			NewPropertyList(baseProperties.PrototypeMetaProperty()),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"remove prototype property with other properties",
			[]properties.Property{baseProperties.PrototypeMetaProperty(), randomProperties[0], randomProperties[1]},
			NewPropertyList(baseProperties.PrototypeMetaProperty(), randomProperties[0], randomProperties[1]),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"add multiple nils",
			[]properties.Property{nil, nil, nil},
			NewPropertyList(),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"remove very large number of properties",
			anyPropertiesToProperties(veryLargeNumberOfProperties...),
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			if got := propertyList.Remove(tt.removeProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyList_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		added   lists.PropertyList
		wantErr errors.Error
	}{
		{
			"one valid property",
			NewPropertyList(randomProperties[randomIndex]),
			nil,
		},
		{
			"two valid properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			nil,
		},
		{
			"many valid properties",
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			nil,
		},
		{
			"valid properties with duplicate",
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex], randomProperties[randomIndex]}},
			errorConstants.IncorrectFormat,
		},
		{
			"valid properties with nil",
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[0], nil, randomProperties[1]}},
			errorConstants.IncorrectFormat,
		},
		{
			"empty property list",
			NewPropertyList(),
			nil,
		},
		{
			"invalid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a"))),
			errorConstants.IncorrectFormat,
		},
		{
			"invalid property with valid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")), randomProperties[0]),
			errorConstants.IncorrectFormat,
		},
		{
			"validate a large number of properties",
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			err := propertyList.ValidateBasic()
			if err == nil && tt.wantErr != nil || err != nil && tt.wantErr == nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("\n got: \n %v \n want: \n %v", err, tt.wantErr)
			}
		})
	}
}

func TestPropertyList_GetProperty(t *testing.T) {
	tests := []struct {
		name       string
		added      lists.PropertyList
		propertyID ids.PropertyID
		want       properties.AnyProperty
	}{
		{
			"one property",
			NewPropertyList(randomProperties[randomIndex]),
			randomProperties[randomIndex].GetID(),
			randomProperties[randomIndex],
		},
		{
			"two properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			randomProperties[1].GetID(),
			randomProperties[1],
		},
		{
			"many properties",
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			randomProperties[randomIndex].GetID(),
			randomProperties[randomIndex],
		},
		{
			"id not in list",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			randomProperties[3].GetID(),
			nil,
		},
		{
			"empty list",
			NewPropertyList(),
			randomProperties[randomIndex].GetID(),
			nil,
		},
		{
			"very large number of properties",
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			veryLargeNumberOfProperties[randomIndex].GetID(),
			veryLargeNumberOfProperties[randomIndex],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			if got := propertyList.GetProperty(tt.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO see placement of this method
func MetaPropertyAsString(anyProperty properties.Property) string {
	metaProperty := anyProperty.(*baseProperties.AnyProperty).Get().(properties.MetaProperty)
	return strings.ReplaceAll(metaProperty.GetID().AsString(), ".S", "") + ":" + metaProperty.GetData().AsString()
}

func MetaPropertiesAsString(properties ...properties.Property) string {
	metaProperties := make([]string, len(properties))
	for i, property := range properties {
		metaProperties[i] = MetaPropertyAsString(property.(*baseProperties.AnyProperty))
	}
	return strings.Join(metaProperties, constants.PropertyListStringSeparator)
}

func TestPropertyList_FromMetaPropertiesString(t *testing.T) {
	tests := []struct {
		name                 string
		added                lists.PropertyList
		metaPropertiesString string
		want                 lists.PropertyList
		wantErr              errors.Error
	}{
		{
			"one property",
			NewPropertyList(randomProperties[randomIndex]),
			MetaPropertyAsString(randomProperties[randomIndex]),
			NewPropertyList(randomProperties[randomIndex]),
			nil,
		},
		{
			"two properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			MetaPropertiesAsString(randomProperties[0], randomProperties[1]),
			NewPropertyList(randomProperties[0], randomProperties[1]),
			nil,
		},
		{
			"many properties",
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			MetaPropertiesAsString(anyPropertiesToProperties(randomProperties...)...),
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			nil,
		},
		{
			"repeated properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			MetaPropertiesAsString(randomProperties[0], randomProperties[1], randomProperties[0]),
			NewPropertyList(randomProperties[0], randomProperties[1]),
			nil,
		},
		{
			"invalid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a"))),
			MetaPropertyAsString(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")).ToAnyProperty()),
			nil,
			errorConstants.IncorrectFormat,
		},
		{
			"invalid property with valid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")), randomProperties[0]),
			MetaPropertiesAsString(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")).ToAnyProperty(), randomProperties[0]),
			nil,
			errorConstants.IncorrectFormat,
		},
		{
			"very large number of properties",
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			MetaPropertiesAsString(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			got, err := propertyList.FromMetaPropertiesString(tt.metaPropertiesString)

			if err == nil && tt.wantErr != nil || err != nil && tt.wantErr == nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("\n got: \n %v \n want: \n %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n got: \n %v \n want: \n %v", got, tt.want)
			}
		})
	}
}
