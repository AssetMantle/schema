package base

import (
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/qualified"
	"github.com/AssetMantle/schema/go/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type identity struct {
	documents.Document
}

var _ documents.Identity = (*identity)(nil)

func (identity identity) GetExpiry() types.Height {
	if property := identity.Document.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}

	return constants.ExpiryHeightProperty.GetData().Get().(data.HeightData).Get()
}
func (identity identity) GetAuthentication() data.ListData {
	var dataList []data.ListableData

	if property := identity.Document.GetProperty(constants.AuthenticationProperty.GetID()); property != nil && property.IsMeta() {
		for _, anyData := range property.Get().(properties.MetaProperty).GetData().Get().(data.ListData).Get() {
			dataList = append(dataList, anyData)
		}
	} else {
		for _, anyData := range constants.AuthenticationProperty.GetData().Get().(data.ListData).Get() {
			dataList = append(dataList, anyData)
		}
	}
	return baseData.NewListData(dataList...)
}
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	_, isProvisioned := identity.GetAuthentication().Search(baseData.NewAccAddressData(accAddress))
	return isProvisioned
}
func (identity identity) GetProvisionedAddressCount() sdkTypes.Int {
	return sdkTypes.NewInt(int64(len(identity.GetAuthentication().Get())))
}
func (identity identity) ProvisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	var accAddressList []data.ListableData
	for _, address := range accAddressesToListableData(accAddresses...) {
		accAddressList = append(accAddressList, address)
	}
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), identity.GetAuthentication().Add(accAddressList...)))
	return identity
}
func (identity identity) UnprovisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	var accAddressList []data.ListableData
	for _, address := range accAddressesToListableData(accAddresses...) {
		accAddressList = append(accAddressList, address)
	}
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), identity.GetAuthentication().Remove(accAddressList...)))
	return identity
}
func accAddressesToListableData(accAddresses ...sdkTypes.AccAddress) []data.ListableData {
	accAddressData := make([]data.ListableData, len(accAddresses))
	for i, accAddress := range accAddresses {
		accAddressData[i] = baseData.NewAccAddressData(accAddress)
	}
	return accAddressData
}

func NewIdentity(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Identity {
	return NewIdentityFromDocument(NewDocument(classificationID, immutables, mutables))
}

func NewIdentityFromDocument(document documents.Document) documents.Identity {
	return identity{Document: document}
}
