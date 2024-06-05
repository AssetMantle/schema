// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Split interface {
	GetValue() sdkTypes.Int
	CanSend(sdkTypes.Int) bool

	Subtract(sdkTypes.Int) Split
	Add(sdkTypes.Int) Split
	ValidateBasic() error
}
