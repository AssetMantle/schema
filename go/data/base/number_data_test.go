package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberData_ValidateBasic(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{"validate empty string", fields{""}, assert.Error},
		{"validate zero value", fields{PrototypeNumberData().ZeroValue().AsString()}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numberData := &NumberData{
				Value: tt.fields.Value,
			}
			tt.wantErr(t, numberData.ValidateBasic(), fmt.Sprintf("ValidateBasic()"))
		})
	}
}
