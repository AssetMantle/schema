package data

import "github.com/AssetMantle/schema/go/traits"

type ListableData interface {
	ToAnyListableData() AnyListableData
	Data
	traits.Listable
}
