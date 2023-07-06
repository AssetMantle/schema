package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
)

var _ ids.CoinID = (*CoinID)(nil)

func (coinID *CoinID) ValidateBasic() error {
	return coinID.StringID.ValidateBasic()
}
func (coinID *CoinID) IsCoinID() {
}
func (coinID *CoinID) GetTypeID() ids.StringID {
	return NewStringID(constants.CoinIDType)
}
func (coinID *CoinID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeCoinID(), nil
	}

	if stringID, err := PrototypeStringID().FromString(idString); err != nil {
		return PrototypeCoinID(), err
	} else {
		coinID := &CoinID{StringID: stringID.(*StringID)}
		if coinID.ValidateBasic() != nil {
			return PrototypeCoinID(), err
		}

		return coinID, nil
	}
}
func (coinID *CoinID) AsString() string {
	return coinID.StringID.AsString()
}

// TODO: Verify
func (coinID *CoinID) Bytes() []byte {
	return []byte(coinID.StringID.IDString)
}
func (coinID *CoinID) IsOwnableID() {}
func (coinID *CoinID) Compare(id ids.ID) int {
	return bytes.Compare(coinID.Bytes(), id.Bytes())
}
func (coinID *CoinID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_OwnableID{
			OwnableID: coinID.ToAnyOwnableID().(*AnyOwnableID),
		},
	}
}
func (coinID *CoinID) ToAnyOwnableID() ids.AnyOwnableID {
	return &AnyOwnableID{
		Impl: &AnyOwnableID_CoinID{
			CoinID: coinID,
		},
	}
}

func NewCoinID(stringID ids.StringID) ids.CoinID {
	return &CoinID{
		StringID: stringID.(*StringID),
	}
}

func PrototypeCoinID() ids.OwnableID {
	return &CoinID{
		StringID: PrototypeStringID().(*StringID),
	}
}
