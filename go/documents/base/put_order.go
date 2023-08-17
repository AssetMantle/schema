package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	propertyConstants "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/AssetMantle/schema/go/types"
	"github.com/AssetMantle/schema/go/types/base"
)

var _ documents.PutOrder = (*order)(nil)

var putOrderClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(propertyConstants.MakerIDProperty, propertyConstants.MakerAssetIDProperty, propertyConstants.TakerAssetIDProperty, propertyConstants.MakerSplitProperty, propertyConstants.TakerSplitProperty, propertyConstants.ExpiryHeightProperty)), baseQualified.NewMutables(baseLists.NewPropertyList()))

func (order order) GetPutOrderID() ids.OrderID {
	return baseIDs.NewOrderID(putOrderClassificationID, order.Document.GetImmutables())
}

func PrototypePutOrder() documents.PutOrder {
	// TODO derive values from zero values of data
	return NewPutOrder(baseIDs.PrototypeIdentityID(), baseIDs.PrototypeAssetID(), baseIDs.PrototypeAssetID(), sdkTypes.ZeroInt(), sdkTypes.ZeroInt(), base.NewHeight(-1))
}

func NewPutOrder(makerID ids.IdentityID, makerAssetID ids.AssetID, takerAssetID ids.AssetID, makerSplit sdkTypes.Int, takerSplit sdkTypes.Int, expiryHeight types.Height) documents.PutOrder {
	return NewOrder(putOrderClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(makerID)),
		baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(makerAssetID)),
		baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(takerAssetID)),
		baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(makerSplit)),
		baseProperties.NewMetaProperty(propertyConstants.TakerSplitProperty.GetKey(), baseData.NewNumberData(takerSplit)),
		baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(expiryHeight)),
	)), baseQualified.NewMutables(baseLists.NewPropertyList())).(order)
}
