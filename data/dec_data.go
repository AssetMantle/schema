// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type DecData interface {
	ListableData
	Get() sdkTypes.Dec
}
