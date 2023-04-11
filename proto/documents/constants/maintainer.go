package constants

import (
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	constantProperties "github.com/AssetMantle/schema/x/properties/constants"
	"github.com/AssetMantle/schema/x/qualified/base"
)

var MaintainerClassificationID = baseIDs.NewClassificationID(base.NewImmutables(baseLists.NewPropertyList(constantProperties.IdentityIDProperty, constantProperties.MaintainedClassificationIDProperty)), base.NewMutables(baseLists.NewPropertyList(constantProperties.MaintainedPropertiesProperty, constantProperties.PermissionsProperty)))
