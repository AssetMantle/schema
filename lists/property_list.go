package lists

import (
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/properties"
)

// PropertyList represents a list of properties.
// It provides methods for retrieving, adding, removing, and mutating properties in the list.
type PropertyList interface {

	// GetProperty searches for a property within the PropertyList using a given PropertyID.
	// If the property is found, it will return the property as output.
	// If the property is not found, the function will return `nil`.
	GetProperty(ids.PropertyID) properties.AnyProperty

	// Get returns all properties in the PropertyList as a slice of AnyProperty.
	Get() []properties.AnyProperty

	// GetPropertyIDList returns a list of all IDs in the PropertyList as an IDList.
	GetPropertyIDList() IDList

	// FromMetaPropertiesString tries to parse a provided string into a PropertyList.
	// If parsing is successful, the function returns the created PropertyList and no error.
	// If parsing fails, the function returns `nil` and the corresponding error.
	FromMetaPropertiesString(string) (PropertyList, error)

	// Add will add one or more new properties to the PropertyList.
	// If a property already exists in the list, the property is updated with the new value.
	// The function will return the updated PropertyList.
	Add(...properties.Property) PropertyList

	// Remove will remove one or more existing properties from the PropertyList.
	// The function will return the updated PropertyList.
	Remove(...properties.Property) PropertyList

	// Mutate will update one or more properties in the PropertyList with new values.
	// The function will return the updated PropertyList.
	Mutate(...properties.Property) PropertyList

	// ScrubData is used to convert all the metadata in the PropertyList to mesadata.
	// The function will return the updated PropertyList.
	ScrubData() PropertyList

	// ValidateBasic will check the PropertyList for basic validity.
	// This could involve checking if all properties have valid IDs, or if all required properties are present.
	// If the PropertyList is valid, the function returns no error.
	// If the PropertyList is not valid, the function returns the corresponding error.
	ValidateBasic() error
}
