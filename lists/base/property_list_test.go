// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/constants"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
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
	randomPropertiesGenerator = func(n int) (unsorted []*baseProperties.AnyProperty, sorted []*baseProperties.AnyProperty) {
		unsorted = make([]*baseProperties.AnyProperty, n)
		for i := 0; i < n; i++ {
			unsorted[i] = randomPropertyGenerator()
		}

		sorted = make([]*baseProperties.AnyProperty, n)
		copy(sorted, unsorted)

		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].Compare(sorted[j]) <= 0
		})

		return unsorted, sorted
	}

	maxProperties                                                    = 100
	randomIndex                                                      = rand.Intn(maxProperties)
	randomUnsortedProperties, randomProperties                       = randomPropertiesGenerator(maxProperties)
	veryLargeNumber                                                  = 50000
	veryLargeNumberOfUnsortedProperties, veryLargeNumberOfProperties = randomPropertiesGenerator(veryLargeNumber)
)

func Test_propertyList_NewPropertyList(t *testing.T) {
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
		{"add many unsorted properties",
			anyPropertiesToProperties(randomUnsortedProperties...),
			&PropertyList{randomProperties},
		},
		{"add duplicate properties",
			anyPropertiesToProperties(randomProperties[randomIndex], randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex]}},
		},
		{"add nil properties",
			anyPropertiesToProperties(nil),
			&PropertyList{},
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
			&PropertyList{},
		},
		{
			"very large number of properties",
			anyPropertiesToProperties(veryLargeNumberOfProperties...),
			&PropertyList{veryLargeNumberOfProperties},
		},
		{
			"very large number of unsorted properties",
			anyPropertiesToProperties(veryLargeNumberOfUnsortedProperties...),
			&PropertyList{veryLargeNumberOfProperties},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldProperties := make([]properties.Property, len(tt.addProperties))
			copy(oldProperties, tt.addProperties)

			if got := NewPropertyList(tt.addProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n got: \n %v \n want: \n %v", got, tt.want)
			}

			// check if input list was modified
			if !reflect.DeepEqual(tt.addProperties, oldProperties) {
				t.Errorf("input modified")
			}
		})
	}
}

