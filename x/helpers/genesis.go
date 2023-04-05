package helpers

import (
	"context"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"

	"github.com/AssetMantle/schema/x/parameters"
)

type Genesis interface {
	proto.Message

	Default() Genesis

	ValidateBasic(ParameterManager) error

	Import(context.Context, Mapper, ParameterManager)
	Export(context.Context, Mapper, ParameterManager) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize([]Mappable, parameters.ParameterList) Genesis
}
