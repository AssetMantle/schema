// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"

	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/AssetMantle/modules/schema/ids"
)

type ParameterList interface {
	Get() []Parameter
	GetParameter(ids.PropertyID) Parameter

	Fetch(context.Context) ParameterList
	Set(context.Context, ...Parameter)

	GetKeyTable() paramsTypes.KeyTable

	Initialize(paramsTypes.Subspace) ParameterList
}
