package base

import (
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
)

var _ documents.CoinAsset = (*asset)(nil)

var coinAssetClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.DenomIDProperty)), baseQualified.NewMutables(baseLists.NewPropertyList()))

func (asset asset) GetDenomID() ids.StringID {
	if property := asset.Document.GetProperty(constantProperties.DenomIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.StringID)
	}

	return constantProperties.DenomIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.StringID)
}

func PrototypeCoinAsset() documents.CoinAsset {
	return NewCoinAsset(baseIDs.PrototypeStringID())
}

func NewCoinAsset(denomID ids.StringID) documents.CoinAsset {
	return NewAsset(coinAssetClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.DenomIDProperty.GetKey(), baseData.NewIDData(denomID)))), baseQualified.NewMutables(baseLists.NewPropertyList())).(asset)
}

func GenerateCoinAssetID(denom string) ids.AssetID {
	coinAsset := NewCoinAsset(baseIDs.NewStringID(denom))
	return baseIDs.NewAssetID(coinAsset.GetClassificationID(), coinAsset.GetImmutables())
}
