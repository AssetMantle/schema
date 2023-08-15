// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

func Test_NewIDData(t *testing.T) {
	tests := []struct {
		name string
		args ids.ID
		want data.Data
	}{
		{name: "+ve", args: baseIDs.NewStringID("test"), want: &IDData{(&baseIDs.StringID{"test"}).ToAnyID().(*baseIDs.AnyID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIDData(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIDData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want bool
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), false},
		//{"-ve", NewIDData(baseIDs.NewStringID("test///")), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err == nil && tt.want {
				t.Errorf("IDDataValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("IDDataValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_IDData_Compare(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want bool
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), true},
		{"+ve", NewIDData(baseIDs.NewStringID("test2")), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(NewIDData(baseIDs.NewStringID("test")))
			if (got == 0) != tt.want {
				t.Errorf("IDData_Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDData_GenerateHashID(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want ids.ID
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), baseIDs.GenerateHashID(NewIDData(baseIDs.NewStringID("test")).Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDData_GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataGet(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want ids.AnyID
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), (&baseIDs.StringID{"test"}).ToAnyID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDDataGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataGetID(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want ids.DataID
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), &baseIDs.DataID{
			TypeID: &baseIDs.StringID{"SI"},
			HashID: baseIDs.GenerateHashID(NewIDData(baseIDs.NewStringID("test")).Bytes()).(*baseIDs.HashID),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDDataGetID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataGetType(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want ids.ID
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), &baseIDs.StringID{"SI"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDDataGetTypeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataAsString(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want string
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), "SI|test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDDataAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want data.Data
	}{
		{"+ve", NewIDData(baseIDs.NewStringID("test")), &IDData{&baseIDs.AnyID{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ZeroValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDDataZeroValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataBytes(t *testing.T) {
	tests := []struct {
		name string
		args data.IDData
		want []byte
	}{
		{name: "+ve", args: NewIDData(baseIDs.NewStringID("test")), want: []byte("test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDDataBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IDDataFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    data.Data
		wantErr bool
	}{
		{"+ve", "SI|id", &IDData{baseIDs.NewStringID("id").ToAnyID().(*baseIDs.AnyID)}, false},
		{"-ve", "U|id", PrototypeIDData(), true},
		{"-ve", "abc", PrototypeIDData(), true},
		{"+ve", "", PrototypeIDData(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeIDData().FromString(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDDataFromString() got = %v, want %v", got, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("IDDataFromString() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("IDDataFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
