package lists

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/AssetMantle/schema/parameters"
)

type ParameterList interface {
	proto.Message
	Get() []parameters.Parameter
}
