package base

import (
	"fmt"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/utilities"
	"github.com/AssetMantle/schema/properties"
	"github.com/AssetMantle/schema/properties/base"
	"sort"
)

var _ lists.PropertyList = (*PropertyList)(nil)

func (propertyList *PropertyList) ValidateBasic() error {
	propertyIDMap := map[string]bool{}

	for i, property := range propertyList.AnyProperties {
		if property == nil {
			return fmt.Errorf("nil property in list at index %d", i)
		}

		if _, found := propertyIDMap[property.GetID().AsString()]; found {
			return fmt.Errorf("duplicate property with ID: %s", property.GetID().AsString())
		} else {
			propertyIDMap[property.GetID().AsString()] = true
		}

		if err := property.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (propertyList *PropertyList) GetProperty(propertyID ids.PropertyID) properties.AnyProperty {
	propertyList.sort()

	if i, found := propertyList.search(base.NewEmptyMesaPropertyFromID(propertyID)); found {
		return propertyList.Get()[i]
	}

	return nil
}
func (propertyList *PropertyList) Get() []properties.AnyProperty {
	anyProperties := make([]properties.AnyProperty, len(propertyList.AnyProperties))

	for i, property := range propertyList.AnyProperties {
		anyProperties[i] = property
	}

	return anyProperties
}
func (propertyList *PropertyList) search(property properties.Property) (index int, found bool) {
	if property == nil {
		return -1, false
	}

	index = sort.Search(
		len(propertyList.AnyProperties),
		func(i int) bool {
			return propertyList.AnyProperties[i].Compare(property) >= 0
		},
	)

	if index < len(propertyList.AnyProperties) && propertyList.AnyProperties[index].Compare(property) == 0 {
		return index, true
	}

	return index, false
}
func (propertyList *PropertyList) sort() {
	sort.Slice(propertyList.AnyProperties, func(i, j int) bool {
		return propertyList.AnyProperties[i].Compare(propertyList.AnyProperties[j]) < 0
	})
}
func (propertyList *PropertyList) GetPropertyIDList() lists.IDList {
	propertyIDList := NewIDList()

	for _, property := range propertyList.Get() {
		propertyIDList = propertyIDList.Add(property.GetID())
	}

	return propertyIDList
}
func (propertyList *PropertyList) FromMetaPropertiesString(metaPropertiesString string) (lists.PropertyList, error) {
	for _, metaPropertyString := range utilities.SplitPropertyListString(metaPropertiesString) {
		if metaPropertyString != "" {
			metaProperty, err := base.PrototypeMetaProperty().FromString(metaPropertyString)
			if err != nil {
				return nil, err
			}

			propertyList.Add(metaProperty)
		}
	}

	return propertyList, nil
}
func (propertyList *PropertyList) Add(properties ...properties.Property) lists.PropertyList {
	propertyList.sort()

	for _, property := range properties {
		if property != nil {
			if index, found := propertyList.search(property); !found {
				propertyList.AnyProperties = append(propertyList.AnyProperties, property.ToAnyProperty().(*base.AnyProperty))
				copy(propertyList.AnyProperties[index+1:], propertyList.AnyProperties[index:])
				propertyList.AnyProperties[index] = property.ToAnyProperty().(*base.AnyProperty)
			} else {
				propertyList.AnyProperties[index] = property.ToAnyProperty().(*base.AnyProperty)
			}
		}
	}

	return propertyList
}
func (propertyList *PropertyList) Remove(properties ...properties.Property) lists.PropertyList {
	propertyList.sort()

	for _, property := range properties {
		if property != nil {
			if index, found := propertyList.search(property); found {
				propertyList.AnyProperties = append(propertyList.AnyProperties[:index], propertyList.AnyProperties[index+1:]...)
			}
		}
	}

	return propertyList
}
func (propertyList *PropertyList) Mutate(properties ...properties.Property) lists.PropertyList {
	propertyList.sort()

	for _, property := range properties {
		if property != nil {
			if index, found := propertyList.search(property); found {
				propertyList.AnyProperties[index] = property.ToAnyProperty().(*base.AnyProperty)
			}
		}
	}

	return propertyList
}
func (propertyList *PropertyList) ScrubData() lists.PropertyList {
	propertyList.sort()

	for index, property := range propertyList.AnyProperties {
		if property != nil {
			if property.IsMeta() {
				propertyList.AnyProperties[index] = property.Get().(properties.MetaProperty).ScrubData().ToAnyProperty().(*base.AnyProperty)
			}
		}
	}

	return propertyList
}
func anyPropertiesToProperties(anyProperties ...*base.AnyProperty) []properties.Property {
	returnProperties := make([]properties.Property, 0)

	for _, anyProperty := range anyProperties {
		if anyProperty != nil {
			returnProperties = append(returnProperties, anyProperty)
		}
	}

	return returnProperties
}
func propertiesToAnyProperties(properties ...properties.Property) []*base.AnyProperty {
	anyProperties := make([]*base.AnyProperty, 0)

	for _, property := range properties {
		if property != nil {
			anyProperties = append(anyProperties, property.ToAnyProperty().(*base.AnyProperty))
		}
	}

	return anyProperties
}

// AnyPropertiesToProperties converts a slice of AnyProperties to a slice of Properties.
// It filters out nil elements from the input slice and returns a new slice containing only non-nil Properties.
func AnyPropertiesToProperties(anyProperties ...properties.AnyProperty) []properties.Property {
	returnProperties := make([]properties.Property, 0)

	for _, anyProperty := range anyProperties {
		if anyProperty != nil {
			returnProperties = append(returnProperties, anyProperty.Get())
		}
	}

	return returnProperties
}

// NewPropertyList creates a new instance of PropertyList and adds the provided properties to it. It returns the created PropertyList.
func NewPropertyList(addProperties ...properties.Property) lists.PropertyList {
	return (&PropertyList{}).Add(addProperties...)
}
