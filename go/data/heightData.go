// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/schema/x/types"
)

type HeightData interface {
	Data
	Get() types.Height
}
