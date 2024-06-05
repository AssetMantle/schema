package documents

import "github.com/AssetMantle/schema/ids"

type ModuleIdentity interface {
	Identity
	GetModuleName() string
	GetModuleIdentityID() ids.IdentityID
}
