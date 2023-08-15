package base

import (
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
)

var _ documents.ModuleIdentity = (*identity)(nil)

var moduleNameClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.ModuleNameProperty)), baseQualified.NewMutables(baseLists.NewPropertyList()))

func (identity identity) GetModuleName() string {
	if property := identity.Document.GetProperty(constantProperties.ModuleNameProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.StringID).Get()
	}

	return constantProperties.ModuleNameProperty.GetData().Get().(data.IDData).Get().Get().(ids.StringID).Get()
}
func (identity identity) GetModuleIdentityID() ids.IdentityID {
	return baseIDs.NewIdentityID(moduleNameClassificationID, identity.Document.GetImmutables()).(ids.IdentityID)
}

func PrototypeModuleIdentity() documents.ModuleIdentity {
	return NewModuleIdentity("")
}

func NewModuleIdentity(moduleName string) documents.ModuleIdentity {
	return NewIdentity(moduleNameClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.ModuleNameProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID(moduleName))))), baseQualified.NewMutables(baseLists.NewPropertyList())).(identity)
}
