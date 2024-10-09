package base

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"strings"
)

type dataGetter interface {
	get() data.Data
}

func (x *AnyData_AccAddressData) get() data.Data {
	return x.AccAddressData
}
func (x *AnyData_BooleanData) get() data.Data {
	return x.BooleanData
}
func (x *AnyData_DecData) get() data.Data {
	return x.DecData
}
func (x *AnyData_HeightData) get() data.Data {
	return x.HeightData
}
func (x *AnyData_IDData) get() data.Data {
	return x.IDData
}
func (x *AnyData_NumberData) get() data.Data {
	return x.NumberData
}
func (x *AnyData_StringData) get() data.Data {
	return x.StringData
}
func (x *AnyData_ListData) get() data.Data {
	return x.ListData
}
func (x *AnyData_LinkedData) get() data.Data {
	return x.LinkedData
}

var _ data.AnyData = (*AnyData)(nil)

func (anyData *AnyData) ValidateBasic() error {
	if anyData.Impl == nil {
		return nil
	}

	return anyData.Impl.(dataGetter).get().ValidateBasic()
}
func (anyData *AnyData) IsAnyData() {}
func (anyData *AnyData) AsString() string {
	return joinDataTypeAndValueStrings(anyData.Impl.(dataGetter).get().GetTypeID().AsString(), anyData.Impl.(dataGetter).get().AsString())
}
func (anyData *AnyData) FromString(dataString string) (data.Data, error) {
	dataTypeString, dataValueString := splitDataTypeAndValueStrings(dataString)
	if dataTypeString != "" {
		var Data data.Data
		var err error

		if strings.HasSuffix(dataTypeString, PrototypeIDData().GetTypeID().AsString()) {
			Data, err = PrototypeIDData().FromString(dataString)
		} else {
			switch baseIDs.NewStringID(dataTypeString).AsString() {
			case PrototypeAccAddressData().GetTypeID().AsString():
				Data, err = PrototypeAccAddressData().FromString(dataValueString)
			case PrototypeBooleanData().GetTypeID().AsString():
				Data, err = PrototypeBooleanData().FromString(dataValueString)
			case PrototypeDecData().GetTypeID().AsString():
				Data, err = PrototypeDecData().FromString(dataValueString)
			case PrototypeHeightData().GetTypeID().AsString():
				Data, err = PrototypeHeightData().FromString(dataValueString)
			case PrototypeListData().GetTypeID().AsString():
				Data, err = PrototypeListData().FromString(dataValueString)
			case PrototypeLinkedData().GetTypeID().AsString():
				Data, err = PrototypeLinkedData().FromString(dataValueString)
			case PrototypeNumberData().GetTypeID().AsString():
				Data, err = PrototypeNumberData().FromString(dataValueString)
			case PrototypeStringData().GetTypeID().AsString():
				Data, err = PrototypeStringData().FromString(dataValueString)
			default:
				Data, err = nil, fmt.Errorf("data type %s identifier is not recognized", dataTypeString)
			}
		}
		if err != nil {
			return nil, err
		}

		return Data.ToAnyData(), nil
	}

	if dataValueString == "" {
		return PrototypeAnyData(), nil
	}

	return nil, fmt.Errorf("type identifier is missing")
}
func (anyData *AnyData) Get() data.Data {
	return anyData.Impl.(dataGetter).get()
}
func (anyData *AnyData) GetID() ids.DataID {
	if anyData.Impl == nil {
		return &baseIDs.DataID{
			TypeID: anyData.GetTypeID().(*baseIDs.StringID),
			HashID: baseIDs.PrototypeHashID().(*baseIDs.HashID),
		}
	}

	return anyData.Impl.(dataGetter).get().GetID()
}
func (anyData *AnyData) Bytes() []byte {
	return anyData.Impl.(dataGetter).get().Bytes()
}
func (anyData *AnyData) GetTypeID() ids.StringID {
	if anyData.Impl != nil {
		return anyData.Impl.(dataGetter).get().GetTypeID()
	}

	return dataConstants.AnyDataTypeID
}
func (anyData *AnyData) ZeroValue() data.Data {
	return anyData.Impl.(dataGetter).get().ZeroValue()
}
func (anyData *AnyData) GenerateHashID() ids.HashID {
	return anyData.Impl.(dataGetter).get().GenerateHashID()
}
func (anyData *AnyData) ToAnyData() data.AnyData {
	return anyData
}
func (anyData *AnyData) GetBondWeight() math.Int {
	return anyData.Impl.(dataGetter).get().GetBondWeight()
}

func PrototypeAnyData() data.AnyData {
	return &AnyData{}
}
func joinDataTypeAndValueStrings(dataType, dataValue string) string {
	return strings.TrimSpace(dataType) + dataTypeAndValueSeparator + strings.TrimSpace(dataValue)
}
func splitDataTypeAndValueStrings(dataTypeAndValueString string) (dataType, dataValue string) {
	if dataTypeAndValue := strings.SplitN(dataTypeAndValueString, dataTypeAndValueSeparator, 2); len(dataTypeAndValue) == 1 {
		return strings.TrimSpace(dataTypeAndValue[0]), ""
	} else if len(dataTypeAndValue) == 2 {
		return strings.TrimSpace(dataTypeAndValue[0]), strings.TrimSpace(dataTypeAndValue[1])
	} else {
		return "", ""
	}
}

const dataTypeAndValueSeparator = "|"
