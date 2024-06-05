package base

import (
	"fmt"
	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

var _ data.NumberData = (*NumberData)(nil)

func (numberData *NumberData) ValidateBasic() error {
	if _, ok := sdkTypes.NewIntFromString(numberData.Value); !ok {
		return fmt.Errorf("number data value %s is not a valid integer", numberData.Value)
	}

	return nil
}
func (numberData *NumberData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(numberData)
}
func (numberData *NumberData) GetBondWeight() sdkTypes.Int {
	return dataConstants.NumberDataWeight
}
func (numberData *NumberData) AsString() string {
	return numberData.Value
}
func (numberData *NumberData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeNumberData(), nil
	}

	value, ok := sdkTypes.NewIntFromString(dataString)
	if !ok {
		return PrototypeNumberData(), fmt.Errorf("number data value %s is not a valid integer", dataString)
	}

	numberData.Value = value.String()
	if err := numberData.ValidateBasic(); err != nil {
		return PrototypeNumberData(), err
	}

	return numberData, nil
}
func (numberData *NumberData) Bytes() []byte {
	return []byte(numberData.Value)
}
func (numberData *NumberData) GetTypeID() ids.StringID {
	return dataConstants.NumberDataTypeID
}
func (numberData *NumberData) ZeroValue() data.Data {
	return NewNumberData(sdkTypes.ZeroInt())
}
func (numberData *NumberData) GenerateHashID() ids.HashID {
	if numberData.Compare(numberData.ZeroValue().(data.ListableData)) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(numberData.Bytes())
}
func (numberData *NumberData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_NumberData{
			NumberData: numberData,
		},
	}
}
func (numberData *NumberData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_NumberData{
			NumberData: numberData,
		},
	}
}
func (numberData *NumberData) Compare(listableData data.ListableData) int {
	if difference := numberData.Get().Sub(listableData.ToAnyListableData().Get().(*NumberData).Get()); difference.IsZero() {
		return 0
	} else if difference.IsPositive() {
		return 1
	}
	return -1
}
func (numberData *NumberData) Get() sdkTypes.Int {
	if value, ok := sdkTypes.NewIntFromString(numberData.Value); !ok {
		panic("invalid number data")
	} else {
		return value
	}
}

func PrototypeNumberData() data.NumberData {
	return NewNumberData(sdkTypes.ZeroInt()).ZeroValue().(*NumberData)
}

func NewNumberData(value sdkTypes.Int) data.NumberData {
	return &NumberData{
		Value: value.String(),
	}
}
