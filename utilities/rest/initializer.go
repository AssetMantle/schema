package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"

	"github.com/AssetMantle/schema/utilities/rest/keys/add"
	"github.com/AssetMantle/schema/utilities/rest/sign"
)

func RegisterRESTRoutes(context client.Context, router *mux.Router) {
	add.RegisterRESTRoutes(context, router)
	sign.RegisterRESTRoutes(context, router)
}