package lists

import (
	"github.com/AssetMantle/schema/x/ids"
	"github.com/AssetMantle/schema/x/properties"
)

type PropertyList interface {
	GetProperty(ids.PropertyID) properties.AnyProperty
	GetList() []properties.AnyProperty
	GetPropertyIDList() IDList

	Add(...properties.Property) PropertyList
	Remove(...properties.Property) PropertyList
	Mutate(...properties.Property) PropertyList

	ScrubData() PropertyList
	ValidateBasic() error
}
