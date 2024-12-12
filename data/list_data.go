// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

type ListData interface {
	Data
	ValidateWithoutLengthCheck() error
	Get() []AnyListableData
	Search(ListableData) (int, bool)
	// Add ignores data with type not matching list data type
	Add(...ListableData) ListData
	Remove(...ListableData) ListData
}
