package base

import (
	"fmt"
	"github.com/AssetMantle/schema/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ types.Split = (*Split)(nil)

func (split *Split) ValidateBasic() error {
	if _, ok := sdkTypes.NewIntFromString(split.Value); !ok {
		return fmt.Errorf("invalid split value %s", split.Value)
	}
	return nil
}
func (split *Split) GetValue() sdkTypes.Int {
	if value, ok := sdkTypes.NewIntFromString(split.Value); !ok {
		panic("invalid split value")
	} else {
		return value
	}
}
func (split *Split) Subtract(outValue sdkTypes.Int) types.Split {
	split.Value = split.GetValue().Sub(outValue).String()
	return split
}
func (split *Split) Add(inValue sdkTypes.Int) types.Split {
	split.Value = split.GetValue().Add(inValue).String()
	return split
}
func (split *Split) CanSend(outValue sdkTypes.Int) bool {
	return split.GetValue().GTE(outValue)
}

func NewSplit(value sdkTypes.Int) types.Split {
	return &Split{
		Value: value.String(),
	}
}
