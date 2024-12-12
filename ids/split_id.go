// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

type SplitID interface {
	ID
	MustGetFromPrefixedStoreKeyBytes(prefixBytes, storeKeyBytes []byte) SplitID
	GetAssetID() AssetID
	GetOwnerID() IdentityID
	IsSplitID()
}
