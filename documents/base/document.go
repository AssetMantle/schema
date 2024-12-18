// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	"github.com/AssetMantle/schema/qualified"
	"github.com/AssetMantle/schema/qualified/base"
)

var _ documents.Document = (*Document)(nil)

func (document *Document) ValidateBasic() error {
	if err := document.ClassificationID.ValidateBasic(); err != nil {
		return err
	}

	if err := baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(append(document.GetImmutables().GetImmutablePropertyList().Get(), document.GetMutables().GetMutablePropertyList().Get()...)...)...).ValidateBasic(); err != nil {
		return err
	}

	return nil
}
func (document *Document) Get() documents.Document {
	return document
}
func (document *Document) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(document.GetClassificationID().Bytes(), document.GetImmutables().GenerateHashID().Bytes())
}
func (document *Document) GetProperty(propertyID ids.PropertyID) properties.AnyProperty {
	if property := document.Immutables.GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Mutables.GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}
func (document *Document) GetClassificationID() ids.ClassificationID {
	return document.ClassificationID
}
func (document *Document) GetImmutables() qualified.Immutables {
	return document.Immutables
}
func (document *Document) GetMutables() qualified.Mutables {
	return document.Mutables
}

// TODO write test case
func (document *Document) Mutate(propertyList ...properties.Property) documents.Document {
	document.Mutables = document.Mutables.Mutate(propertyList...).(*base.Mutables)
	return document
}

func NewDocument(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Document {
	return &Document{
		ClassificationID: classificationID.(*baseIDs.ClassificationID),
		Immutables:       immutables.(*base.Immutables),
		Mutables:         mutables.(*base.Mutables),
	}
}
