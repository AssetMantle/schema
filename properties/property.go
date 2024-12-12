// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/ids"
)

type Property interface {
	GetID() ids.PropertyID
	GetDataID() ids.DataID
	GetKey() ids.StringID
	GetDataTypeID() ids.StringID
	GetBondWeight() math.Int

	IsMeta() bool
	ValidateBasic() error

	ToAnyProperty() AnyProperty

	Mutate(data.Data) (Property, error)
	Compare(Property) int
}
