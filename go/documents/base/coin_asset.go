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

var coinAssetClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.DenomProperty)), baseQualified.NewMutables(baseLists.NewPropertyList()))

func (asset asset) GetDenom() string {
	if property := asset.Document.GetProperty(constantProperties.DenomProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.StringData).Get()
	}

	return constantProperties.DenomProperty.GetData().Get().(data.StringData).Get()
}
func (asset asset) GetCoinAssetID() ids.AssetID {
	return baseIDs.NewAssetID(coinAssetClassificationID, asset.Document.GetImmutables())
}

func PrototypeCoinAsset() documents.CoinAsset {
	return NewCoinAsset("")
}

func NewCoinAsset(denom string) documents.CoinAsset {
	return NewAsset(coinAssetClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.DenomProperty.GetKey(), baseData.NewStringData(denom)))), baseQualified.NewMutables(baseLists.NewPropertyList())).(asset)
}
