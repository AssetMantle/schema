package base

import (
	"github.com/AssetMantle/schema/go/ids/constants"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/ids"
)

func Test_HashIDFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    ids.HashID
		wantErr bool
	}{
		{"+ve", testValidBase64String, &HashID{testBytes}, false},
		{"-ve", "asd", &HashID{[]byte{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeHashID().FromString(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashIDFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashIDFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HashIDType(t *testing.T) {
	tests := []struct {
		name string
		args ids.HashID
		want ids.StringID
	}{
		{"+ve", &HashID{testBytes}, NewStringID(constants.HashIDType)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.GetTypeID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashIDType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HashIDAsString(t *testing.T) {
	tests := []struct {
		name string
		args ids.HashID
		want string
	}{
		{"+ve", &HashID{testBytes}, testValidBase64URLString},
		{"+ve", &HashID{[]byte{}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.AsString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashIDAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HashIDBytes(t *testing.T) {
	tests := []struct {
		name string
		args ids.HashID
		want []byte
	}{
		{"+ve", &HashID{testBytes}, testBytes},
		{"+ve", &HashID{[]byte{}}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashIDByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HashIDCompare(t *testing.T) {
	hashID := &HashID{testBytes}
	tests := []struct {
		name string
		args ids.HashID
		want bool
	}{
		{"+ve", &HashID{testBytes}, true},
		{"-ve", &HashID{[]byte("test")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Compare(hashID)
			if !reflect.DeepEqual(got == 0, tt.want) {
				t.Errorf("HashIDCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
