package documents

import "github.com/AssetMantle/schema/go/ids"

type NubIdentity interface {
	Identity
	GetNubID() ids.StringID
}
