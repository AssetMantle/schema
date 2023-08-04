package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/qualified"
	"github.com/AssetMantle/schema/go/types"
)

type order struct {
	documents.Document
}

var _ documents.Order = (*order)(nil)

func (order order) GetMakerID() ids.IdentityID {
	if property := order.GetProperty(constants.MakerIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
	}
	return constants.MakerIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
}
func (order order) GetMakerOwnableID() ids.OwnableID {
	if property := order.GetProperty(constants.MakerOwnableIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.OwnableID)
	}
	return constants.MakerOwnableIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.OwnableID)
}
func (order order) GetTakerOwnableID() ids.OwnableID {
	if property := order.GetProperty(constants.TakerOwnableIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.OwnableID)
	}
	return constants.TakerOwnableIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.OwnableID)
}
func (order order) GetMakerOwnableSplit() sdkTypes.Int {
	if property := order.GetProperty(constants.MakerSplitProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}
	return constants.MakerSplitProperty.GetData().Get().(data.NumberData).Get()
}
func (order order) GetTakerOwnableSplit() sdkTypes.Int {
	if property := order.GetProperty(constants.TakerSplitProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}
	return constants.TakerSplitProperty.GetData().Get().(data.NumberData).Get()
}
func (order order) GetExpiryHeight() types.Height {
	if property := order.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}
	return constants.ExpiryHeightProperty.GetData().Get().(data.HeightData).Get()
}
func (order order) GetTakerID() ids.IdentityID {
	if property := order.GetProperty(constants.TakerIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
	}
	return constants.TakerIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
}
func (order order) GetExchangeRate() sdkTypes.Dec {
	if property := order.GetProperty(constants.ExchangeRateProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.DecData).Get()
	}
	return constants.ExchangeRateProperty.GetData().Get().(data.DecData).Get()
}
func (order order) GetExecutionHeight() types.Height {
	if property := order.GetProperty(constants.ExecutionHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}
	return constants.ExecutionHeightProperty.GetData().Get().(data.HeightData).Get()
}

func NewOrder(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Order {
	return order{Document: NewDocument(classificationID, immutables, mutables)}
}

func NewOrderFromDocument(document documents.Document) documents.Order {
	return order{Document: document}
}
