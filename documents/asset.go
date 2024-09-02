// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/AssetMantle/schema/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Asset interface {
	ValidateAsset() error
	GetBurnHeight() types.Height
	GetLockHeight() types.Height
	GetSupply() sdkTypes.Int
	Document
}
