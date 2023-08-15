package documents

import "github.com/AssetMantle/schema/go/ids"

type NameIdentity interface {
	Identity
	GetName() string
	GetNameIdentityID() ids.IdentityID
}
