package base

import (
	"bytes"
	"strings"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	stringUtilities "github.com/AssetMantle/schema/go/ids/utilities"
)

var _ ids.SplitID = (*SplitID)(nil)

func (splitID *SplitID) ValidateBasic() error {
	if err := splitID.OwnableID.ValidateBasic(); err != nil {
		return err
	}
	if err := splitID.OwnerID.ValidateBasic(); err != nil {
		return err
	}
	return nil
}
func (splitID *SplitID) GetTypeID() ids.StringID {
	return NewStringID(constants.SplitIDType)
}
func (splitID *SplitID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeSplitID(), nil
	}

	ownableIDAndOwnerID := stringUtilities.SplitCompositeIDString(idString)
	if len(ownableIDAndOwnerID) != 2 {
		return PrototypeSplitID(), errorConstants.IncorrectFormat.Wrapf("expected composite id")
	} else if ownableID, err := PrototypeOwnableID().FromString(ownableIDAndOwnerID[0]); err != nil {
		return PrototypeSplitID(), err
	} else if ownerID, err := PrototypeIdentityID().FromString(ownableIDAndOwnerID[1]); err != nil {
		return PrototypeSplitID(), err
	} else {
		splitID := &SplitID{OwnableID: ownableID.(ids.OwnableID).ToAnyOwnableID().(*AnyOwnableID), OwnerID: ownerID.(*IdentityID)}
		if err := splitID.ValidateBasic(); err != nil {
			return PrototypeSplitID(), err
		}

		return splitID, nil
	}
}
func (splitID *SplitID) AsString() string {
	return stringUtilities.JoinIDStrings(splitID.OwnableID.AsString(), splitID.OwnerID.AsString())
}
func (splitID *SplitID) GetOwnableID() ids.OwnableID {
	return splitID.OwnableID
}
func (splitID *SplitID) GetOwnerID() ids.IdentityID {
	return splitID.OwnerID
}
func (splitID *SplitID) IsSplitID() {}
func (splitID *SplitID) Bytes() []byte {
	return append(
		splitID.OwnableID.Bytes(),
		splitID.OwnerID.Bytes()...)
}
func (splitID *SplitID) SplitIDString() string {
	return stringUtilities.JoinIDStrings(splitID.OwnableID.AsString(), splitID.OwnerID.AsString())
}
func (splitID *SplitID) Compare(id ids.ID) int {
	return bytes.Compare(splitID.Bytes(), id.Bytes())
}
func (splitID *SplitID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_SplitID{
			SplitID: splitID,
		},
	}
}

func splitIDFromInterface(i interface{}) *SplitID {
	switch value := i.(type) {
	case *SplitID:
		return value
	default:
		panic(i)
	}
}

func NewSplitID(ownableID ids.OwnableID, ownerID ids.IdentityID) ids.SplitID {
	return &SplitID{
		OwnableID: ownableID.ToAnyOwnableID().(*AnyOwnableID),
		OwnerID:   ownerID.(*IdentityID),
	}
}

func PrototypeSplitID() ids.SplitID {
	return &SplitID{
		OwnableID: PrototypeOwnableID().(*AnyOwnableID),
		OwnerID:   PrototypeIdentityID().(*IdentityID),
	}
}
