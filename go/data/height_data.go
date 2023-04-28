// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/schema/go/types"
)

type HeightData interface {
	ListableData
	Get() types.Height
}
