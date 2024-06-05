// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/ids"
)

type Property interface {
	GetID() ids.PropertyID
	GetDataID() ids.DataID
	GetKey() ids.StringID
	GetDataTypeID() ids.StringID
	GetBondWeight() sdkTypes.Int

	IsMeta() bool
	ValidateBasic() error

	ToAnyProperty() AnyProperty

	Mutate(data.Data) (Property, error)
	Compare(Property) int
}
