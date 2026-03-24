// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"cosmossdk.io/math"
)

type DecData interface {
	ListableData
	Get() math.LegacyDec
}
