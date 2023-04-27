package base

import (
	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/documents"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/qualified"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type classification struct {
	documents.Document
}

var _ documents.Classification = (*classification)(nil)

func (classification classification) GetBondAmount() sdkTypes.Int {
	if property := classification.Document.GetProperty(constants.BondAmountProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	return constants.BondAmountProperty.GetData().Get().(data.NumberData).Get()
}
func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) documents.Classification {
	return classification{
		Document: NewDocument(baseIDs.NewClassificationID(immutables, mutables), immutables, mutables),
	}
}

func NewClassificationFromDocument(document documents.Document) documents.Classification {
	return classification{
		Document: document,
	}
}
