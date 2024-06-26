// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/properties"
)

type Parameter interface {
	proto.Message
	ValidateBasic() error
	GetMetaProperty() properties.MetaProperty
	Mutate(data.Data) Parameter
}
