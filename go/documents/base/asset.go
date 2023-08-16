package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/qualified"
	"github.com/AssetMantle/schema/go/types"
)

type asset struct {
	documents.Document
}

var _ documents.Asset = (*asset)(nil)

func (asset asset) GetBurnHeight() types.Height {
	if property := asset.GetProperty(constants.BurnHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}

	return constants.BurnHeightProperty.GetData().Get().(data.HeightData).Get()
}
func (asset asset) GetLockHeight() types.Height {
	if property := asset.GetProperty(constants.LockHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}

	return constants.LockHeightProperty.GetData().Get().(data.HeightData).Get()
}
func (asset asset) GetSupply() sdkTypes.Int {
	if property := asset.GetProperty(constants.SupplyProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	return constants.SupplyProperty.GetData().Get().(data.NumberData).Get()
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
