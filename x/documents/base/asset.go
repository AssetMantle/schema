package base

import (
	"github.com/AssetMantle/schema/x/documents"
	"github.com/AssetMantle/schema/x/ids"
	"github.com/AssetMantle/schema/x/properties"
	"github.com/AssetMantle/schema/x/properties/constants"
	"github.com/AssetMantle/schema/x/qualified"
)

type asset struct {
	documents.Document
}

var _ documents.Asset = (*asset)(nil)

func (asset asset) GetBurn() properties.Property {
	if burn := asset.GetProperty(constants.BurnHeightProperty.GetID()); burn != nil {
		return burn
	}

	return constants.BurnHeightProperty
}
func (asset asset) GetLock() properties.Property {
	if lock := asset.GetProperty(constants.LockProperty.GetID()); lock != nil {
		return lock
	}

	return constants.LockProperty
}
func (asset asset) GetSupply() properties.Property {
	if supply := asset.GetProperty(constants.SupplyProperty.GetID()); supply != nil {
		return supply
	}

	return constants.SupplyProperty
}

func NewAsset(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Asset {
	return asset{
		Document: NewDocument(classificationID, immutables, mutables),
	}
}

func NewAssetFromDocument(document documents.Document) documents.Asset {
	return asset{
		Document: document,
	}
}