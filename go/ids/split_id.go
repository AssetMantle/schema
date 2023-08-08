package ids

type SplitID interface {
	ID
	MustGetFromPrefixedStoreKeyBytes(prefixBytes, storeKeyBytes []byte) SplitID
	GetAssetID() AssetID
	GetOwnerID() IdentityID
	IsSplitID()
}
