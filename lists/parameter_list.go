package lists

import (
	"github.com/AssetMantle/schema/parameters"
	"github.com/cosmos/gogoproto/proto"
)

type ParameterList interface {
	proto.Message
	Get() []parameters.Parameter
}
