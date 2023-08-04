// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/types"
)

type Order interface {
	GetMakerID() ids.IdentityID
	GetMakerOwnableID() ids.OwnableID
	GetTakerOwnableID() ids.OwnableID
	GetMakerOwnableSplit() sdkTypes.Int
	GetTakerOwnableSplit() sdkTypes.Int
	GetExpiryHeight() types.Height

	GetTakerID() ids.IdentityID
	GetExchangeRate() sdkTypes.Dec
	GetExecutionHeight() types.Height

	Document
}
