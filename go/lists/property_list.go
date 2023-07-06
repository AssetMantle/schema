package lists

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/properties"
)

type PropertyList interface {
	GetProperty(ids.PropertyID) properties.AnyProperty
	Get() []properties.AnyProperty
	GetPropertyIDList() IDList

	FromMetaPropertiesString(string) (PropertyList, error)

	Add(...properties.Property) PropertyList
	Remove(...properties.Property) PropertyList
	Mutate(...properties.Property) PropertyList

	ScrubData() PropertyList
	ValidateBasic() error
}
