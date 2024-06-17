// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package qualified

import (
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/properties"
)

type Immutables interface {
	// GetImmutablePropertyList return the immutable properties object
	// does not return nil
	GetImmutablePropertyList() lists.PropertyList

	GetProperty(id ids.PropertyID) properties.AnyProperty

	GenerateHashID() ids.ID

	ValidateBasic() error
}
