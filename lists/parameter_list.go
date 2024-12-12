// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/parameters"
	"github.com/cosmos/cosmos-sdk/codec"
)

type ParameterList interface {
	ValidateBasic() error
	Get() []parameters.Parameter
	GetParameter(ids.PropertyID) parameters.Parameter
	Mutate(...parameters.Parameter) ParameterList

	codec.ProtoMarshaler
}
