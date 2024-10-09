package base

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/properties"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified"
	"github.com/AssetMantle/schema/types"
)

type asset struct {
	documents.Document
}

var _ documents.Asset = (*asset)(nil)

func (asset asset) ValidateAsset() error {
	if err := asset.ValidateBasic(); err != nil {
		return err
	}

	if err := asset.GetBurnHeight().ValidateBasic(); err != nil {
		return err
	}

	if err := asset.GetLockHeight().ValidateBasic(); err != nil {
		return err
	}

	if asset.GetSupply().IsNegative() || asset.GetSupply().IsZero() {
		return fmt.Errorf("supply %s is negative", asset.GetSupply().String())
	}

	return nil
}
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
func (asset asset) GetSupply() math.Int {
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
