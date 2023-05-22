// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, IDList{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, ParameterList{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, PropertyList{})
}
