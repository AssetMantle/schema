// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/properties"
	"github.com/AssetMantle/schema/qualified"
)

type Document interface {
	Get() Document
	GenerateHashID() ids.HashID
	GetClassificationID() ids.ClassificationID
	// GetProperty returns property from a document searching in both Mutables and Immutables
	// * Returns nil if property is not found
	GetProperty(ids.PropertyID) properties.AnyProperty
	GetImmutables() qualified.Immutables
	GetMutables() qualified.Mutables

	ValidateBasic() error
	Mutate(...properties.Property) Document
}
