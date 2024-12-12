// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
)

func Test_new_Classification(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))

	testClassification := NewClassification(immutables, mutables)
	type fields struct {
		immutables qualified.Immutables
		mutables   qualified.Mutables
	}
	tests := []struct {
		name      string
		fields    fields
		want      documents.Classification
		wantError bool
	}{
		{"+ve", fields{immutables, mutables}, testClassification, false},
		{"nil immutables", fields{nil, mutables}, testClassification, true},
		{"nil mutables", fields{immutables, nil}, testClassification, true},
		{"nil", fields{nil, nil}, testClassification, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantError {
					t.Errorf("error = %v, wantErr %v", r, tt.wantError)
				}
			}()
			if got := NewClassification(tt.fields.immutables, tt.fields.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClassification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classifiaction_GetSupply(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithSupply := NewDocument(classificationID, immutables, baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(999))))))
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   math.Int
	}{
		{"+ve", fields{testDocument}, sdkTypes.ZeroInt()},
		{"+ve with bondAmount", fields{testDocumentWithSupply}, sdkTypes.NewInt(999)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			classification := NewClassificationFromDocument(tt.fields.Document)
			if got := classification.GetBondAmount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}