func Test_propertyList_Add(t *testing.T) {
	tests := []struct {
		name       string
		added      lists.PropertyList
		properties []properties.Property
		want       lists.PropertyList
	}{
		{"add one property",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			anyPropertiesToProperties(randomProperties[2]),
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[0], randomProperties[1], randomProperties[2]}},
		},
		{"add already added property",
			NewPropertyList(randomProperties[randomIndex]),
			anyPropertiesToProperties(randomProperties[randomIndex]),
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex]}},
		},
		{"add two properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			anyPropertiesToProperties(randomProperties[2], randomProperties[3]),
			&PropertyList{randomProperties[0:4]},
		},
		{"add many properties",
			NewPropertyList(),
			anyPropertiesToProperties(randomProperties...),
			&PropertyList{randomProperties},
		},
		{"add many unsorted properties",
			NewPropertyList(),
			anyPropertiesToProperties(randomUnsortedProperties...),
			&PropertyList{randomProperties},
		},
		{"add duplicate properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			anyPropertiesToProperties(randomProperties[2], randomProperties[2]),
			&PropertyList{randomProperties[0:3]},
		},
		{"add nil properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			anyPropertiesToProperties(nil),
			&PropertyList{randomProperties[0:2]},
		},
		{"add nil property with other properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			anyPropertiesToProperties(randomProperties[2], nil, randomProperties[3]),
			&PropertyList{randomProperties[0:4]},
		},
		{
			"add properties out of order",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			anyPropertiesToProperties(randomProperties[2], randomProperties[3]),
			&PropertyList{randomProperties[0:4]},
		},
		{
			"add prototype property",
			NewPropertyList(),
			[]properties.Property{baseProperties.PrototypeMetaProperty()},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty)}},
		},
		{
			"add prototype property with other properties",
			NewPropertyList(),
			[]properties.Property{baseProperties.PrototypeMetaProperty(), randomProperties[0], randomProperties[1]},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty), randomProperties[0], randomProperties[1]}},
		},
		{
			"add multiple nils",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			[]properties.Property{nil, nil, nil},
			&PropertyList{randomProperties[0:2]},
		},
		{
			"very large number of properties",
			NewPropertyList(),
			anyPropertiesToProperties(veryLargeNumberOfProperties...),
			&PropertyList{veryLargeNumberOfProperties},
		},
		{
			"very large number of unsorted properties",
			NewPropertyList(),
			anyPropertiesToProperties(veryLargeNumberOfUnsortedProperties...),
			&PropertyList{veryLargeNumberOfProperties},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldProperties := make([]properties.Property, len(tt.properties))
			copy(oldProperties, tt.properties)

			if got := tt.added.Add(tt.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = \n %v \n want \n %v", got, tt.want)
			}

			// check if input list was modified
			if !reflect.DeepEqual(tt.properties, oldProperties) {
				t.Errorf("input modified")
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
		{"remove many unsorted properties",
			anyPropertiesToProperties(randomUnsortedProperties...),
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
			&PropertyList{},
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
			"remove multiple nils",
			[]properties.Property{nil, nil, nil},
			NewPropertyList(randomProperties[0], randomProperties[1]),
			&PropertyList{randomProperties[0:2]},
		},
		{
			"remove very large number of properties",
			anyPropertiesToProperties(veryLargeNumberOfProperties...),
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
		{
			"remove very large number of unsorted properties",
			anyPropertiesToProperties(veryLargeNumberOfUnsortedProperties...),
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			&PropertyList{[]*baseProperties.AnyProperty{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldProperties := make([]properties.Property, len(tt.removeProperties))
			copy(oldProperties, tt.removeProperties)

			if got := tt.added.Remove(tt.removeProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}

			// check if input list was modified
			if !reflect.DeepEqual(tt.removeProperties, oldProperties) {
				t.Errorf("input modified")
			}
		})
	}
}

func TestPropertyList_Mutate(t *testing.T) {
	tests := []struct {
		name       string
		added      lists.PropertyList
		properties []properties.Property
		want       lists.PropertyList
	}{
		{
			"empty list",
			NewPropertyList(),
			[]properties.Property{randomProperties[randomIndex]},
			NewPropertyList(),
		},
		{
			"one property",
			NewPropertyList(randomProperties[0]),
			[]properties.Property{randomProperties[randomIndex]},
			NewPropertyList(randomProperties[0]),
		},
		{
			"two properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			[]properties.Property{randomProperties[2], randomProperties[3]},
			NewPropertyList(randomProperties[0], randomProperties[1]),
		},
		{
			"many properties",
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			[]properties.Property{randomProperties[2], randomProperties[3]},
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
		},
		{
			"mutate one property in one property list",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(0)))),
			[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).(properties.Property)},
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).(properties.Property)),
		},
		{
			"mutate one property in two property list",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(0))), randomProperties[randomIndex]),
			[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).(properties.Property)},
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).(properties.Property), randomProperties[randomIndex]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			oldProperties := make([]properties.Property, len(tt.properties))
			copy(oldProperties, tt.properties)

			if got := propertyList.Mutate(tt.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
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

func TestPropertyList_ScrubData(t *testing.T) {
	tests := []struct {
		name  string
		added lists.PropertyList
		want  lists.PropertyList
	}{
		{
			"empty list",
			NewPropertyList(),
			NewPropertyList(),
		},
		{
			"one property",
			NewPropertyList(randomProperties[randomIndex]),
			NewPropertyList(randomProperties[randomIndex].Get().(properties.MetaProperty).ScrubData().ToAnyProperty().(*baseProperties.AnyProperty)),
		},
		{
			"two properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			NewPropertyList(randomProperties[0].Get().(properties.MetaProperty).ScrubData().ToAnyProperty().(*baseProperties.AnyProperty), randomProperties[1].Get().(properties.MetaProperty).ScrubData().ToAnyProperty().(*baseProperties.AnyProperty)),
		},
		{
			"nil properties",
			NewPropertyList(nil),
			NewPropertyList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			if got := propertyList.ScrubData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrubData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyList_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		added   lists.PropertyList
		wantErr bool
	}{
		{
			"one valid property",
			NewPropertyList(randomProperties[randomIndex]),
			false,
		},
		{
			"two valid properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			false,
		},
		{
			"many valid properties",
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			false,
		},
		{
			"valid properties with duplicate",
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex], randomProperties[randomIndex]}},
			true,
		},
		{
			"valid properties with nil",
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[0], nil, randomProperties[1]}},
			true,
		},
		{
			"empty property list",
			NewPropertyList(),
			false,
		},
		{
			"invalid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a"))),
			true,
		},
		{
			"invalid property with valid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")), randomProperties[0]),
			true,
		},
		{
			"validate a large number of properties",
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			err := propertyList.ValidateBasic()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr              bool
	}{
		{
			"one property",
			NewPropertyList(randomProperties[randomIndex]),
			MetaPropertyAsString(randomProperties[randomIndex]),
			NewPropertyList(randomProperties[randomIndex]),
			false,
		},
		{
			"two properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			MetaPropertiesAsString(randomProperties[0], randomProperties[1]),
			NewPropertyList(randomProperties[0], randomProperties[1]),
			false,
		},
		{
			"many properties",
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			MetaPropertiesAsString(anyPropertiesToProperties(randomProperties...)...),
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			false,
		},
		{
			"repeated properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			MetaPropertiesAsString(randomProperties[0], randomProperties[1], randomProperties[0]),
			NewPropertyList(randomProperties[0], randomProperties[1]),
			false,
		},
		{
			"invalid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a"))),
			MetaPropertyAsString(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")).ToAnyProperty()),
			nil,
			true,
		},
		{
			"invalid property with valid property",
			NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")), randomProperties[0]),
			MetaPropertiesAsString(baseProperties.NewMetaProperty(baseIDs.NewStringID("(invalid)"), baseData.NewStringData("a")).ToAnyProperty(), randomProperties[0]),
			nil,
			true,
		},
		//{
		//	"very large number of properties",
		//	NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
		//	MetaPropertiesAsString(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
		//	NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
		//	nil,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			got, err := propertyList.FromMetaPropertiesString(tt.metaPropertiesString)

			if (err != nil) != tt.wantErr {
				t.Errorf("FromMetaPropertiesString() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n got: \n %v \n want: \n %v", got, tt.want)
			}
		})
	}
}

