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

var _ documents.NameIdentity = (*identity)(nil)

var nameIdentityClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.NameProperty)), baseQualified.NewMutables(baseLists.NewPropertyList(constantProperties.AuthenticationProperty)))

func (identity identity) GetName() string {
	if property := identity.Document.GetProperty(constantProperties.NameProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.StringID).Get()
	}

	return constantProperties.NameProperty.GetData().Get().(data.IDData).Get().Get().(ids.StringID).Get()
}
func (identity identity) GetNameIdentityID() ids.IdentityID {
	return baseIDs.NewIdentityID(nameIdentityClassificationID, identity.Document.GetImmutables()).(ids.IdentityID)
}

func PrototypeNameIdentity() documents.NameIdentity {
	return NewNameIdentity("", baseData.PrototypeListData())
}

func NewNameIdentity(name string, authentication data.ListData) documents.NameIdentity {
	return NewIdentity(nameIdentityClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.NameProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID(name))))), baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.AuthenticationProperty.GetKey(), authentication)))).(identity)
}
