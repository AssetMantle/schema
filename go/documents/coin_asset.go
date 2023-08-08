package documents

import "github.com/AssetMantle/schema/go/ids"

type CoinAsset interface {
	Asset
	GetDenomID() ids.StringID
}
