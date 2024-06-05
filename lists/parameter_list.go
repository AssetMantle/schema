package lists

import (
	"github.com/gogo/protobuf/proto"

	"github.com/AssetMantle/schema/parameters"
)

type ParameterList interface {
	proto.Message
	Get() []parameters.Parameter
}
