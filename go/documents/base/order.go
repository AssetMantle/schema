package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/documents"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
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

func (order order) ValidateBasic() error {
	if err := order.Document.ValidateBasic(); err != nil {
		return err
	}

	if property := order.GetProperty(constants.MakerIDProperty.GetID()); property == nil || !property.IsMeta() {
		return errorConstants.IncorrectFormat.Wrapf("order must have a revealed %s", constants.MakerIDProperty.GetID())
	}

	if property := order.GetProperty(constants.MakerAssetIDProperty.GetID()); property == nil || !property.IsMeta() {
		return errorConstants.IncorrectFormat.Wrapf("order must have a revealed %s", constants.MakerAssetIDProperty.GetID())
	}

	if property := order.GetProperty(constants.MakerSplitProperty.GetID()); property == nil || !property.IsMeta() {
		return errorConstants.IncorrectFormat.Wrapf("order must have a revealed %s", constants.MakerSplitProperty.GetID())
	}

	if property := order.GetProperty(constants.ExpiryHeightProperty.GetID()); property == nil || !property.IsMeta() {
		return errorConstants.IncorrectFormat.Wrapf("order must have a revealed %s", constants.ExpiryHeightProperty.GetID())
	}

	if property := order.GetProperty(constants.TakerAssetIDProperty.GetID()); property == nil {
		return errorConstants.IncorrectFormat.Wrapf("order must have a %s", constants.TakerAssetIDProperty.GetID())
	}

	if property := order.GetProperty(constants.TakerSplitProperty.GetID()); property == nil {
		return errorConstants.IncorrectFormat.Wrapf("order must have a %s", constants.TakerSplitProperty.GetID())
	}

	return nil
}
func (order order) GetMakerID() ids.IdentityID {
	if property := order.GetProperty(constants.MakerIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
	}
	return constants.MakerIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.IdentityID)
}
func (order order) GetMakerAssetID() ids.AssetID {
	if property := order.GetProperty(constants.MakerAssetIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.AssetID)
	}
	return constants.MakerAssetIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.AssetID)
}
func (order order) GetTakerAssetID() ids.AssetID {
	if property := order.GetProperty(constants.TakerAssetIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().Get().(ids.AssetID)
	}
	return constants.TakerAssetIDProperty.GetData().Get().(data.IDData).Get().Get().(ids.AssetID)
}
func (order order) GetMakerSplit() sdkTypes.Int {
	if property := order.GetProperty(constants.MakerSplitProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}
	return constants.MakerSplitProperty.GetData().Get().(data.NumberData).Get()
}
func (order order) GetTakerSplit() sdkTypes.Int {
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
	// TODO change definition to maker/taker exchange rate
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
