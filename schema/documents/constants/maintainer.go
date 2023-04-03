package constants

import (
	baseIDs "github.com/AssetMantle/schema/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/schema/lists/base"
	constantProperties "github.com/AssetMantle/schema/schema/properties/constants"
	"github.com/AssetMantle/schema/schema/qualified/base"
)

var MaintainerClassificationID = baseIDs.NewClassificationID(base.NewImmutables(baseLists.NewPropertyList(constantProperties.IdentityIDProperty, constantProperties.MaintainedClassificationIDProperty)), base.NewMutables(baseLists.NewPropertyList(constantProperties.MaintainedPropertiesProperty, constantProperties.PermissionsProperty)))
