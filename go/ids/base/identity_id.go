package base

import (
	"bytes"
	"fmt"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"github.com/AssetMantle/schema/go/qualified"
	"strings"
)

var _ ids.IdentityID = (*IdentityID)(nil)

func (identityID *IdentityID) ValidateBasic() error {
	return identityID.HashID.ValidateBasic()
}
func (identityID *IdentityID) IsIdentityID() {}
func (identityID *IdentityID) GetTypeID() ids.StringID {
	return NewStringID(constants.IdentityIDType)
}
func (identityID *IdentityID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeIdentityID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeIdentityID(), err
	} else {
		identityID := &IdentityID{HashID: hashID.(*HashID)}
		if err := identityID.ValidateBasic(); err != nil {
			return PrototypeIdentityID(), err
		}

		return identityID, nil
	}
}
func (identityID *IdentityID) AsString() string {
	return identityID.HashID.AsString()
}
func (identityID *IdentityID) GetHashID() ids.HashID {
	return identityID.HashID
}
func (identityID *IdentityID) Bytes() []byte {
	return identityID.HashID.Bytes()
}
func (identityID *IdentityID) Compare(id ids.ID) int {
	return bytes.Compare(identityID.Bytes(), id.ToAnyID().Get().(*IdentityID).Bytes())
}
func (identityID *IdentityID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_IdentityID{
			IdentityID: identityID,
		},
	}
}

func identityIDFromInterface(i interface{}) *IdentityID {
	switch value := i.(type) {
	case *IdentityID:
		return value
	default:
		panic(fmt.Errorf("expected *IdentityID, got %T", i))
	}
}

func NewIdentityID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.IdentityID {
	return &IdentityID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeIdentityID() ids.IdentityID {
	return &IdentityID{
		HashID: PrototypeHashID().(*HashID),
	}
}
