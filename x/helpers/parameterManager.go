// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/AssetMantle/schema/x/ids"
	"github.com/AssetMantle/schema/x/parameters"
)

type ParameterManager interface {
	Get() parameters.ParameterList
	GetValidatableParameter(ids.PropertyID) ValidatableParameter
	GetParameter(ids.PropertyID) parameters.Parameter
	ValidateParameter(parameters.Parameter) error

	Fetch(context.Context) ParameterManager
	Set(context.Context, parameters.ParameterList)

	GetKeyTable() paramsTypes.KeyTable
	RESTQueryHandler(client.Context) http.HandlerFunc
	Initialize(paramsTypes.Subspace) ParameterManager
}
