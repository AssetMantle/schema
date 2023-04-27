package base

import (
	"github.com/AssetMantle/schema/go/errors"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ errors.Error = (*sdkErrors.Error)(nil)

func NewError(codeSpace string, code uint32, description string) errors.Error {
	return sdkErrors.Register(codeSpace, code, description)
}
