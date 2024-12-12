// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/types"
)

type Asset interface {
	ValidateAsset() error
	GetBurnHeight() types.Height
	GetLockHeight() types.Height
	GetSupply() math.Int
	Document
}
