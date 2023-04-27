// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/parameters"
	"github.com/AssetMantle/schema/go/properties/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func dummyValidator(interface{}) error {
	return nil
}

func createTestInput() (*baseIDs.StringID, data.Data, parameters.Parameter) {
	id := baseIDs.NewStringID("ID")
	stringData := baseData.NewStringData("Data")

	testParameter := NewParameter(base.NewMetaProperty(id, stringData))
	return id.(*baseIDs.StringID), stringData, testParameter
}

func TestNewParameter(t *testing.T) {
	id, testData, _ := createTestInput()
	type args struct {
		id        ids.StringID
		data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name    string
		args    args
		want    parameters.Parameter
		wantErr bool
	}{

		{"+ve", args{id, testData, dummyValidator}, &Parameter{MetaProperty: base.NewMetaProperty(id, testData).(*base.MetaProperty)}, false},
		{"panic empty", args{}, &Parameter{}, true},
		{"nil", args{nil, nil, nil}, &Parameter{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("Error %v", r)
				}
			}()
			if got := NewParameter(base.NewMetaProperty(tt.args.id, tt.args.data)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_Mutate(t *testing.T) {
	id, testData, _ := createTestInput()
	newData := baseData.NewStringData("Data")
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	type args struct {
		data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   parameters.Parameter
	}{

		{name: "+ve", fields: fields{id, testData, dummyValidator}, args: args{newData}, want: &Parameter{MetaProperty: base.NewMetaProperty(id, testData).(*base.MetaProperty)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{base.NewMetaProperty(tt.fields.ID, tt.fields.Data).(*base.MetaProperty)}
			if got := parameter.Mutate(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_Validate(t *testing.T) {
	id, testData, _ := createTestInput()
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with stringData", fields{id, testData, dummyValidator}, false},
		{"+ve with decData", fields{baseIDs.NewStringID("ID"), baseData.NewDecData(sdkTypes.SmallestDec()), dummyValidator}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{base.NewMetaProperty(tt.fields.ID, tt.fields.Data).(*base.MetaProperty)}
			if err := parameter.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
