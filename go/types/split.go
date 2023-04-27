// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/schema/go/ids"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Split interface {
	GetOwnerID() ids.IdentityID
	GetOwnableID() ids.OwnableID
	GetValue() sdkTypes.Int
	CanSend(sdkTypes.Int) bool

	Send(sdkTypes.Int) Split
	Receive(sdkTypes.Int) Split
	ValidateBasic() error
}
