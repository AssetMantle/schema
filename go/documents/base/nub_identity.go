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

var _ documents.NubIdentity = (*identity)(nil)

var nubIdentityClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.NubIDProperty)), baseQualified.NewMutables(baseLists.NewPropertyList(constantProperties.AuthenticationProperty)))

func (identity identity) GetNubID() ids.StringID {
	if property := identity.Document.GetProperty(constantProperties.NubIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.StringID)
	}

	return constantProperties.NubIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.StringID)
}

func PrototypeNubIdentity() documents.NubIdentity {
	return NewNubIdentity(baseIDs.PrototypeStringID(), baseData.PrototypeListData())
}

func NewNubIdentity(nubID ids.StringID, authentication data.ListData) documents.NubIdentity {
	return NewIdentity(nubIdentityClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.NubIDProperty.GetKey(), baseData.NewIDData(nubID)))), baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.AuthenticationProperty.GetKey(), authentication)))).(identity)
}
