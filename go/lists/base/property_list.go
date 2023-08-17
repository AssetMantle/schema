package base

import (
	"sort"

	"github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/lists/utilities"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/properties/base"
)

var _ lists.PropertyList = (*PropertyList)(nil)

func (propertyList *PropertyList) ValidateBasic() error {
	propertyIDMap := map[string]bool{}

	for _, property := range propertyList.AnyProperties {

		if _, found := propertyIDMap[property.GetID().AsString()]; found {
			return constants.IncorrectFormat.Wrapf("duplicate property ID: %s", property.GetID().AsString())
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
	updatedList := NewPropertyList(anyPropertiesToProperties(propertyList.Get()...)...).(*PropertyList)
	if i, found := updatedList.search(base.NewEmptyMesaPropertyFromID(propertyID)); found {
		return updatedList.Get()[i]
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

// NOTE: No need to sort since it's an internal function, function calling it is already sorting it
func (propertyList *PropertyList) search(property properties.Property) (index int, found bool) {
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
func (propertyList *PropertyList) GetPropertyIDList() lists.IDList {
	propertyIDList := NewIDList()

	for _, property := range propertyList.Get() {
		propertyIDList = propertyIDList.Add(property.GetID())
	}

	return propertyIDList
}
func (propertyList *PropertyList) FromMetaPropertiesString(metaPropertiesString string) (lists.PropertyList, error) {
	var Properties []properties.Property

	metaProperties := utilities.SplitPropertyListString(metaPropertiesString)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, err := base.PrototypeMetaProperty().FromString(metaPropertyString)
			if err != nil {
				return nil, err
			}

			Properties = append(Properties, metaProperty)
		}
	}

	updatedPropertyList := NewPropertyList(Properties...)
	if err := updatedPropertyList.ValidateBasic(); err != nil {
		return nil, err
	}

	return updatedPropertyList, nil
}
func (propertyList *PropertyList) Add(properties ...properties.Property) lists.PropertyList {
	updatedList := NewPropertyList(anyPropertiesToProperties(propertyList.Get()...)...).(*PropertyList)

	for _, property := range properties {
		if index, found := updatedList.search(property); !found {
			updatedList.AnyProperties = append(updatedList.AnyProperties, property.ToAnyProperty().(*base.AnyProperty))
			copy(updatedList.AnyProperties[index+1:], updatedList.AnyProperties[index:])
			updatedList.AnyProperties[index] = property.ToAnyProperty().(*base.AnyProperty)
		} else {
			updatedList.AnyProperties[index] = property.ToAnyProperty().(*base.AnyProperty)
		}
	}

	return updatedList
}
func (propertyList *PropertyList) Remove(properties ...properties.Property) lists.PropertyList {
	updatedList := NewPropertyList(anyPropertiesToProperties(propertyList.Get()...)...).(*PropertyList)

	for _, listable := range properties {
		if index, found := updatedList.search(listable); found {
			updatedList.AnyProperties = append(updatedList.AnyProperties[:index], updatedList.AnyProperties[index+1:]...)
		}
	}

	return updatedList
}
func (propertyList *PropertyList) Mutate(properties ...properties.Property) lists.PropertyList {
	updatedList := NewPropertyList(anyPropertiesToProperties(propertyList.Get()...)...).(*PropertyList)

	for _, listable := range properties {
		if index, found := updatedList.search(listable); found {
			updatedList.AnyProperties[index] = listable.ToAnyProperty().(*base.AnyProperty)
		}
	}

	return updatedList
}
func (propertyList *PropertyList) ScrubData() lists.PropertyList {
	updatedList := NewPropertyList(anyPropertiesToProperties(propertyList.Get()...)...).(*PropertyList)

	for index, property := range updatedList.AnyProperties {
		if property.IsMeta() {
			updatedList.AnyProperties[index] = property.Get().(properties.MetaProperty).ScrubData().ToAnyProperty().(*base.AnyProperty)
		}
	}

	return updatedList
}
func anyPropertiesToProperties(anyProperties ...properties.AnyProperty) []properties.Property {
	returnProperties := make([]properties.Property, len(anyProperties))

	for i, anyProperty := range anyProperties {
		returnProperties[i] = anyProperty
	}

	return returnProperties
}
func propertiesToAnyProperties(properties ...properties.Property) []*base.AnyProperty {
	anyProperties := make([]*base.AnyProperty, len(properties))
	for i, property := range properties {
		anyProperties[i] = property.ToAnyProperty().(*base.AnyProperty)
	}
	return anyProperties
}
func AnyPropertiesToProperties(anyProperties ...properties.AnyProperty) []properties.Property {
	returnProperties := make([]properties.Property, len(anyProperties))

	for i, anyProperty := range anyProperties {
		returnProperties[i] = anyProperty
	}

	return returnProperties
}

func NewPropertyList(properties ...properties.Property) lists.PropertyList {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i].Compare(properties[j]) <= 0
	})

	return &PropertyList{propertiesToAnyProperties(properties...)}
}

func PrototypePropertyList() lists.PropertyList {
	return &PropertyList{}
}
