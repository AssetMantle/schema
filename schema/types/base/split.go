package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids/base"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/types"
)

// type split struct {
//	OwnerID   ids.IdentityID
//	OwnableID ids.OwnableID
//	Value     sdkTypes.Dec
// }

var _ types.Split = (*Split)(nil)

func (split *Split) ValidateBasic() error {
	if err := split.OwnerID.ValidateBasic(); err != nil {
		return err
	}
	if err := split.OwnableID.ValidateBasic(); err != nil {
		return err
	}
	if value, err := sdkTypes.NewDecFromStr(split.Value); err != nil {
		return err
	} else if !sdkTypes.ValidSortableDec(value) {
		return errorConstants.IncorrectFormat
	}
	return nil
}
func (split *Split) GetOwnerID() ids.IdentityID {
	return split.OwnerID
}
func (split *Split) GetOwnableID() ids.OwnableID {
	return split.OwnableID
}
func (split *Split) GetValue() sdkTypes.Dec {
	value, _ := sdkTypes.NewDecFromStr(split.Value)
	return value
}
func (split *Split) Send(outValue sdkTypes.Dec) types.Split {
	value, _ := sdkTypes.NewDecFromStr(split.Value)
	split.Value = value.Sub(outValue).String()
	return split
}
func (split *Split) Receive(inValue sdkTypes.Dec) types.Split {
	value, _ := sdkTypes.NewDecFromStr(split.Value)
	split.Value = value.Add(inValue).String()
	return split
}
func (split *Split) CanSend(outValue sdkTypes.Dec) bool {
	value, _ := sdkTypes.NewDecFromStr(split.Value)
	return value.GTE(outValue)
}

func NewSplit(ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) types.Split {
	return &Split{
		OwnerID:   ownerID.(*base.IdentityID),
		OwnableID: ownableID.ToAnyOwnableID().(*base.AnyOwnableID),
		Value:     value.String(),
	}
}
