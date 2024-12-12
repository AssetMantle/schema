// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

type ID interface {
	Compare(ID) int
	GetTypeID() StringID
	ValidateBasic() error
	FromString(string) (ID, error)
	AsString() string
	Bytes() []byte
	ToAnyID() AnyID
}
