// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Order interface {
	GetMakerID() ids.IdentityID
	GetMakerAssetID() ids.AssetID
	GetTakerAssetID() ids.AssetID
	GetMakerSplit() math.Int
	GetTakerSplit() math.Int
	GetExpiryHeight() types.Height

	GetTakerID() ids.IdentityID
	GetExchangeRate() sdkTypes.Dec
	GetExecutionHeight() types.Height

	Document
}