func TestPropertyList_search(t *testing.T) {

	tests := []struct {
		name      string
		added     lists.PropertyList
		property  properties.Property
		wantIndex int
		wantFound bool
	}{
		{
			"empty list",
			NewPropertyList(),
			randomProperties[randomIndex],
			0,
			false,
		},
		{
			"one property",
			NewPropertyList(randomProperties[randomIndex]),
			randomProperties[randomIndex],
			0,
			true,
		},
		{
			"two properties",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			randomProperties[1],
			1,
			true,
		},
		{
			"many properties",
			NewPropertyList(anyPropertiesToProperties(randomProperties...)...),
			randomProperties[randomIndex],
			randomIndex,
			true,
		},
		{
			"unsorted properties",
			NewPropertyList(anyPropertiesToProperties(randomUnsortedProperties...)...),
			randomProperties[randomIndex],
			randomIndex,
			true,
		},
		{
			"nil property",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			nil,
			-1,
			false,
		},
		{
			"property not in list",
			NewPropertyList(randomProperties[0], randomProperties[1]),
			randomProperties[2],
			2,
			false,
		},
		{
			"very large number of properties",
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfProperties...)...),
			veryLargeNumberOfProperties[randomIndex],
			randomIndex,
			true,
		},
		{
			"very large number of unsorted properties",
			NewPropertyList(anyPropertiesToProperties(veryLargeNumberOfUnsortedProperties...)...),
			veryLargeNumberOfProperties[randomIndex],
			randomIndex,
			true,
		},
		{
			"prototype property",
			NewPropertyList(baseProperties.PrototypeMetaProperty()),
			baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty),
			0,
			true,
		},
		{
			"prototype property with other properties",
			NewPropertyList(baseProperties.PrototypeMetaProperty(), randomProperties[0], randomProperties[1]),
			baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty),
			0,
			true,
		},
		{
			"multiple nils",
			NewPropertyList(nil, nil, nil),
			nil,
			-1,
			false,
		},
		{
			"multiple nils with other properties",
			NewPropertyList(randomProperties[0], nil, randomProperties[1]),
			nil,
			-1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added
			oldProperties := make([]properties.AnyProperty, len(tt.added.Get()))
			copy(oldProperties, tt.added.Get())

			gotIndex, gotFound := propertyList.(*PropertyList).search(tt.property)

			if gotIndex != tt.wantIndex {
				t.Errorf("search() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}

			if gotFound != tt.wantFound {
				t.Errorf("search() gotFound = %v, want %v", gotFound, tt.wantFound)
			}

			if !reflect.DeepEqual(tt.added.Get(), oldProperties) {
				t.Errorf("property list modified")
			}
		})
	}
}

