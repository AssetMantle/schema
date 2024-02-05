package lists

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/properties"
)

type PropertyList interface {
	// GetProperty returns the property with the given PropertyID
	// returns nil if not found
	GetProperty(ids.PropertyID) properties.AnyProperty
	Get() []properties.AnyProperty
	GetPropertyIDList() IDList

	FromMetaPropertiesString(string) (PropertyList, error)

	// Add will add a new property or update if already present
	Add(...properties.Property) PropertyList
	Remove(...properties.Property) PropertyList
	Mutate(...properties.Property) PropertyList

	ScrubData() PropertyList
	ValidateBasic() error
}
