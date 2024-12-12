// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/AssetMantle/schema/types"
)

var _ types.Split = (*Split)(nil)

func (split *Split) ValidateBasic() error {
	if _, ok := math.NewIntFromString(split.Value); !ok {
		return fmt.Errorf("invalid split value %s", split.Value)
	}
	return nil
}
func (split *Split) GetValue() math.Int {
	if value, ok := math.NewIntFromString(split.Value); !ok {
		panic("invalid split value")
	} else {
		return value
	}
}
func (split *Split) Subtract(outValue math.Int) types.Split {
	split.Value = split.GetValue().Sub(outValue).String()
	return split
}
func (split *Split) Add(inValue math.Int) types.Split {
	split.Value = split.GetValue().Add(inValue).String()
	return split
}
func (split *Split) CanSend(outValue math.Int) bool {
	return split.GetValue().GTE(outValue)
}

func NewSplit(value math.Int) types.Split {
	return &Split{
		Value: value.String(),
	}
}
