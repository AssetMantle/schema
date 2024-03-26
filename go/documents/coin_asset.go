package documents

import "github.com/AssetMantle/schema/go/ids"

type CoinAsset interface {
	Asset
	ValidateCoinAsset() error
	GetDenom() string
	GetCoinAssetID() ids.AssetID
}
