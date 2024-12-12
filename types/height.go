// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

// Height type to represent blockchain's block height
// 0 height === chain genesis block height
// -1 height === infinite block height
type Height interface {
	Bytes() []byte
	Get() int64
	ValidateBasic() error

	// Compare returns 1 if height is greater than compareHeight
	// Compare returns 0 if height is equal to compareHeight
	// Compare returns -1 if height is less than compareHeight
	Compare(Height) int
}
