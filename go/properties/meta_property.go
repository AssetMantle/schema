// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"github.com/AssetMantle/schema/go/data"
)

type MetaProperty interface {
	IsMetaProperty()
	FromString(string) (MetaProperty, error)
	GetData() data.AnyData
	ScrubData() MesaProperty
	Property
}
