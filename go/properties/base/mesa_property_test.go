// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	dataConstants "github.com/AssetMantle/schema/go/data/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties"
)

func ValidatedID[V *baseIDs.PropertyID | *baseIDs.DataID](value any) V {
	if value == nil {
		return nil
	}
	return value.(V)
}

func createTestInputForMesaProperty() (ids.StringID, ids.PropertyID, data.Data, properties.MesaProperty) {
	testKey := baseIDs.NewStringID("ID")
	testData := baseData.NewStringData("Data")
	testPropertyID := baseIDs.NewPropertyID(testKey, testData.GetTypeID())
	testProperty := NewMesaProperty(testKey, testData)
	return testKey, testPropertyID, testData, testProperty
}

func TestNewEmptyMesaPropertyFromID(t *testing.T) {
	_, testPropertyID, _, _ := createTestInputForMesaProperty()

	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name    string
		args    args
		want    properties.Property
		wantErr bool
	}{
		{"+ve", args{testPropertyID}, &MesaProperty{ID: testPropertyID.(*baseIDs.PropertyID)}, false},
		{"panic with nil", args{}, &MesaProperty{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := NewEmptyMesaPropertyFromID(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyMesaPropertyFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMesaProperty(t *testing.T) {
	testKey, testPropertyID, testData, _ := createTestInputForMesaProperty()
	type args struct {
		key  ids.StringID
		data data.Data
	}
	tests := []struct {
		name      string
		args      args
		want      properties.Property
		wantPanic bool
	}{
		{"nil", args{}, &MesaProperty{}, true},
		{"+ve", args{testKey, testData}, &MesaProperty{testPropertyID.(*baseIDs.PropertyID), testData.GetID().(*baseIDs.DataID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("NewMesaProperty() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := NewMesaProperty(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMesaProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_Compare(t *testing.T) {
	_, testMesaPropertyID, testData, testMesaProperty := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	type args struct {
		property properties.Property
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"+ve compare with property with no Data", fields{testMesaPropertyID, testData.GetID()}, args{&MesaProperty{ID: baseIDs.NewPropertyID(baseIDs.NewStringID("ID"), baseIDs.NewStringID("S")).(*baseIDs.PropertyID)}}, 0, false},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, args{&MesaProperty{ID: baseIDs.NewPropertyID(baseIDs.NewStringID("ID"), baseIDs.NewStringID("S")).(*baseIDs.PropertyID), DataID: baseData.NewStringData("Data2").GetID().(*baseIDs.DataID)}}, 0, false},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, args{testMesaProperty}, 0, false},
		{"+ve nil dataID", fields{testMesaPropertyID, nil}, args{testMesaProperty}, 0, false},
		{"panic nil propertyID", fields{nil, testData.GetID()}, args{testMesaProperty}, 0, true},
		{"panic all nil", fields{}, args{testMesaProperty}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.Compare(tt.args.property); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetDataID(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		{"+ve with nil", fields{}, (*baseIDs.DataID)(nil)},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testData.GetID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.GetDataID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetHash(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()

	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name    string
		fields  fields
		want    ids.ID
		wantErr bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testData.GetID().GetHashID(), false},
		{"+ve nil PropertyID", fields{nil, testData.GetID()}, testData.GetID().GetHashID(), false},
		{"panic nil DataID", fields{testMesaPropertyID, nil}, testData.GetID().GetHashID(), true},
		{"panic all nil", fields{nil, nil}, testData.GetID().GetHashID(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.GetHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetID(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.PropertyID
	}{
		{"+ve with nil", fields{}, (*baseIDs.PropertyID)(nil)},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testMesaPropertyID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetKey(t *testing.T) {
	testKey, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()

	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name    string
		fields  fields
		want    ids.StringID
		wantErr bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testKey, false},
		{"+ve nil dataID", fields{testMesaPropertyID, nil}, testKey, false},
		{"panic nil propertyID", fields{nil, testData.GetID()}, testKey, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetType(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name    string
		fields  fields
		want    ids.StringID
		wantErr bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, baseIDs.NewStringID("S"), false},
		{"+ve nil dataID", fields{testMesaPropertyID, nil}, baseIDs.NewStringID("S"), false},
		{"panic nil PropertyID", fields{nil, testData.GetID()}, baseIDs.NewStringID("S"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.GetDataTypeID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTypeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_IsMesa(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.IsMesa(); got != tt.want {
				t.Errorf("IsMesa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_IsMeta(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.IsMeta(); got != tt.want {
				t.Errorf("IsMesa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MesaPropertyGetBondedWeight(t *testing.T) {
	_, _, _, testMesaProperty := createTestInputForMesaProperty()
	tests := []struct {
		name string
		args properties.MesaProperty
		want sdkTypes.Int
	}{
		{"+ve", testMesaProperty, dataConstants.StringDataWeight},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetBondWeight()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MesaPropertyGetBondedWeight() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MesaPropertyMutate(t *testing.T) {
	testKey, _, _, testMesaProperty := createTestInputForMesaProperty()
	tests := []struct {
		name    string
		args    properties.MesaProperty
		change  data.Data
		want    properties.MesaProperty
		wantErr bool
	}{
		{"+ve", testMesaProperty, baseData.NewStringData("Data2"), NewMesaProperty(testKey, baseData.NewStringData("Data2")), false},
		{"-ve", testMesaProperty, baseData.NewNumberData(sdkTypes.NewInt(10)), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.Mutate(tt.change)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MesaPropertyGetBondedWeight() got = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("MesaPropertyMutate() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("MesaPropertyMutate() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_MesaPropertyValidateBasic(t *testing.T) {
	_, testPropertyID, testData, testMesaProperty := createTestInputForMesaProperty()
	tests := []struct {
		name string
		args properties.MesaProperty
		want bool
	}{
		{"+ve", testMesaProperty, false},
		// Regex not work in testing
		{"-ve", &MesaProperty{
			ID:     baseIDs.NewPropertyID(baseIDs.NewStringID("/////ID"), testData.GetTypeID()).(*baseIDs.PropertyID),
			DataID: testData.GetID().(*baseIDs.DataID),
		}, true},
		{"-ve", &MesaProperty{
			ID:     testPropertyID.(*baseIDs.PropertyID),
			DataID: baseData.NewNumberData(sdkTypes.NewInt(10)).GetID().(*baseIDs.DataID),
		}, true},
		{"-ve", &MesaProperty{
			ID:     testPropertyID.(*baseIDs.PropertyID),
			DataID: &baseIDs.DataID{&baseIDs.StringID{"S"}, &baseIDs.HashID{[]byte("a")}}},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err == nil && tt.want {
				t.Errorf("MesaPropertyValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("MesaPropertyValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}
