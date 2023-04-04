package helpers

import (
	"github.com/AssetMantle/schema/x/data"
)

type ValidatableParameter interface {
	GetParameter() Parameter
	Mutate(data.Data) ValidatableParameter
	GetValidator() func(i interface{}) error
	Validate() error
}
