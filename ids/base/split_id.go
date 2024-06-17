package base

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/constants"
	stringUtilities "github.com/AssetMantle/schema/ids/utilities"
)

var _ ids.SplitID = (*SplitID)(nil)

func (splitID *SplitID) ValidateBasic() error {
	if splitID == nil {
		return fmt.Errorf("split ID is empty")
	}

	if err := splitID.AssetID.ValidateBasic(); err != nil {
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

	assetIDAndOwnerID := stringUtilities.SplitCompositeIDString(idString)
	assetID, err := PrototypeAssetID().FromString(assetIDAndOwnerID[0])
	if err != nil {
		return PrototypeSplitID(), err
	}

	ownerID := PrototypeIdentityID()

	if len(assetIDAndOwnerID) == 2 {
		OwnerID, err := PrototypeIdentityID().FromString(assetIDAndOwnerID[1])
		if err != nil {
			return PrototypeSplitID(), err
		}

		ownerID = OwnerID.(*IdentityID)
	}

	splitID := &SplitID{AssetID: assetID.(*AssetID), OwnerID: ownerID.(*IdentityID)}
	if err := splitID.ValidateBasic(); err != nil {
		return PrototypeSplitID(), err
	}

	return splitID, nil
}
func (splitID *SplitID) AsString() string {
	return stringUtilities.JoinIDStrings(splitID.AssetID.AsString(), splitID.OwnerID.AsString())
}
func (splitID *SplitID) GetAssetID() ids.AssetID {
	return splitID.AssetID
}
func (splitID *SplitID) GetOwnerID() ids.IdentityID {
	return splitID.OwnerID
}
func (splitID *SplitID) IsSplitID() {}
func (splitID *SplitID) Bytes() []byte {
	// NOTE: if modified, also modify the MustGetFromPrefixedStoreKeyBytes Method
	return append(
		splitID.AssetID.Bytes(),
		splitID.OwnerID.Bytes()...)
}

// TODO optimize this
func (*SplitID) MustGetFromPrefixedStoreKeyBytes(prefixBytes, storeKeyBytes []byte) ids.SplitID {
	keyBytes := bytes.TrimPrefix(storeKeyBytes, prefixBytes)

	splitID := &SplitID{
		AssetID: &AssetID{HashID: &HashID{IDBytes: keyBytes[:len(keyBytes)-32]}},
		OwnerID: &IdentityID{HashID: &HashID{IDBytes: keyBytes[len(keyBytes)-32:]}},
	}

	if bytes.Compare(splitID.Bytes(), keyBytes) != 0 {
		panic("invalid key")
	}

	return splitID
}
func (splitID *SplitID) SplitIDString() string {
	return stringUtilities.JoinIDStrings(splitID.AssetID.AsString(), splitID.OwnerID.AsString())
}
func (splitID *SplitID) Compare(id ids.ID) int {
	return bytes.Compare(splitID.Bytes(), id.ToAnyID().Get().(*SplitID).Bytes())
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

func NewSplitID(assetID ids.AssetID, ownerID ids.IdentityID) ids.SplitID {
	return &SplitID{
		AssetID: assetID.(*AssetID),
		OwnerID: ownerID.(*IdentityID),
	}
}

func PrototypeSplitID() ids.SplitID {
	return &SplitID{
		AssetID: PrototypeAssetID().(*AssetID),
		OwnerID: PrototypeIdentityID().(*IdentityID),
	}
}
