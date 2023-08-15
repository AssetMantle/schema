package documents

import "github.com/AssetMantle/schema/go/ids"

type CoinAsset interface {
	Asset
	GetDenom() string
	GetCoinAssetID() ids.AssetID
}
