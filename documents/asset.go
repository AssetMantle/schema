// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/types"
)

type Asset interface {
	ValidateAsset() error
	GetBurnHeight() types.Height
	GetLockHeight() types.Height
	GetSupply() sdkTypes.Int
	Document
}
