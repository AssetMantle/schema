package base

import (
	"github.com/AssetMantle/schema/schema/data"
	"github.com/AssetMantle/schema/schema/documents"
	baseIDs "github.com/AssetMantle/schema/schema/ids/base"
	"github.com/AssetMantle/schema/schema/properties"
	"github.com/AssetMantle/schema/schema/properties/constants"
	"github.com/AssetMantle/schema/schema/qualified"
)

type classification struct {
	documents.Document
}

var _ documents.Classification = (*classification)(nil)

func (classification classification) GetBondAmount() int64 {
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
