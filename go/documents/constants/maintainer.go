package constants

import (
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/qualified/base"
)

var MaintainerClassificationID = baseIDs.NewClassificationID(base.NewImmutables(baseLists.NewPropertyList(constantProperties.IdentityIDProperty, constantProperties.MaintainedClassificationIDProperty)), base.NewMutables(baseLists.NewPropertyList(constantProperties.MaintainedPropertiesProperty, constantProperties.PermissionsProperty)))
