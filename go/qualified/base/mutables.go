// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/qualified"
)

var _ qualified.Mutables = (*Mutables)(nil)

func (mutables *Mutables) ValidateBasic() error {
	return mutables.PropertyList.ValidateBasic()
}
func (mutables *Mutables) GetMutablePropertyList() lists.PropertyList {
	if mutables.PropertyList == nil {
		return baseLists.NewPropertyList()
	}

	return mutables.PropertyList
}
func (mutables *Mutables) GetProperty(id ids.PropertyID) properties.AnyProperty {
	return mutables.GetMutablePropertyList().GetProperty(id)
}
func (mutables *Mutables) Mutate(propertyList ...properties.Property) qualified.Mutables {
	for _, property := range propertyList {
		mutables.PropertyList = mutables.PropertyList.Mutate(property).(*baseLists.PropertyList)
	}

	return mutables
}

func NewMutables(propertyList lists.PropertyList) qualified.Mutables {
	if propertyList == nil {
		return &Mutables{baseLists.NewPropertyList().(*baseLists.PropertyList)}
	}
	return &Mutables{
		PropertyList: propertyList.(*baseLists.PropertyList),
	}
}
