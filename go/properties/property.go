// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/traits"
)

type Property interface {
	GetID() ids.PropertyID
	GetDataID() ids.DataID
	GetKey() ids.StringID
	GetDataTypeID() ids.StringID
	GetBondWeight() int64

	IsMeta() bool
	ValidateBasic() error

	ToAnyProperty() AnyProperty

	Mutate(data.Data) Property
	traits.Listable
}
