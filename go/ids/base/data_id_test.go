// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/ids"
	stringUtilities "github.com/AssetMantle/schema/go/ids/utilities"
)

func Test_dataIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      *DataID
		wantError bool
	}{
		{"+ve", args{&DataID{NewStringID("B").(*StringID), GenerateHashID([]byte("B")).(*HashID)}}, &DataID{NewStringID("B").(*StringID), GenerateHashID().(*HashID)}, false},
		{"-ve", args{&DataID{}}, &DataID{}, false},
		{"-ve", args{nil}, &DataID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantError {
					t.Errorf("MetadataFromInterface() error = %v, wantError %v", r, tt.wantError)
				}
			}()
			got := dataIDFromInterface(tt.args.i)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataIDFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_Bytes(t *testing.T) {
	type fields struct {
		Type   ids.StringID
		HashID ids.HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{NewStringID("B"), GenerateHashID()}, append(append([]byte{}, NewStringID("B").Bytes()...), GenerateHashID().Bytes()...)},
		{"+ve", fields{NewStringID("B"), GenerateHashID()}, append(append([]byte{}, NewStringID("B").Bytes()...), GenerateHashID().Bytes()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type.(*StringID),
				HashID: tt.fields.HashID.(*HashID),
			}
			if got := dataID.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_Compare(t *testing.T) {
	type fields struct {
		Type   ids.StringID
		HashID ids.HashID
	}
	type args struct {
		id ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"+ve", fields{NewStringID("B"), GenerateHashID()}, args{&DataID{NewStringID("B").(*StringID), GenerateHashID().(*HashID)}}, 0},
		{"+ve", fields{NewStringID("B"), GenerateHashID()}, args{&DataID{NewStringID("B").(*StringID), GenerateHashID().(*HashID)}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type.(*StringID),
				HashID: tt.fields.HashID.(*HashID),
			}
			if got := dataID.Compare(tt.args.id); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_GetHashID(t *testing.T) {
	type fields struct {
		Type   *StringID
		HashID *HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		{"+ve", fields{}, (&DataID{}).HashID},
		{"+ve", fields{NewStringID("B").(*StringID), GenerateHashID().(*HashID)}, GenerateHashID()},
		{"+ve", fields{NewStringID("B").(*StringID), GenerateHashID().(*HashID)}, GenerateHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type,
				HashID: tt.fields.HashID,
			}
			if got := dataID.GetHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_String(t *testing.T) {
	type fields struct {
		Type   *StringID
		HashID *HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{NewStringID("B").(*StringID), GenerateHashID().(*HashID)}, stringUtilities.JoinIDStrings(NewStringID("B").AsString(), GenerateHashID().AsString())},
		{"+ve", fields{NewStringID("B").(*StringID), GenerateHashID().(*HashID)}, stringUtilities.JoinIDStrings(NewStringID("B").AsString(), GenerateHashID().AsString())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type,
				HashID: tt.fields.HashID,
			}
			if got := dataID.AsString(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDataID(t *testing.T) {
	type args struct {
		dataIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.DataID
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", args{stringUtilities.JoinIDStrings(NewStringID("B").AsString(), GenerateHashID().AsString())}, &DataID{}, false},
		{"+ve with empty string", args{""}, PrototypeDataID(), false},
		{"+ve with nil", args{}, PrototypeDataID(), false},
		{"-ve", args{stringUtilities.JoinIDStrings(NewStringID("j").AsString(), "0")}, &DataID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeDataID().FromString(tt.args.dataIDString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDataID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDataID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
