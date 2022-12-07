// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.IDData = (*IDData)(nil)

func (idData *IDData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(idData)
}
func (idData *IDData) Compare(listable traits.Listable) int {
	compareIDData, err := idDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Bytes(), compareIDData.Bytes())
}
func (idData *IDData) Bytes() []byte {
	return idData.Bytes()
}
func (idData *IDData) GetType() ids.StringID {
	return dataConstants.IDDataID
}
func (idData *IDData) ZeroValue() data.Data {
	return IDDataPrototype()
}
func (idData *IDData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(idData.Bytes())
}
func (idData *IDData) Get() ids.ID {
	return idData.Value
}

func idDataFromInterface(listable traits.Listable) (data.IDData, error) {
	switch value := listable.(type) {
	case *IDData:
		return value, nil
	default:
		return nil, constants.MetaDataError
	}
}

func IDDataPrototype() data.IDData {
	return NewIDData(baseIDs.NewStringID(""))
}

func NewIDData(value ids.ID) data.IDData {
	return &IDData{
		Value: value,
	}
}
