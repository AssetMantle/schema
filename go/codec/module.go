package codec

import (
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterModuleConcrete(legacyAmino *codec.LegacyAmino, o interface{}) {
	legacyAmino.RegisterConcrete(o, reflect.TypeOf(o).PkgPath()+"/"+reflect.TypeOf(o).Name(), nil)
}
