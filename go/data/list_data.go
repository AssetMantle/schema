// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import "github.com/AssetMantle/schema/go/ids"

type ListData interface {
	Data
	ValidateWithType(expectedTypeID ids.StringID) error
	Get() []ListableData
	Search(ListableData) (int, bool)
	Add(...ListableData) ListData
	Remove(...ListableData) ListData
}
