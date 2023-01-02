// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package genesis

import (
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

// TODO ***** add default classes NUB,maintainer, self class
func Prototype() helpers.Genesis {
	return baseHelpers.NewGenesis(key.Prototype, PrototypeGenesisState())
}
