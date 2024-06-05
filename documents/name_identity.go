package documents

import "github.com/AssetMantle/schema/ids"

type NameIdentity interface {
	Identity
	GetName() string
	GetNameIdentityID() ids.IdentityID
}
