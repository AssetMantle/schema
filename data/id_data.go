// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/schema/ids"
)

type IDData interface {
	ListableData
	Get() ids.AnyID
}
