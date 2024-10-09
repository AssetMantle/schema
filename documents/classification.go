// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import "cosmossdk.io/math"

type Classification interface {
	GetBondAmount() math.Int
	Document
}
