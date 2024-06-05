package documents

import "github.com/AssetMantle/schema/ids"

type CoinAsset interface {
	Asset
	ValidateCoinAsset() error
	GetDenom() string
	GetCoinAssetID() ids.AssetID
}
