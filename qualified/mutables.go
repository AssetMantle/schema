// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package qualified

import (
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/properties"
)

type Mutables interface {
	// GetMutablePropertyList return the mutable properties object
	// does not return nil
	GetMutablePropertyList() lists.PropertyList

	GetProperty(ids.PropertyID) properties.AnyProperty

	Mutate(...properties.Property) Mutables

	ValidateBasic() error
}
