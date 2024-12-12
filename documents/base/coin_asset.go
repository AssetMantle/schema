// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ documents.CoinAsset = (*asset)(nil)

var coinAssetClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.DenomProperty)), baseQualified.NewMutables(baseLists.NewPropertyList()))

func (asset asset) ValidateCoinAsset() error {
	if err := asset.ValidateAsset(); err != nil {
		return err
	}

	if err := sdkTypes.ValidateDenom(asset.GetDenom()); err != nil {
		return err
	}

	if err := asset.GetCoinAssetID().ValidateBasic(); err != nil {
		return err
	}

	return nil
}
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
