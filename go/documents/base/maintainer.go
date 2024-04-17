package base

import (
	"fmt"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
)

type maintainer struct {
	documents.Document
}

var _ documents.Maintainer = (*maintainer)(nil)

var maintainerClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.IdentityIDProperty, constantProperties.MaintainedClassificationIDProperty)), baseQualified.NewMutables(baseLists.NewPropertyList(constantProperties.MaintainedPropertiesProperty, constantProperties.PermissionsProperty)))

func (maintainer maintainer) ValidateBasic() error {
	if err := maintainer.Document.ValidateBasic(); err != nil {
		return err
	}

	if property := maintainer.GetProperty(constantProperties.IdentityIDProperty.GetID()); property == nil || !property.IsMeta() {
		return fmt.Errorf("maintainer must have a revealed %s", constantProperties.IdentityIDProperty.GetID())
	}

	if property := maintainer.GetProperty(constantProperties.MaintainedClassificationIDProperty.GetID()); property == nil || !property.IsMeta() {
		return fmt.Errorf("maintainer must have a revealed %s", constantProperties.MaintainedClassificationIDProperty.GetID())
	}

	if property := maintainer.GetProperty(constantProperties.MaintainedPropertiesProperty.GetID()); property == nil || !property.IsMeta() {
		return fmt.Errorf("maintainer must have a revealed %s", constantProperties.MaintainedPropertiesProperty.GetID())
	}

	if property := maintainer.GetProperty(constantProperties.PermissionsProperty.GetID()); property == nil || !property.IsMeta() {
		return fmt.Errorf("maintainer must have a revealed %s", constantProperties.PermissionsProperty.GetID())
	}

	return nil
}

func (maintainer maintainer) GetIdentityID() ids.IdentityID {
	if property := maintainer.GetProperty(constantProperties.IdentityIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
	}
	return constantProperties.IdentityIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
}
func (maintainer maintainer) GetMaintainedClassificationID() ids.ClassificationID {
	if property := maintainer.GetProperty(constantProperties.MaintainedClassificationIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.ClassificationID)
	}
	return constantProperties.MaintainedClassificationIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.ClassificationID)
}
func (maintainer maintainer) GetMaintainedProperties() data.ListData {
	if property := maintainer.GetProperty(constantProperties.MaintainedPropertiesProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.ListData)
	}

	return constantProperties.MaintainedPropertiesProperty.GetData().Get().(data.ListData)
}
func (maintainer maintainer) GetPermissions() data.ListData {
	if property := maintainer.GetProperty(constantProperties.PermissionsProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.ListData)
	}

	return constantProperties.PermissionsProperty.GetData().Get().(data.ListData)
}
func (maintainer maintainer) IsPermitted(permissionID ids.StringID) bool {
	_, found := maintainer.GetPermissions().Search(baseData.NewIDData(permissionID))
	return found
}
func (maintainer maintainer) MaintainsProperty(propertyID ids.PropertyID) bool {
	_, found := maintainer.GetMaintainedProperties().Search(baseData.NewIDData(propertyID))
	return found
}

// TODO: Move to a common package
func idListToListData(idList lists.IDList) data.ListData {
	dataList := baseData.NewListData()
	for _, id := range idList.GetList() {
		dataList = dataList.Add(baseData.NewIDData(id))
	}
	return dataList
}

func PrototypeMaintainer() documents.Maintainer {
	return NewMaintainer(baseIDs.PrototypeIdentityID(), baseIDs.PrototypeClassificationID(), baseLists.NewIDList(), baseLists.NewIDList())
}

func NewMaintainerFromDocument(document documents.Document) documents.Maintainer {
	return maintainer{Document: document}
}

func NewMaintainer(identityID ids.IdentityID, maintainedClassificationID ids.ClassificationID, maintainedPropertyIDList lists.IDList, permissions lists.IDList) documents.Maintainer {
	return maintainer{
		Document: NewDocument(maintainerClassificationID,
			baseQualified.NewImmutables(baseLists.NewPropertyList(
				baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(identityID)),
				baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(maintainedClassificationID)),
			)),
			baseQualified.NewMutables(baseLists.NewPropertyList(
				baseProperties.NewMetaProperty(constantProperties.MaintainedPropertiesProperty.GetKey(), idListToListData(maintainedPropertyIDList)),
				baseProperties.NewMetaProperty(constantProperties.PermissionsProperty.GetKey(), idListToListData(permissions)),
			)),
		),
	}
}
