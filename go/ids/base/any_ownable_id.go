package base

import (
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
)

type ownableIDGetter interface {
	get() ids.OwnableID
}

var _ ownableIDGetter = (*AnyOwnableID_AssetID)(nil)

func (m *AnyOwnableID_AssetID) get() ids.OwnableID {
	return m.AssetID
}

var _ ownableIDGetter = (*AnyOwnableID_CoinID)(nil)

func (m *AnyOwnableID_CoinID) get() ids.OwnableID {
	return m.CoinID
}

var _ ids.AnyOwnableID = (*AnyOwnableID)(nil)

func (m *AnyOwnableID) Get() ids.OwnableID {
	return m.Impl.(ownableIDGetter).get()
}
func (m *AnyOwnableID) Compare(id ids.ID) int {
	return m.Impl.(ownableIDGetter).get().Compare(id)
}
func (m *AnyOwnableID) GetTypeID() ids.StringID {
	return m.Impl.(ownableIDGetter).get().GetTypeID()
}
func (m *AnyOwnableID) FromString(idString string) (ids.ID, error) {
	idTypeString, valueString := splitIDTypeAndValueStrings(idString)
	if idTypeString != "" {
		var id ids.ID
		var err error

		switch NewStringID(idTypeString).AsString() {
		case PrototypeAssetID().GetTypeID().AsString():
			id, err = PrototypeAssetID().FromString(valueString)
		case PrototypeCoinID().GetTypeID().AsString():
			id, err = PrototypeCoinID().FromString(valueString)
		default:
			id, err = nil, errorConstants.IncorrectFormat.Wrapf("type identifier is not recognised")
		}

		if err != nil {
			return nil, err
		}

		return id, nil
	}

	return nil, errorConstants.IncorrectFormat.Wrapf("type identifier is missing")
}
func (m *AnyOwnableID) AsString() string {
	return m.Impl.(ownableIDGetter).get().AsString()
}
func (m *AnyOwnableID) Bytes() []byte {
	return m.Impl.(ownableIDGetter).get().Bytes()
}
func (m *AnyOwnableID) ToAnyID() ids.AnyID {
	return m.Impl.(ownableIDGetter).get().ToAnyID()
}
func (m *AnyOwnableID) ToAnyOwnableID() ids.AnyOwnableID {
	return m.Impl.(ownableIDGetter).get().ToAnyOwnableID()
}
func (m *AnyOwnableID) IsAnyOwnableID() {
}

func (m *AnyOwnableID) ValidateBasic() error {
	return m.Impl.(ownableIDGetter).get().ValidateBasic()
}

func PrototypeOwnableID() ids.AnyOwnableID {
	return PrototypeAssetID().ToAnyOwnableID()
}
