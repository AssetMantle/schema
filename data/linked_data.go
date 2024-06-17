package data

import (
	"github.com/AssetMantle/schema/ids"
)

type LinkedData interface {
	ListableData
	GetResourceID() ids.HashID
	GetExtensionID() ids.StringID
	GetServiceEndpoint() string
}
