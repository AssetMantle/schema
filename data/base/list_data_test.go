// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"fmt"
	"github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var (
	testListData       = NewListData(NewStringData("Data2"), NewStringData("Data1"))
	sortedTestListData = NewListData(NewStringData("Data1"), NewStringData("Data2"))
)

func Test_ListDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want bool
	}{
		{"+ve", NewListData(NewStringData("Data")), false},
		{"-ve", NewListData(&HeightData{&base.Height{-10}}), true},
		{"-ve", &ListData{[]*AnyListableData{(&StringData{"Data"}).ToAnyListableData().(*AnyListableData), (&HeightData{&base.Height{10}}).ToAnyListableData().(*AnyListableData)}}, true},
		{"-ve", &ListData{[]*AnyListableData{(&IDData{(&baseIDs.StringID{"ID1"}).ToAnyID().(*baseIDs.AnyID)}).ToAnyListableData().(*AnyListableData), (&IDData{(&baseIDs.HashID{testBytes}).ToAnyID().(*baseIDs.AnyID)}).ToAnyListableData().(*AnyListableData)}}, true},
	}
	for _, tt := range tests {
		err := tt.args.ValidateBasic()
		if err != nil {
			fmt.Println(err.Error())
		}
		if err == nil && tt.want {
			t.Errorf("ListDataValidateBasic() got = %v, want = %v", err, tt.want)
		}
		if err != nil && !tt.want {
			t.Errorf("ListDataValidateBasic() got = %v, want = %v", err, tt.want)
		}
	}
}
func TestListDataPrototype(t *testing.T) {
	type args struct {
		value data.ListData
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve", args{}, &ListData{[]*AnyListableData{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrototypeListData()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListDataPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewListData(t *testing.T) {
	tests := []struct {
		name string
		args []data.ListableData
		want data.ListData
	}{
		{"+ve for some id", []data.ListableData{NewStringData("Data1"), NewStringData("Data2")}, &ListData{[]*AnyListableData{NewStringData("Data1").ToAnyListableData().(*AnyListableData), NewStringData("Data2").ToAnyListableData().(*AnyListableData)}}},
		{"+ve for empty String", []data.ListableData{NewStringData(""), NewStringData("Data1")}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData), NewStringData("Data1").ToAnyListableData().(*AnyListableData)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListData(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ListDataSearch(t *testing.T) {
	tests := []struct {
		name   string
		args   data.ListData
		search data.ListableData
		want1  int
		want2  bool
	}{
		{"+ve", NewListData([]data.ListableData{NewStringData("Data1"), NewStringData("Data2")}...), NewStringData("Data1"), 0, true},
		{"+ve", NewListData([]data.ListableData{NewStringData("Data1"), NewStringData("Data2")}...), NewStringData("Data2"), 1, true},
		{"+ve", NewListData([]data.ListableData{NewStringData("Data1"), NewStringData("")}...), NewStringData(""), 0, true},
		{"+ve", NewListData([]data.ListableData{NewStringData("Data1"), NewStringData("")}...), NewStringData("test"), 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.args.Search(tt.search)
			if !reflect.DeepEqual(got1, tt.want1) || !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ListDataSearch() got1 = %v, got2 = %v, want1 = %v want2 = %v", got1, got2, tt.want1, tt.want2)
			}
		})
	}
}

func Test_ListDataAdd(t *testing.T) {
	tests := []struct {
		name string
		args []data.ListableData
		want data.ListData
	}{
		{"+ve", []data.ListableData{NewStringData("Data2"), NewHeightData(base.NewHeight(10)), NewStringData("Data1")}, &ListData{[]*AnyListableData{NewStringData("Data1").ToAnyListableData().(*AnyListableData), NewStringData("Data2").ToAnyListableData().(*AnyListableData)}}},
		{"+ve", []data.ListableData{NewStringData("Data"), NewStringData(""), NewStringData("Data")}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData), NewStringData("Data").ToAnyListableData().(*AnyListableData)}}},
		{"+ve", []data.ListableData{NewStringData("Data")}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}},
		{"+ve", []data.ListableData{NewStringData("")}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}},
		{"-ve", []data.ListableData{NewStringData("Data")}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}},
		{"-ve", []data.ListableData{NewStringData("Data"), NewStringData("Data")}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}},
		{"-ve", []data.ListableData{NewStringData("Data"), NewNumberData(sdkTypes.NewInt(1))}, &ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}},
		{"-ve", []data.ListableData{NewIDData(baseIDs.NewStringID("ID1")), NewIDData(baseIDs.GenerateHashID(testBytes))}, &ListData{[]*AnyListableData{NewIDData(baseIDs.NewStringID("ID1")).ToAnyListableData().(*AnyListableData)}}},
		{"+ve", []data.ListableData{NewStringData("Data4"), NewStringData("Data3"), NewStringData("Data2"), NewHeightData(base.NewHeight(10)), NewStringData("Data1")}, &ListData{[]*AnyListableData{NewStringData("Data1").ToAnyListableData().(*AnyListableData), NewStringData("Data2").ToAnyListableData().(*AnyListableData), NewStringData("Data3").ToAnyListableData().(*AnyListableData), NewStringData("Data4").ToAnyListableData().(*AnyListableData)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := &ListData{}
			if got := listData.Add(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListDataAdd got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_ListDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want []byte
	}{
		{"+ve", NewListData(NewStringData("Data")), NewStringData("Data").Bytes()}, // for a single data no loop iteration is required, so directly it's byte should match
		{"+ve", testListData, bytes.Join([][]byte{NewStringData("Data1").Bytes(), NewStringData("Data2").Bytes()}, dataConstants.ListBytesSeparator)},
		{"+ve", NewListData(NewStringData("")), []byte(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListDataBytes() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want ids.HashID
	}{
		{"+ve", testListData, baseIDs.GenerateHashID(testListData.Bytes())},
		{"+ve", NewListData(NewStringData("")), baseIDs.GenerateHashID(NewStringData("").Bytes())},
		{"+ve", NewListData(), baseIDs.GenerateHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listData_GenerateHashID() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_Get(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want []data.AnyListableData
	}{
		{"+ve", NewListData(NewStringData("Data")), (&ListData{[]*AnyListableData{NewStringData("Data").ToAnyListableData().(*AnyListableData)}}).Get()},
		{"+ve", NewListData(NewStringData("")), (&ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}).Get()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listData_Get() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_GetID(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want ids.DataID
	}{
		{"+ve", testListData, baseIDs.GenerateDataID(&ListData{[]*AnyListableData{NewStringData("Data1").ToAnyListableData().(*AnyListableData), NewStringData("Data2").ToAnyListableData().(*AnyListableData)}})},
		{"+ve", NewListData(NewStringData("")), baseIDs.GenerateDataID(&ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listData_GetID() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_GetType(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want ids.StringID
	}{
		{"+ve", testListData, dataConstants.ListDataTypeID},
		{"+ve", NewListData(NewHeightData(base.NewHeight(10))), dataConstants.ListDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listData_GetType() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_Remove(t *testing.T) {
	tests := []struct {
		name   string
		args   data.ListData
		remove []data.ListableData
		want   data.ListData
	}{
		{"+ve", NewListData(NewStringData("")), []data.ListableData{}, &ListData{[]*AnyListableData{NewStringData("").ToAnyListableData().(*AnyListableData)}}},
		{"+ve", NewListData(NewStringData("")), []data.ListableData{NewStringData("")}, &ListData{[]*AnyListableData{}}},
		{"+ve", NewListData(NewStringData("data")), []data.ListableData{NewStringData("data")}, &ListData{[]*AnyListableData{}}},
		{"+ve", testListData, []data.ListableData{NewStringData("Data2")}, &ListData{[]*AnyListableData{NewStringData("Data1").ToAnyListableData().(*AnyListableData)}}},
		{"+ve", testListData, []data.ListableData{NewStringData("Data1"), NewStringData("Data2")}, &ListData{[]*AnyListableData{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Remove(tt.remove...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listData_Remove() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_AsString(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want string
	}{
		{"+ve", testListData, "Data1,,Data2"},
		{"+ve", NewListData(NewHeightData(base.NewHeight(20)), NewHeightData(base.NewHeight(11))), "11,,20"},
		{"+ve", NewListData(NewStringData("")), ""},
		{"+ve", NewListData(NewHeightData(base.NewHeight(-1))), "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listData_AsString() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_listData_ZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.ListData
		want data.Data
	}{
		{"+ve", NewListData(NewStringData("Data")), NewListData([]data.ListableData{}...)},
		{"+ve", NewListData(NewStringData("")), NewListData([]data.ListableData{}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listData_ZeroValue() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_ListDataFromString(t *testing.T) {
	tests := []struct {
		name      string
		args      string
		want      data.Data
		wantError bool
	}{
		{"+ve", "", &ListData{[]*AnyListableData{}}, false},
		{"+ve", "S|Data1,,S|Data2", sortedTestListData, false},
		{"-ve", "L|S|test", sortedTestListData, true},
		{"-ve", "K|test,,H|10", testListData, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeListData().FromString(tt.args)
			if !tt.wantError {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ListDataFromString() got = %v, want = %v", got, tt.want)
				}
			}
			if err != nil && !tt.wantError {
				t.Errorf("ListDataFromString() got = %v, want = %v", err, tt.wantError)
			}
			if err == nil && tt.wantError {
				t.Errorf("ListDataFromString() got = %v, want = %v", err, tt.wantError)
			}
		})
	}
}
