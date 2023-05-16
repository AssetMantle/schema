// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"sort"
	"strings"

	"github.com/AssetMantle/schema/go/data"
	dataConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/data/utilities"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ data.ListData = (*ListData)(nil)

func (listData *ListData) ValidateWithType(expectedTypeID ids.StringID) error {
	for _, anyListableDatum := range listData.AnyListableData {
		if anyListableDatum.GetTypeID().Compare(expectedTypeID) == 0 {
			return errorConstants.IncorrectFormat.Wrapf("data type doesnt conform to expected type for list")
		}
	}

	return listData.ValidateBasic()
}
func (listData *ListData) ValidateBasic() error {
	for _, anyListableData := range listData.AnyListableData {
		if err := anyListableData.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (listData *ListData) Get() []data.AnyListableData {
	anyListableData := make([]data.AnyListableData, len(listData.AnyListableData))

	for i, anyListableDataList := range listData.AnyListableData {
		anyListableData[i] = anyListableDataList
	}

	return anyListableData
}
func (listData *ListData) GetBondWeight() sdkTypes.Int {
	return dataConstants.ListDataWeight
}
func (listData *ListData) AsString() string {
	sortedListData := NewListData(anyListableDataToListableData(listData.Get()...)...).(*ListData)
	dataStrings := make([]string, len(sortedListData.AnyListableData))

	for i, anyListableData := range sortedListData.AnyListableData {
		dataStrings[i] = anyListableData.AsString()
	}

	return utilities.JoinListStrings(dataStrings...)
}
func (listData *ListData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeListData(), nil
	}

	dataStringList := utilities.SplitListString(dataString)
	dataList := make([]data.ListableData, len(dataStringList))

	for i, datumString := range dataStringList {
		if datum, err := PrototypeAnyData().FromString(datumString); err != nil {
			return PrototypeListData(), err
		} else if listableData, ok := datum.(data.ListableData); !ok {
			return PrototypeListData(), errorConstants.IncorrectFormat.Wrapf("data type %T is not listable", datum)
		} else {
			dataList[i] = listableData
		}
	}

	return NewListData(dataList...), nil
}
func (listData *ListData) Search(data data.ListableData) (int, bool) {
	sortedListData := NewListData(anyListableDataToListableData(listData.Get()...)...).(*ListData)

	index := sort.Search(
		len(sortedListData.AnyListableData),
		func(i int) bool {
			return sortedListData.AnyListableData[i].Compare(data) >= 0
		},
	)

	if index < len(sortedListData.AnyListableData) && sortedListData.AnyListableData[index].Compare(data) == 0 {
		return index, true
	}

	return index, false
}
func (listData *ListData) Add(listableData ...data.ListableData) data.ListData {
	updatedListData := NewListData(anyListableDataToListableData(listData.Get()...)...).(*ListData)

	for _, datum := range listableData {
		if index, found := updatedListData.Search(datum); !found {
			updatedListData.AnyListableData = append(updatedListData.AnyListableData, datum.ToAnyListableData().(*AnyListableData))
			copy(updatedListData.AnyListableData[index+1:], updatedListData.AnyListableData[index:])
			updatedListData.AnyListableData[index] = datum.ToAnyListableData().(*AnyListableData)
		}
	}

	return updatedListData
}
func (listData *ListData) Remove(data ...data.ListableData) data.ListData {
	updatedListData := NewListData(anyListableDataToListableData(listData.Get()...)...).(*ListData)

	for _, listable := range data {
		if index, found := updatedListData.Search(listable); found {
			updatedListData.AnyListableData = append(updatedListData.AnyListableData[:index], updatedListData.AnyListableData[index+1:]...)
		}
	}

	return updatedListData
}
func (listData *ListData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(listData)
}
func (listData *ListData) Bytes() []byte {
	bytesList := make([][]byte, len(listData.AnyListableData))

	for i, anyListableData := range listData.AnyListableData {
		if anyListableData != nil {
			bytesList[i] = anyListableData.Bytes()
		}
	}
	return bytes.Join(bytesList, nil)
}
func (listData *ListData) GetTypeID() ids.StringID {
	return dataConstants.ListDataTypeID
}
func (listData *ListData) ZeroValue() data.Data {
	return NewListData([]data.ListableData{}...)
}
func (listData *ListData) GenerateHashID() ids.HashID {
	if len(listData.AnyListableData) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(listData.Bytes())
}
func (listData *ListData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_ListData{
			ListData: listData,
		},
	}
}
func listableDataToAnyListableData(listableData ...data.ListableData) []*AnyListableData {
	anyListableData := make([]*AnyListableData, len(listableData))

	for i, listableDatum := range listableData {
		anyListableData[i] = listableDatum.ToAnyListableData().(*AnyListableData)
	}

	return anyListableData
}
func anyListableDataToListableData(anyListableData ...data.AnyListableData) []data.ListableData {
	listableData := make([]data.ListableData, len(anyListableData))

	for i, anyListableDatum := range anyListableData {
		listableData[i] = anyListableDatum.Get().(data.ListableData)
	}

	return listableData
}
func PrototypeListData() data.ListData {
	return (&ListData{}).ZeroValue().(data.ListData)
}

// NewListData
// * onus of ensuring all data are of the same type is on DataList
func NewListData(listableData ...data.ListableData) data.ListData {
	sort.Slice(listableData, func(i, j int) bool {
		return listableData[i].Compare(listableData[j]) < 0
	})

	return &ListData{listableDataToAnyListableData(listableData...)}
}
