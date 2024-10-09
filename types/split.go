// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"cosmossdk.io/math"
)

type Split interface {
	GetValue() math.Int
	CanSend(math.Int) bool

	Subtract(math.Int) Split
	Add(math.Int) Split
	ValidateBasic() error
}
