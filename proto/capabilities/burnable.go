// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package capabilities

import (
	"github.com/AssetMantle/schema/x/properties"
)

type Burnable interface {
	GetBurn() properties.Property
}
