package base

import (
	"bytes"
	"fmt"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"github.com/AssetMantle/schema/go/qualified"
	"strings"
)

var _ ids.MaintainerID = (*MaintainerID)(nil)

func (maintainerID *MaintainerID) ValidateBasic() error {
	if maintainerID == nil {
		return fmt.Errorf("maintainer ID is empty")
	}

	return maintainerID.HashID.ValidateBasic()
}
func (maintainerID *MaintainerID) GetTypeID() ids.StringID {
	return NewStringID(constants.MaintainerIDType)
}
func (maintainerID *MaintainerID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeMaintainerID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeMaintainerID(), err
	} else {
		maintainerID := &MaintainerID{HashID: hashID.(*HashID)}
		if err := maintainerID.ValidateBasic(); err != nil {
			return PrototypeMaintainerID(), err
		}

		return maintainerID, nil
	}
}
func (maintainerID *MaintainerID) AsString() string {
	return maintainerID.HashID.AsString()
}
func (maintainerID *MaintainerID) Bytes() []byte {
	return maintainerID.HashID.IDBytes
}
func (maintainerID *MaintainerID) IsMaintainerID() {}
func (maintainerID *MaintainerID) Compare(id ids.ID) int {
	return bytes.Compare(maintainerID.Bytes(), id.ToAnyID().Get().(*MaintainerID).Bytes())
}
func (maintainerID *MaintainerID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_MaintainerID{
			MaintainerID: maintainerID,
		},
	}
}

func NewMaintainerID(maintainerClassificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return &MaintainerID{
		HashID: GenerateHashID(maintainerClassificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return &MaintainerID{
		HashID: PrototypeHashID().(*HashID),
	}
}
