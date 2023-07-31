package base

import (
	"bytes"
	"strings"

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
func (*SplitID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeSplitID(), nil
	}

	ownableIDAndOwnerID := stringUtilities.SplitCompositeIDString(idString)
	ownableID, err := PrototypeAnyOwnableID().FromString(ownableIDAndOwnerID[0])
	if err != nil {
		return PrototypeSplitID(), err
	}

	ownerID := PrototypeIdentityID()

	if len(ownableIDAndOwnerID) == 2 {
		OwnerID, err := PrototypeIdentityID().FromString(ownableIDAndOwnerID[1])
		if err != nil {
			return PrototypeSplitID(), err
		}

		ownerID = OwnerID.(*IdentityID)
	}

	splitID := &SplitID{OwnableID: ownableID.(ids.OwnableID).ToAnyOwnableID().(*AnyOwnableID), OwnerID: ownerID.(*IdentityID)}
	if err := splitID.ValidateBasic(); err != nil {
		return PrototypeSplitID(), err
	}

	return splitID, nil
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
	// NOTE: if modified, also modify the FromBytes Method
	return append(
		splitID.OwnableID.Bytes(),
		splitID.OwnerID.Bytes()...)
}

// TODO optimize this
func (*SplitID) MustGetFromPrefixedStoreKeyBytes(prefixBytes, storeKeyBytes []byte) ids.SplitID {
	keyBytes := bytes.TrimPrefix(storeKeyBytes, prefixBytes)
	anyOwnableIDBytes := keyBytes[:len(keyBytes)-32]

	var anyOwnableID *AnyOwnableID
	if len(anyOwnableIDBytes) < 32 {
		anyOwnableID = NewCoinID(NewStringID(string(anyOwnableIDBytes))).ToAnyOwnableID().(*AnyOwnableID)
	} else {
		anyOwnableID = (&AssetID{HashID: &HashID{IDBytes: anyOwnableIDBytes}}).ToAnyOwnableID().(*AnyOwnableID)
	}

	splitID := &SplitID{
		OwnableID: anyOwnableID,
		OwnerID:   &IdentityID{HashID: &HashID{IDBytes: keyBytes[len(keyBytes)-32:]}},
	}

	if bytes.Compare(splitID.Bytes(), keyBytes) != 0 {
		panic("invalid key")
	}

	return splitID
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
		OwnableID: PrototypeAnyOwnableID().(*AnyOwnableID),
		OwnerID:   PrototypeIdentityID().(*IdentityID),
	}
}
