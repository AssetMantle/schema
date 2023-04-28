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
	"github.com/AssetMantle/schema/go/traits"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ data.ListData = (*ListData)(nil)

func (listData *ListData) ValidateBasic() error {
	for _, anyListableData := range listData.AnyListableDataList {
		if err := anyListableData.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (listData *ListData) Get() []data.ListableData {
	listData.sort()
	anyDataList := make([]data.ListableData, len(listData.AnyListableDataList))
	for i, anyListableDataList := range listData.AnyListableDataList {
		anyDataList[i] = anyListableDataList
	}

	return anyDataList
}
func (listData *ListData) GetBondWeight() sdkTypes.Int {
	return dataConstants.ListDataWeight
}
func (listData *ListData) AsString() string {
	listData.sort()
	dataStrings := make([]string, len(listData.AnyListableDataList))

	for i, anyListableData := range listData.AnyListableDataList {
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
	listData.sort()
	index := sort.Search(
		len(listData.AnyListableDataList),
		func(i int) bool {
			return listData.AnyListableDataList[i].Compare(data) >= 0
		},
	)

	if index < len(listData.AnyListableDataList) && listData.AnyListableDataList[index].Compare(data) == 0 {
		return index, true
	}

	return index, false
}
func (listData *ListData) Add(listableData ...data.ListableData) data.ListData {
	updatedListData := listData.sort()
	for _, datum := range listableData {
		if index, found := updatedListData.Search(datum); !found {
			updatedListData.AnyListableDataList = append(updatedListData.AnyListableDataList, datum.ToAnyListableData().(*AnyListableData))
			copy(updatedListData.AnyListableDataList[index+1:], updatedListData.AnyListableDataList[index:])
			updatedListData.AnyListableDataList[index] = datum.ToAnyListableData().(*AnyListableData)
		}
	}

	return updatedListData
}
func (listData *ListData) Remove(data ...data.ListableData) data.ListData {
	updatedListData := listData.sort()

	for _, listable := range data {
		if index, found := updatedListData.Search(listable); found {
			updatedListData.AnyListableDataList = append(updatedListData.AnyListableDataList[:index], updatedListData.AnyListableDataList[index+1:]...)
		}
	}

	return updatedListData
}
func (listData *ListData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(listData.sort())
}
func (listData *ListData) Compare(listable traits.Listable) int {
	compareListData, err := listDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	// TODO check for optimization
	return bytes.Compare(listData.Bytes(), compareListData.Bytes())
}
func (listData *ListData) Bytes() []byte {
	bytesList := make([][]byte, len(listData.sort().AnyListableDataList))

	for i, anyListableData := range listData.AnyListableDataList {
		if anyListableData != nil {
			bytesList[i] = anyListableData.Bytes()
		}
	}
	// TODO see if separator required
	return bytes.Join(bytesList, nil)
}
func (listData *ListData) GetTypeID() ids.StringID {
	return dataConstants.ListDataTypeID
}
func (listData *ListData) ZeroValue() data.Data {
	return NewListData([]data.ListableData{}...)
}
func (listData *ListData) GenerateHashID() ids.HashID {
	if listData.Compare(listData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(listData.sort().Bytes())
}
func (listData *ListData) sort() *ListData {
	sort.Slice(listData.AnyListableDataList, func(i, j int) bool {
		return listData.AnyListableDataList[i].Compare(listData.AnyListableDataList[j]) < 0
	})

	return listData
}
func (listData *ListData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_ListData{
			ListData: listData,
		},
	}
}
func listDataFromInterface(listable traits.Listable) (*ListData, error) {
	switch value := listable.(type) {
	case *ListData:
		return value.sort(), nil
	default:
		return &ListData{}, errorConstants.IncorrectFormat.Wrapf("unsupported type")
	}
}

func PrototypeListData() data.ListData {
	return (&ListData{}).ZeroValue().(data.ListData)
}

// NewListData
// * onus of ensuring all data are of the same type is on DataList
func NewListData(data ...data.ListableData) data.ListData {
	return (&ListData{}).Add(data...)
}
