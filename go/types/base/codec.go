// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	codecUtilities "github.com/AssetMantle/schema/go/utilities"
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Height{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, signature{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, Split{})
}
