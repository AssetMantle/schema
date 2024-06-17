package data

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type NumberData interface {
	ListableData
	Get() sdkTypes.Int
}
