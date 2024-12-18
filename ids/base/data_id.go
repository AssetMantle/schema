// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"fmt"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/constants"
	stringUtilities "github.com/AssetMantle/schema/ids/utilities"
	"strings"
)

var _ ids.DataID = (*DataID)(nil)

func (dataID *DataID) ValidateBasic() error {
	if dataID == nil {
		return fmt.Errorf("data ID is empty")
	}

	if err := dataID.TypeID.ValidateBasic(); err != nil {
		return err
	}
	if err := dataID.HashID.ValidateBasic(); err != nil {
		return err
	}
	return nil
}
func (dataID *DataID) GetTypeID() ids.StringID {
	return NewStringID(constants.DataIDType)
}
func (dataID *DataID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeDataID(), nil
	}

	typeIDAndHashID := stringUtilities.SplitCompositeIDString(idString)
	if len(typeIDAndHashID) != 2 {
		return PrototypeDataID(), fmt.Errorf("invalid dataID string %s", idString)
	} else if typeID, err := PrototypeStringID().FromString(typeIDAndHashID[0]); err != nil {
		return PrototypeDataID(), err
	} else if hashID, err := PrototypeHashID().FromString(typeIDAndHashID[1]); err != nil {
		return PrototypeDataID(), err
	} else {
		dataID := &DataID{TypeID: typeID.(*StringID), HashID: hashID.(*HashID)}
		if dataID.ValidateBasic() != nil {
			return PrototypeDataID(), err
		}

		return dataID, nil
	}
}
func (dataID *DataID) AsString() string {
	return stringUtilities.JoinIDStrings(dataID.TypeID.AsString(), dataID.HashID.AsString())
}
func (dataID *DataID) GetHashID() ids.HashID {
	return dataID.HashID
}
func (dataID *DataID) IsDataID() {
}
func (dataID *DataID) DataIDString() string {
	return stringUtilities.JoinIDStrings(dataID.TypeID.AsString(), dataID.HashID.AsString())
}
func (dataID *DataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.TypeID.Bytes()...)
	Bytes = append(Bytes, dataID.HashID.Bytes()...)

	return Bytes
}
func (dataID *DataID) Compare(id ids.ID) int {
	return bytes.Compare(dataID.Bytes(), id.ToAnyID().Get().(*DataID).Bytes())
}
func (dataID *DataID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_DataID{
			DataID: dataID,
		},
	}
}

func GenerateDataID(data data.Data) ids.DataID {
	if data == nil {
		panic(fmt.Errorf("data is nil"))
	}

	return &DataID{
		TypeID: data.GetTypeID().(*StringID),
		HashID: data.GenerateHashID().(*HashID),
	}
}

func PrototypeDataID() ids.DataID {
	return &DataID{
		TypeID: PrototypeStringID().(*StringID),
		HashID: PrototypeHashID().(*HashID),
	}
}
