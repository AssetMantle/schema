// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"github.com/AssetMantle/schema/go/ids/utilities"
)

var _ ids.StringID = (*StringID)(nil)

func (stringID *StringID) Get() string {
	return stringID.IDString
}
func (stringID *StringID) ValidateBasic() error {
	if !utilities.IsValidStringID(stringID.AsString()) {
		return errorConstants.IncorrectFormat.Wrapf("regular expression check failed")
	}
	return nil
}
func (stringID *StringID) IsStringID() {}
func (stringID *StringID) GetTypeID() ids.StringID {
	return NewStringID(constants.StringIDType)
}
func (stringID *StringID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeStringID(), nil
	}

	stringID.IDString = idString
	if err := stringID.ValidateBasic(); err != nil {
		return PrototypeStringID(), err
	}

	return stringID, nil
}
func (stringID *StringID) AsString() string {
	return stringID.IDString
}
func (stringID *StringID) Bytes() []byte {
	return []byte(stringID.IDString)
}
func (stringID *StringID) Compare(id ids.ID) int {
	return strings.Compare(stringID.AsString(), id.ToAnyID().Get().(*StringID).AsString())
}
func (stringID *StringID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_StringID{
			StringID: stringID,
		},
	}
}

func stringIDFromInterface(i interface{}) *StringID {
	switch value := i.(type) {
	case *StringID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("expected *StringID, got %T", i))
	}
}

func StringIDsToIDs(stringIDs []ids.StringID) []ids.ID {
	IDs := make([]ids.ID, len(stringIDs))

	for i, stringID := range stringIDs {
		IDs[i] = stringID
	}

	return IDs
}

func NewStringID(idString string) ids.StringID {
	return &StringID{IDString: idString}
}

func PrototypeStringID() ids.StringID {
	return &StringID{
		IDString: "",
	}
}
