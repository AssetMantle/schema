// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package capabilities

import (
	"github.com/AssetMantle/schema/go/properties"
)

type Splittable interface {
	GetSupply() properties.Property
}
