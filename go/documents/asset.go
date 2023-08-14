// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/types"
)

type Asset interface {
	GetBurnHeight() types.Height
	GetLockHeight() types.Height
	GetSupply() sdkTypes.Int
	Document
}
