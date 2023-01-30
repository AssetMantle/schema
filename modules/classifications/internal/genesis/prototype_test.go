// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package genesis

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"

	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.NotPanics(t, func() {
		require.Equal(t, Prototype(), baseHelpers.NewGenesis(key.Prototype, PrototypeGenesisState()))
	})
}
