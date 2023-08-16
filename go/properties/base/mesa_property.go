// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties"
)

var _ properties.MesaProperty = (*MesaProperty)(nil)

func (mesaProperty *MesaProperty) ValidateBasic() error {
	if err := mesaProperty.ID.ValidateBasic(); err != nil {
		return err
	}
	if err := mesaProperty.DataID.ValidateBasic(); err != nil {
		return err
	}
	if mesaProperty.DataID.TypeID.Compare(mesaProperty.ID.TypeID) != 0 {
		return errorConstants.IncorrectFormat.Wrapf("property %s typeID %s does not match data type %s", mesaProperty.ID.KeyID.AsString(), mesaProperty.ID.TypeID.AsString(), mesaProperty.DataID.TypeID.AsString())
	}
	return nil
}
func (mesaProperty *MesaProperty) GetID() ids.PropertyID {
	return mesaProperty.ID
}
func (mesaProperty *MesaProperty) GetDataID() ids.DataID {
	return mesaProperty.DataID
}
func (mesaProperty *MesaProperty) GetKey() ids.StringID {
	return mesaProperty.ID.GetKey()
}
func (mesaProperty *MesaProperty) GetDataTypeID() ids.StringID {
	return mesaProperty.ID.GetDataTypeID()
}
func (mesaProperty *MesaProperty) GetBondWeight() sdkTypes.Int {
	if zeroData, err := base.PrototypeAnyData().FromString(mesaProperty.GetDataTypeID().AsString()); err != nil {
		panic(err)
	} else {
		return zeroData.GetBondWeight()
	}
}
func (mesaProperty *MesaProperty) GetHash() ids.HashID {
	return mesaProperty.DataID.GetHashID()
}
func (mesaProperty *MesaProperty) IsMeta() bool {
	return false
}
func (mesaProperty *MesaProperty) IsMesa() bool {
	return true
}
func (mesaProperty *MesaProperty) Compare(property properties.Property) int {
	// NOTE: compare property can be meta or mesa, so listable must only be cast to Property Interface and not MesaProperty
	return mesaProperty.GetID().Compare(property.GetID())
}
func (mesaProperty *MesaProperty) ToAnyProperty() properties.AnyProperty {
	return &AnyProperty{
		Impl: &AnyProperty_MesaProperty{
			MesaProperty: mesaProperty,
		},
	}
}

// Mutate NOTE: Need to check if propertyID type is not different from data after mutating
func (mesaProperty *MesaProperty) Mutate(data data.Data) properties.Property {
	mesaProperty.DataID = data.GetID().(*baseIDs.DataID)
	if err := mesaProperty.ValidateBasic(); err != nil {
		panic(err)
	}
	return mesaProperty
}
func NewEmptyMesaPropertyFromID(propertyID ids.PropertyID) properties.MesaProperty {
	return &MesaProperty{
		ID: propertyID.(*baseIDs.PropertyID),
	}
}
func NewMesaProperty(key ids.StringID, data data.Data) properties.MesaProperty {
	return &MesaProperty{
		ID:     baseIDs.NewPropertyID(key, data.GetTypeID()).(*baseIDs.PropertyID),
		DataID: data.GetID().(*baseIDs.DataID),
	}
}
