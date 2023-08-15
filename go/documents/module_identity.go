package documents

import "github.com/AssetMantle/schema/go/ids"

type ModuleIdentity interface {
	Identity
	GetModuleName() string
	GetModuleIdentityID() ids.IdentityID
}
