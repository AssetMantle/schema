// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	stringUtilities "github.com/AssetMantle/schema/go/ids/utilities"
)

var _ ids.PropertyID = (*PropertyID)(nil)

func (propertyID *PropertyID) ValidateBasic() error {
	if err := propertyID.KeyID.ValidateBasic(); err != nil {
		return err
	}
	if err := propertyID.TypeID.ValidateBasic(); err != nil {
		return err
	}
	return nil
}
func (propertyID *PropertyID) GetTypeID() ids.StringID {
	return NewStringID(constants.PropertyIDType)
}
func (propertyID *PropertyID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypePropertyID(), nil
	}

	keyIDAndTypeID := stringUtilities.SplitCompositeIDString(idString)
	if len(keyIDAndTypeID) != 2 {
		return PrototypePropertyID(), errorConstants.IncorrectFormat.Wrapf("expected composite id")
	} else if keyID, err := PrototypeStringID().FromString(keyIDAndTypeID[0]); err != nil {
		return PrototypePropertyID(), err
	} else if typeID, err := PrototypeStringID().FromString(keyIDAndTypeID[1]); err != nil {
		return PrototypePropertyID(), err
	} else {
		propertyID := &PropertyID{KeyID: keyID.(*StringID), TypeID: typeID.(*StringID)}
		if propertyID.ValidateBasic() != nil {
			return PrototypePropertyID(), err
		}

		return propertyID, nil
	}
}
func (propertyID *PropertyID) AsString() string {
	return stringUtilities.JoinIDStrings(propertyID.KeyID.AsString(), propertyID.TypeID.AsString())
}
func (propertyID *PropertyID) IsPropertyID() {}
func (propertyID *PropertyID) GetKey() ids.StringID {
	return propertyID.KeyID
}
func (propertyID *PropertyID) GetDataTypeID() ids.StringID {
	return propertyID.TypeID
}
func (propertyID *PropertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.KeyID.Bytes()...)
	Bytes = append(Bytes, propertyID.TypeID.Bytes()...)

	return Bytes
}
func (propertyID *PropertyID) Compare(id ids.ID) int {
	return bytes.Compare(propertyID.Bytes(), id.ToAnyID().Get().(*PropertyID).Bytes())
}
func (propertyID *PropertyID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_PropertyID{
			PropertyID: propertyID,
		},
	}
}

func PrototypePropertyID() ids.PropertyID {
	return &PropertyID{
		KeyID:  PrototypeStringID().(*StringID),
		TypeID: PrototypeStringID().(*StringID),
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return &PropertyID{
		KeyID:  key.(*StringID),
		TypeID: Type.(*StringID),
	}
}
