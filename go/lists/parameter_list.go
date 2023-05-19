package lists

import (
	"github.com/gogo/protobuf/proto"

	"github.com/AssetMantle/schema/go/types"
)

type ParameterList interface {
	proto.Message
	Get() []types.Parameter
}
