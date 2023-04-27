package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/types"
)

var _ types.Split = (*Split)(nil)

func (split *Split) ValidateBasic() error {
	if err := split.OwnerID.ValidateBasic(); err != nil {
		return err
	}
	if err := split.OwnableID.ValidateBasic(); err != nil {
		return err
	}
	if _, err := sdkTypes.NewDecFromStr(split.Value); err != nil {
		return err
	}
	return nil
}
func (split *Split) GetOwnerID() ids.IdentityID {
	return split.OwnerID
}
func (split *Split) GetOwnableID() ids.OwnableID {
	return split.OwnableID
}
func (split *Split) GetValue() sdkTypes.Int {
	if value, ok := sdkTypes.NewIntFromString(split.Value); !ok {
		panic("invalid split value")
	} else {
		return value
	}
}
func (split *Split) Send(outValue sdkTypes.Int) types.Split {
	split.Value = split.GetValue().Sub(outValue).String()
	return split
}
func (split *Split) Receive(inValue sdkTypes.Int) types.Split {
	split.Value = split.GetValue().Add(inValue).String()
	return split
}
func (split *Split) CanSend(outValue sdkTypes.Int) bool {
	return split.GetValue().GTE(outValue)
}

func NewSplit(ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Int) types.Split {
	return &Split{
		OwnerID:   ownerID.(*baseIDs.IdentityID),
		OwnableID: ownableID.ToAnyOwnableID().(*baseIDs.AnyOwnableID),
		Value:     value.String(),
	}
}
