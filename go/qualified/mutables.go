// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

/*
 *  Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
 *  SPDX-License-Identifier: Apache-2.0
 */

package qualified

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/properties"
)

type Mutables interface {
	// GetMutablePropertyList return the mutable properties object
	// does not return nil
	GetMutablePropertyList() lists.PropertyList

	GetProperty(ids.PropertyID) properties.AnyProperty

	Mutate(...properties.Property) Mutables

	ValidateBasic() error
}
