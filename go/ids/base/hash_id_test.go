package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/ids"
)

func TestReadHashID(t *testing.T) {
	type args struct {
		hashIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.HashID
		wantErr bool
	}{
		{name: "empty string", args: args{""}, want: &HashID{[]uint8{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeHashID().FromString(tt.args.hashIDString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadHashID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashID_AsString(t *testing.T) {
	type fields struct {
		IDBytes []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{GenerateHashID(NewStringID("sdfsdfsdf").Bytes()).Bytes()}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashID := &HashID{
				IDBytes: tt.fields.IDBytes,
			}
			if got := hashID.AsString(); got != tt.want {
				t.Errorf("AsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
