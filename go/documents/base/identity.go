package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/qualified"
	"github.com/AssetMantle/schema/go/types"
)

type identity struct {
	documents.Document
}

var _ documents.Identity = (*identity)(nil)

func (identity identity) ValidateBasic() error {
	if err := identity.Document.ValidateBasic(); err != nil {
		return err
	}

	if len(identity.GetAuthentication().Get()) > 0 && identity.GetAuthentication().Get()[0].Get().GetTypeID().Compare(baseData.PrototypeAccAddressData().GetTypeID()) != 0 {
		return errorConstants.MetaDataError.Wrapf("authentication property must contain a list of addresses only")
	}

	return nil
}
func (identity identity) GetExpiry() types.Height {
	if property := identity.Document.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}

	return constants.ExpiryHeightProperty.GetData().Get().(data.HeightData).Get()
}
func (identity identity) GetAuthentication() data.ListData {
	if property := identity.GetProperty(constants.AuthenticationProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.ListData)
	}

	return constants.AuthenticationProperty.GetData().Get().(data.ListData)
}
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	_, isProvisioned := identity.GetAuthentication().Search(baseData.NewAccAddressData(accAddress))
	return isProvisioned
}
func (identity identity) GetProvisionedAddressCount() sdkTypes.Int {
	return sdkTypes.NewInt(int64(len(identity.GetAuthentication().Get())))
}
func (identity identity) ProvisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	for _, accAddress := range accAddresses {
		identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), identity.GetAuthentication().Add(baseData.NewAccAddressData(accAddress))))
	}

	return identity
}
func (identity identity) UnprovisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	for _, accAddress := range accAddresses {
		identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), identity.GetAuthentication().Remove(baseData.NewAccAddressData(accAddress))))
	}

	return identity
}

func NewIdentity(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Identity {
	return NewIdentityFromDocument(NewDocument(classificationID, immutables, mutables))
}

func NewIdentityFromDocument(document documents.Document) documents.Identity {
	return identity{Document: document}
}
