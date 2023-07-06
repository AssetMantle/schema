package base

import (
	"bytes"
	"strings"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"github.com/AssetMantle/schema/go/qualified"
)

var _ ids.ClassificationID = (*ClassificationID)(nil)

func (classificationID *ClassificationID) ValidateBasic() error {
	return classificationID.HashID.ValidateBasic()
}
func (classificationID *ClassificationID) GetTypeID() ids.StringID {
	return NewStringID(constants.ClassificationIDType)
}
func (classificationID *ClassificationID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeClassificationID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeClassificationID(), err
	} else {
		classificationID := &ClassificationID{HashID: hashID.(*HashID)}
		if err := classificationID.ValidateBasic(); err != nil {
			return PrototypeClassificationID(), err
		}

		return classificationID, nil
	}
}
func (classificationID *ClassificationID) AsString() string {
	return classificationID.HashID.AsString()
}
func (classificationID *ClassificationID) Bytes() []byte {
	return classificationID.HashID.IDBytes
}
func (classificationID *ClassificationID) IsClassificationID() {}
func (classificationID *ClassificationID) Compare(id ids.ID) int {
	return bytes.Compare(classificationID.Bytes(), id.Bytes())
}
func (classificationID *ClassificationID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_ClassificationID{
			ClassificationID: classificationID,
		},
	}
}

func classificationIDFromInterface(i interface{}) *ClassificationID {
	switch value := i.(type) {
	case *ClassificationID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("Expected ClassificationID, got %T", i))
	}
}

func NewClassificationID(immutables qualified.Immutables, mutables qualified.Mutables) ids.ClassificationID {
	immutableIDByteList := make([][]byte, len(immutables.GetImmutablePropertyList().Get()))
	for i, property := range immutables.GetImmutablePropertyList().Get() {
		immutableIDByteList[i] = property.GetID().Bytes()
	}

	mutableIDByteList := make([][]byte, len(mutables.GetMutablePropertyList().Get()))
	for i, property := range mutables.GetMutablePropertyList().Get() {
		mutableIDByteList[i] = property.GetID().Bytes()
	}

	return &ClassificationID{HashID: GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), immutables.GenerateHashID().Bytes()).(*HashID)}
}

func PrototypeClassificationID() ids.ClassificationID {
	return &ClassificationID{
		HashID: PrototypeHashID().(*HashID),
	}
}
