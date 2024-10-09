package base

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/documents"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/properties"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified"
)

type classification struct {
	documents.Document
}

var _ documents.Classification = (*classification)(nil)

func (classification classification) ValidateBasic() error {
	if err := classification.Document.ValidateBasic(); err != nil {
		return err
	}

	if property := classification.GetProperty(constantProperties.BondAmountProperty.GetID()); property != nil && !property.IsMeta() {
		return fmt.Errorf("classification must have a revealed %s", constantProperties.BondAmountProperty.GetID())
	}

	return nil
}
func (classification classification) GetBondAmount() math.Int {
	if property := classification.Document.GetProperty(constantProperties.BondAmountProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	return constantProperties.BondAmountProperty.GetData().Get().(data.NumberData).Get()
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
