package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/schema/utilities/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyProperty{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, MetaProperty{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, MesaProperty{})
}
