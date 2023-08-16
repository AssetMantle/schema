// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"github.com/AssetMantle/schema/go/qualified"
)

var _ ids.AssetID = (*AssetID)(nil)

func (assetID *AssetID) ValidateBasic() error {
	return assetID.HashID.ValidateBasic()
}
func (assetID *AssetID) GetTypeID() ids.StringID {
	return NewStringID(constants.AssetIDType)
}
func (assetID *AssetID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeAssetID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeAssetID(), err
	} else {
		assetID := &AssetID{HashID: hashID.(*HashID)}
		if assetID.ValidateBasic() != nil {
			return PrototypeAssetID(), err
		}

		return assetID, nil
	}
}
func (assetID *AssetID) AsString() string {
	return assetID.HashID.AsString()
}
func (assetID *AssetID) Bytes() []byte {
	return assetID.HashID.IDBytes
}
func (assetID *AssetID) IsAssetID() {}
func (assetID *AssetID) Compare(id ids.ID) int {
	return bytes.Compare(assetID.Bytes(), id.ToAnyID().Get().(*AssetID).Bytes())
}
func (assetID *AssetID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_AssetID{
			AssetID: assetID,
		},
	}
}

func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return &AssetID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeAssetID() ids.AssetID {
	return &AssetID{
		HashID: PrototypeHashID().(*HashID),
	}
}
