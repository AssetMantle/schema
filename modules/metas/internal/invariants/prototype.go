package invariants

import (
	"github.com/AssetMantle/modules/modules/metas/internal/invariants/dummy"
	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Invariants {
	return base.NewInvariants(module.Name, "", dummy.Invariant)
}