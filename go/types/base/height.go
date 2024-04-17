// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/AssetMantle/schema/go/types"
	"github.com/AssetMantle/schema/go/types/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ types.Height = (*Height)(nil)

func (height *Height) ValidateBasic() error {
	if height.Value < constants.InfiniteHeight {
		return fmt.Errorf("height value %d is out of bound", height.Value)
	}
	return nil
}
func (height *Height) Bytes() []byte {
	Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(Bytes, uint64(height.Get()))
	return Bytes
}
func (height *Height) Compare(compareHeight types.Height) int {
	if difference := height.Get() - compareHeight.Get(); difference == 0 {
		return 0
	} else if difference > 0 || height.Get() == constants.InfiniteHeight {
		return 1
	}

	return -1
}
func (height *Height) Get() int64 { return height.Value }

func NewHeight(value int64) types.Height {
	if value < 0 {
		value = constants.InfiniteHeight
	}

	return &Height{Value: value}
}

func CurrentHeight(context context.Context) types.Height {
	return NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())
}
