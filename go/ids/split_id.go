package ids

type SplitID interface {
	ID
	GetOwnableID() OwnableID
	GetOwnerID() IdentityID
	IsSplitID()
}
