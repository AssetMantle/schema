// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Order interface {
	GetExchangeRate() sdkTypes.Dec
	GetCreationHeight() types.Height
	GetMakerOwnableID() ids.OwnableID
	GetTakerOwnableID() ids.OwnableID
	GetMakerID() ids.IdentityID
	GetTakerID() ids.IdentityID
	GetExpiryHeight() types.Height
	GetMakerOwnableSplit() sdkTypes.Int

	Document
}
