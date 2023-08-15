package constants

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

var (
	AccAddressDataWeight = sdkTypes.NewInt(90)
	BooleanDataWeight    = sdkTypes.NewInt(1)
	DecDataWeight        = sdkTypes.NewInt(16)
	HeightDataWeight     = sdkTypes.NewInt(8)
	IDDataWeight         = sdkTypes.NewInt(64)
	ListDataWeight       = sdkTypes.NewInt(1024)
	LinkedDataWeight     = sdkTypes.NewInt(384)
	NumberDataWeight     = sdkTypes.NewInt(8)
	StringDataWeight     = sdkTypes.NewInt(256)
)
