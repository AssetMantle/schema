// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	"github.com/AssetMantle/schema/qualified"
)

var _ qualified.Immutables = (*Immutables)(nil)

func (immutables *Immutables) ValidateBasic() error {
	return immutables.PropertyList.ValidateBasic()
}

// TODO write test case
func (immutables *Immutables) GetImmutablePropertyList() lists.PropertyList {
	if immutables.PropertyList.Get() == nil {
		return baseLists.NewPropertyList()
	}

	return immutables.PropertyList
}
func (immutables *Immutables) GetProperty(id ids.PropertyID) properties.AnyProperty {
	return immutables.GetImmutablePropertyList().GetProperty(id)
}
func (immutables *Immutables) GenerateHashID() ids.ID {
	metaList := make([][]byte, len(immutables.PropertyList.Get()))

	for i, immutableProperty := range immutables.PropertyList.Get() {
		metaList[i] = immutableProperty.GetDataID().(ids.DataID).GetHashID().Bytes()
	}

	return baseIDs.GenerateHashID(metaList...)
}

func NewImmutables(propertyList lists.PropertyList) qualified.Immutables {
	if propertyList == nil {
		return &Immutables{baseLists.NewPropertyList().(*baseLists.PropertyList)}
	}
	return &Immutables{
		PropertyList: propertyList.(*baseLists.PropertyList),
	}
}
