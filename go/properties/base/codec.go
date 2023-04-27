package base

import (
	codecUtilities "github.com/AssetMantle/schema/go/utilities"
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyProperty{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, MetaProperty{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, MesaProperty{})
}