func TestPropertyList_sort(t *testing.T) {
	tests := []struct {
		name  string
		added lists.PropertyList
		want  lists.PropertyList
	}{
		{
			"empty list",
			NewPropertyList(),
			NewPropertyList(),
		},
		{
			"one property",
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[randomIndex]}},
			NewPropertyList(randomProperties[randomIndex]),
		},
		{
			"two properties",
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[1], randomProperties[0]}},
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[0], randomProperties[1]}},
		},
		{
			"three properties",
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[1], randomProperties[0], randomProperties[2]}},
			&PropertyList{[]*baseProperties.AnyProperty{randomProperties[0], randomProperties[1], randomProperties[2]}},
		},
		{
			"many properties",
			&PropertyList{randomUnsortedProperties},
			&PropertyList{randomProperties},
		},
		{
			"many properties already sorted",
			&PropertyList{randomProperties},
			&PropertyList{randomProperties},
		},
		{
			"very large number of properties",
			&PropertyList{veryLargeNumberOfUnsortedProperties},
			&PropertyList{veryLargeNumberOfProperties},
		},
		{
			"very large number of properties already sorted",
			&PropertyList{veryLargeNumberOfProperties},
			&PropertyList{veryLargeNumberOfProperties},
		},
		{
			"prototype property",
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty)}},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty)}},
		},
		{
			"prototype property with other properties",
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty), randomProperties[0], randomProperties[1]}},
			&PropertyList{[]*baseProperties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty), randomProperties[0], randomProperties[1]}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := tt.added.(*PropertyList)

			propertyList.sort()

			if !reflect.DeepEqual(tt.added.Get(), tt.want.Get()) {
				t.Errorf("sort() got = %v, want %v", tt.added.Get(), tt.want.Get())
			}
		})
	}
}

func TestAnyPropertiesToProperties(t *testing.T) {
	tests := []struct {
		name          string
		anyProperties []properties.AnyProperty
		want          []properties.Property
	}{
		{
			"empty list",
			[]properties.AnyProperty{},
			[]properties.Property{},
		},
		{
			"one property",
			[]properties.AnyProperty{randomProperties[randomIndex]},
			[]properties.Property{randomProperties[randomIndex].Get()},
		},
		{
			"two properties",
			[]properties.AnyProperty{randomProperties[0], randomProperties[1]},
			[]properties.Property{randomProperties[0].Get(), randomProperties[1].Get()},
		},
		{
			"many properties",
			[]properties.AnyProperty{randomProperties[0], randomProperties[1], randomProperties[2]},
			[]properties.Property{randomProperties[0].Get(), randomProperties[1].Get(), randomProperties[2].Get()},
		},
		{
			"nil properties",
			[]properties.AnyProperty{nil},
			[]properties.Property{},
		},
		{
			"nil property with other properties",
			[]properties.AnyProperty{randomProperties[0], nil, randomProperties[1]},
			[]properties.Property{randomProperties[0].Get(), randomProperties[1].Get()},
		},
		{
			"prototype property",
			[]properties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty)},
			[]properties.Property{baseProperties.PrototypeMetaProperty()},
		},
		{
			"prototype property with other properties",
			[]properties.AnyProperty{baseProperties.PrototypeMetaProperty().ToAnyProperty().(*baseProperties.AnyProperty), randomProperties[0], randomProperties[1]},
			[]properties.Property{baseProperties.PrototypeMetaProperty(), randomProperties[0].Get(), randomProperties[1].Get()},
		},
		{
			"multiple nils",
			[]properties.AnyProperty{nil, nil, nil},
			[]properties.Property{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldAnyProperties := make([]properties.AnyProperty, len(tt.anyProperties))
			copy(oldAnyProperties, tt.anyProperties)

			if got := AnyPropertiesToProperties(tt.anyProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n got: \n %v \n want: \n %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.anyProperties, oldAnyProperties) {
				t.Errorf("input modified")
			}
		})
	}
}
