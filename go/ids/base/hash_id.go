package base

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strings"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
)

// type hashID struct {
//	HashBytes []byte
// }

var _ ids.HashID = (*HashID)(nil)

func (hashID *HashID) IsHashID() {}
func (hashID *HashID) ValidateBasic() error {
	if len(hashID.IDBytes) != 0 && len(hashID.IDBytes) != 32 && hashID.IDBytes != nil {
		return errorConstants.IncorrectFormat.Wrapf("invalid hashID length")
	}
	return nil
}
func (hashID *HashID) GetTypeID() ids.StringID {
	return NewStringID(constants.HashIDType)
}
func (hashID *HashID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeHashID(), nil
	}

	if hashBytes, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(strings.ReplaceAll(idString, "-", "+"), "_", "/")); err != nil {
		return PrototypeHashID(), err
	} else {
		hashID := &HashID{IDBytes: hashBytes}
		if err := hashID.ValidateBasic(); err != nil {
			return PrototypeHashID(), err
		}

		return hashID, nil
	}
}

// TODO test if nil and empty result in ""
func (hashID *HashID) AsString() string {
	return base64.StdEncoding.EncodeToString(hashID.IDBytes)
}
func (hashID *HashID) Bytes() []byte {
	return hashID.IDBytes
}
func (hashID *HashID) Compare(id ids.ID) int {
	return bytes.Compare(hashID.Bytes(), id.ToAnyID().Get().(*HashID).Bytes())
}
func (hashID *HashID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_HashID{
			HashID: hashID,
		},
	}
}

func hashIDFromInterface(i interface{}) *HashID {
	switch value := i.(type) {
	case *HashID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("expected *HashID, got %T", i))
	}
}

// TODO test
func GenerateHashID(toHashList ...[]byte) ids.HashID {
	var nonEmptyByteList [][]byte

	for _, value := range toHashList {
		if len(value) != 0 {
			nonEmptyByteList = append(nonEmptyByteList, value)
		}
	}

	if len(nonEmptyByteList) == 0 {
		return &HashID{IDBytes: []byte{}}
	}

	sort.Slice(nonEmptyByteList, func(i, j int) bool { return bytes.Compare(nonEmptyByteList[i], nonEmptyByteList[j]) == -1 })

	hash := sha256.New()

	// TODO check if nil elements in slice
	if _, err := hash.Write(bytes.Join(nonEmptyByteList, nil)); err != nil {
		panic(err)
	}

	return &HashID{IDBytes: hash.Sum(nil)}
}

func PrototypeHashID() ids.HashID {
	return GenerateHashID()
}
