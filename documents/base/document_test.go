// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
)

func ValidatedParameter[V *baseIDs.ClassificationID | *baseQualified.Immutables | *baseQualified.Mutables](value any) V {
	if value == nil {
		return nil
	}
	return value.(V)
}

func createTestInput() (ids.ClassificationID, qualified.Immutables, qualified.Mutables, documents.Document) {
	testImmutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	testMutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData().ZeroValue()), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiry"), baseData.NewHeightData(baseTypes.NewHeight(10))), baseProperties.NewMetaProperty(baseIDs.NewStringID("expiryHeight"), baseData.NewHeightData(baseTypes.NewHeight(100)))))
	classificationID := baseIDs.NewClassificationID(testImmutables, testMutables)
	testDocument := NewDocument(classificationID, testImmutables, testMutables)
	return classificationID, testImmutables, testMutables, testDocument
}

func TestNewDocument(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name    string
		args    args
		want    documents.Document
		wantErr bool
	}{
		{"+ve", args{classificationID: classificationID, immutables: testImmutables, mutables: testMutables}, &Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, false},
		{"+ve prototype classificationID", args{baseIDs.PrototypeClassificationID(), testImmutables, testMutables}, &Document{ClassificationID: baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, false},
		{"+ve empty immutables", args{classificationID, baseQualified.NewImmutables(nil), testMutables}, &Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: baseQualified.NewImmutables(nil).(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, false},
		{"+ve empty mutables", args{classificationID, testImmutables, baseQualified.NewMutables(nil)}, &Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: baseQualified.NewMutables(nil).(*baseQualified.Mutables)}, false},
		{"+ve all empty", args{baseIDs.PrototypeClassificationID(), baseQualified.NewImmutables(nil), baseQualified.NewMutables(nil)}, &Document{ClassificationID: baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID), Immutables: baseQualified.NewImmutables(nil).(*baseQualified.Immutables), Mutables: baseQualified.NewMutables(nil).(*baseQualified.Mutables)}, false},
		{"panic with nil immutables", args{classificationID: classificationID, immutables: nil, mutables: testMutables}, &Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: nil, Mutables: testMutables.(*baseQualified.Mutables)}, true},
		{"panic with nil mutables", args{classificationID: classificationID, immutables: testImmutables, mutables: nil}, &Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: nil}, true},
		{"panic with all nil", args{}, &Document{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := NewDocument(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetClassificationID(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	tests := []struct {
		name    string
		fields  fields
		want    ids.ClassificationID
		wantErr bool
	}{
		{"+ve", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, classificationID, false},
		{"+ve empty classificationID", fields{ClassificationID: baseIDs.PrototypeClassificationID(), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, baseIDs.PrototypeClassificationID(), false},
		{"+ve empty immutables", fields{classificationID, baseQualified.NewImmutables(nil), testMutables}, classificationID, false},
		{"+ve empty mutables", fields{classificationID, testImmutables, baseQualified.NewMutables(nil)}, classificationID, false},
		{"+ve all empty", fields{baseIDs.PrototypeClassificationID(), baseQualified.NewImmutables(nil), baseQualified.NewMutables(nil)}, baseIDs.PrototypeClassificationID(), false},
		{"panic with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, nil, true},
		{"panic with nil immutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: nil, Mutables: testMutables.(*baseQualified.Mutables)}, classificationID, true},
		{"panic with nil mutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: nil}, classificationID, true},
		{"panic with all nil", fields{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := NewDocument(tt.fields.ClassificationID, tt.fields.Immutables, tt.fields.Mutables).GetClassificationID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClassificationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetImmutables(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   qualified.Immutables
	}{
		{"+ve", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, testImmutables},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, testImmutables},
		{"+ve with nil immutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: nil, Mutables: testMutables.(*baseQualified.Mutables)}, (*baseQualified.Immutables)(nil)},
		{"+ve with nil mutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: nil}, testImmutables},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := &Document{
				ClassificationID: ValidatedParameter[*baseIDs.ClassificationID](tt.fields.ClassificationID),
				Immutables:       ValidatedParameter[*baseQualified.Immutables](tt.fields.Immutables),
				Mutables:         ValidatedParameter[*baseQualified.Mutables](tt.fields.Mutables),
			}
			if got := document.GetImmutables(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImmutables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetMutables(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   qualified.Mutables
	}{
		{"+ve", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, testMutables},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, testMutables},
		{"+ve with nil immutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: nil, Mutables: testMutables.(*baseQualified.Mutables)}, testMutables},
		{"+ve with nil mutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: nil}, (*baseQualified.Mutables)(nil)},
		{"+ve with all nil", fields{}, (*baseQualified.Mutables)(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := &Document{
				ClassificationID: ValidatedParameter[*baseIDs.ClassificationID](tt.fields.ClassificationID),
				Immutables:       ValidatedParameter[*baseQualified.Immutables](tt.fields.Immutables),
				Mutables:         ValidatedParameter[*baseQualified.Mutables](tt.fields.Mutables),
			}
			if got := document.GetMutables(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMutables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetProperty(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	testImmutablePropertyID := baseIDs.NewPropertyID(baseIDs.NewStringID("ID1"), baseIDs.NewStringID("ImmutableData"))
	testMutablePropertyID := baseIDs.NewPropertyID(baseIDs.NewStringID("ID2"), baseIDs.NewStringID("MutableData"))

	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      properties.Property
		wantPanic bool
	}{
		// TODO: Update unit test after issue # fix
		{"+ve for Immutable Property", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID), false},
		{"+ve with nil classificationID for Immutable Property", fields{ClassificationID: nil, Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID), false},
		{"+ve with nil immutables for Immutable Property", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: baseQualified.NewImmutables(baseLists.NewPropertyList()), Mutables: testMutables.(*baseQualified.Mutables)}, args{testImmutablePropertyID}, nil, false}, // TODO: panics for empty immutable struct
		{"+ve with nil mutables for Immutable Property", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: baseQualified.NewMutables(baseLists.NewPropertyList())}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID), false},
		{"+ve with all nil", fields{nil, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList())}, args{testImmutablePropertyID}, nil, false}, // TODO: panics for empty immutable struct
		{"+ve for Mutable Property", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID), false},
		{"+ve with nil classificationID for Mutable Property", fields{ClassificationID: nil, Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID), false},
		{"+ve with nil immutables for Mutable Property", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: baseQualified.NewImmutables(baseLists.NewPropertyList()), Mutables: testMutables.(*baseQualified.Mutables)}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID), false}, // TODO: panics for empty immutable struct
		{"+ve with nil mutables for Mutable Property", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: baseQualified.NewMutables(baseLists.NewPropertyList())}, args{testMutablePropertyID}, nil, false},
		// Panics for empty structs
		// GetProperty() searches by traversing through Immutables and Mutables of a document, hence setting them as nil should cause a fatal panic
		{"panic case nil immutables", fields{classificationID, nil, testMutables}, args{testImmutablePropertyID}, nil, true},
		{"panic case nil mutables", fields{classificationID, testImmutables, nil}, args{testImmutablePropertyID}, nil, true},
		{"panic case nil document", fields{nil, nil, nil}, args{testImmutablePropertyID}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := &Document{
				ClassificationID: ValidatedParameter[*baseIDs.ClassificationID](tt.fields.ClassificationID),
				Immutables:       ValidatedParameter[*baseQualified.Immutables](tt.fields.Immutables),
				Mutables:         ValidatedParameter[*baseQualified.Mutables](tt.fields.Mutables),
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					document.GetProperty(tt.args.propertyID)
				})
			} else if got := document.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_Mutate(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	testMutateProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("MutableData1"))
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	type args struct {
		propertyList []properties.Property
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      documents.Document
		wantPanic bool
	}{
		// TODO: Update after #59 fix https://github.com/AssetMantle/schema/issues/59
		{"+ve with no mutation", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, args{}, &Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, false},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}, args{[]properties.Property{testMutateProperty}}, (&Document{ClassificationID: nil, Immutables: testImmutables.(*baseQualified.Immutables), Mutables: testMutables.(*baseQualified.Mutables)}).Mutate(testMutateProperty), false},
		{"+ve with nil immutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: nil, Mutables: testMutables.(*baseQualified.Mutables)}, args{[]properties.Property{testMutateProperty}}, (&Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: nil, Mutables: testMutables.(*baseQualified.Mutables)}).Mutate(testMutateProperty), false},
		{"+ve with nil mutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: baseQualified.NewMutables(nil)}, args{[]properties.Property{testMutateProperty}}, (&Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: baseQualified.NewMutables(baseLists.NewPropertyList()).(*baseQualified.Mutables)}).Mutate(testMutateProperty), false}, // TODO: fix this
		{"+ve with all nil", fields{nil, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList())}, args{[]properties.Property{testMutateProperty}}, (&Document{nil, baseQualified.NewImmutables(baseLists.NewPropertyList()).(*baseQualified.Immutables), baseQualified.NewMutables(baseLists.NewPropertyList()).(*baseQualified.Mutables)}).Mutate(testMutateProperty), false},                                                                                        // TODO: fix this
		{"panic case nil document", fields{nil, nil, nil}, args{[]properties.Property{testMutateProperty}}, nil, true},
		{"panic case nil mutables", fields{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: nil}, args{[]properties.Property{testMutateProperty}}, (&Document{ClassificationID: classificationID.(*baseIDs.ClassificationID), Immutables: testImmutables.(*baseQualified.Immutables), Mutables: baseQualified.NewMutables(baseLists.NewPropertyList()).(*baseQualified.Mutables)}).Mutate(testMutateProperty), true}, // TODO: fix this
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := &Document{
				ClassificationID: ValidatedParameter[*baseIDs.ClassificationID](tt.fields.ClassificationID),
				Immutables:       ValidatedParameter[*baseQualified.Immutables](tt.fields.Immutables),
				Mutables:         ValidatedParameter[*baseQualified.Mutables](tt.fields.Mutables),
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					document.Mutate(tt.args.propertyList...)
				})
			} else if got := document.Mutate(tt.args.propertyList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_GenerateHashID(t *testing.T) {
	zeroStringDataProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("test"), baseData.NewStringData("").ZeroValue()).ToAnyProperty().(*baseProperties.AnyProperty)
	zeroNumberDataProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("test"), baseData.NewNumberData(types.ZeroInt()).ZeroValue()).ToAnyProperty().(*baseProperties.AnyProperty)

	sameKeyDifferentDataImmutables := baseQualified.NewImmutables(baseLists.NewPropertyList(zeroStringDataProperty, zeroNumberDataProperty))
	sameKeyAndDataImmutables := baseQualified.NewImmutables(baseLists.NewPropertyList(zeroStringDataProperty, zeroStringDataProperty))

	emptyMutables := baseQualified.NewMutables(baseLists.NewPropertyList())

	sameKeyDifferentDataDocument := NewDocument(baseIDs.NewClassificationID(sameKeyDifferentDataImmutables, emptyMutables), sameKeyDifferentDataImmutables, emptyMutables)
	sameKeyAndDataDocument := NewDocument(baseIDs.NewClassificationID(sameKeyAndDataImmutables, emptyMutables), sameKeyAndDataImmutables, emptyMutables)
	emptyDocument := NewDocument(&baseIDs.ClassificationID{HashID: &baseIDs.HashID{}}, &baseQualified.Immutables{PropertyList: &baseLists.PropertyList{}}, emptyMutables)

	assert.NotEqual(t, sameKeyDifferentDataDocument.GenerateHashID(), sameKeyAndDataDocument.GenerateHashID(), "hash of same key and different data should not be same")

	tests := []struct {
		name   string
		fields documents.Document
		want   ids.HashID
	}{
		{name: "generate hash of empty classification, immutables and mutables", fields: emptyDocument, want: &baseIDs.HashID{[]byte{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.GenerateHashID()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHashID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DocumentValidateBasic(t *testing.T) {
	zeroStringDataProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("test"), baseData.NewStringData("").ZeroValue()).ToAnyProperty().(*baseProperties.AnyProperty)
	zeroNumberDataProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("test"), baseData.NewNumberData(types.ZeroInt()).ZeroValue()).ToAnyProperty().(*baseProperties.AnyProperty)

	immutables1 := baseQualified.NewImmutables(baseLists.NewPropertyList(zeroStringDataProperty, zeroNumberDataProperty))
	immutables2 := baseQualified.NewImmutables(baseLists.NewPropertyList(zeroStringDataProperty, zeroStringDataProperty))

	mutables1 := baseQualified.NewMutables(baseLists.NewPropertyList(zeroStringDataProperty))
	mutables2 := baseQualified.NewMutables(baseLists.NewPropertyList(zeroNumberDataProperty))
	mutables3 := baseQualified.NewMutables(baseLists.NewPropertyList())

	document1 := NewDocument(baseIDs.NewClassificationID(immutables1, mutables1), immutables1, mutables1)
	document2 := NewDocument(baseIDs.NewClassificationID(immutables2, mutables2), immutables2, mutables2)
	document3 := NewDocument(baseIDs.NewClassificationID(immutables1, mutables3), immutables1, mutables3)
	document4 := &Document{&baseIDs.ClassificationID{&baseIDs.HashID{[]byte("a")}}, immutables1.(*baseQualified.Immutables), mutables3.(*baseQualified.Mutables)}

	tests := []struct {
		name string
		args documents.Document
		want bool
	}{
		{"-ve 1", document1, false},
		{"-ve 2", document2, false},
		{"-ve 3", document4, true},
		{"+ve 4", document3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.ValidateBasic()
			if err != nil {
				fmt.Println(err.Error())
			}
			if err == nil && tt.want {
				t.Errorf("DocumentValidateBasic() = %v, want %v", err, tt.want)
			}
			if err != nil && !tt.want {
				t.Errorf("DocumentValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}
