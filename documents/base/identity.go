package base

import (
	"fmt"
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified"
	"github.com/AssetMantle/schema/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
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
		return fmt.Errorf("authentication data type %s does not conform to expected type %s for list", identity.GetAuthentication().Get()[0].Get().GetTypeID().AsString(), baseData.PrototypeAccAddressData().GetTypeID().AsString())
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
