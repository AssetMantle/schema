package ids

type SplitID interface {
	ID
	MustGetFromPrefixedStoreKeyBytes(prefixBytes, storeKeyBytes []byte) SplitID
	GetOwnableID() OwnableID
	GetOwnerID() IdentityID
	IsSplitID()
}
