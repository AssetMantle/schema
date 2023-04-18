// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/schema/go/ids"
)

type IDData interface {
	Data
	Get() ids.AnyID
}
