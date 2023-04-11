package parameters

import (
	"github.com/gogo/protobuf/proto"
)

type ParameterList interface {
	proto.Message
	Get() []Parameter
}
