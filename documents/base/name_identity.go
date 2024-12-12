// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
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
	return NewNameIdentity(baseIDs.PrototypeStringID(), baseData.PrototypeListData())
}

func NewNameIdentity(name ids.StringID, authentication data.ListData) documents.NameIdentity {
	return NewIdentity(nameIdentityClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.NameProperty.GetKey(), baseData.NewIDData(name)))), baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.AuthenticationProperty.GetKey(), authentication)))).(identity)
}
