// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/ids"
	baseProperties "github.com/AssetMantle/schema/properties/base"
)

func TestNewMaintainer(t *testing.T) {
	classificationID, _, _, testDocument := createTestInput()
	testIdentity := NewIdentityFromDocument(testDocument)
	identityID := baseIDs.NewIdentityID(testIdentity.GetClassificationID(), testIdentity.GetImmutables())
	maintainerProperties := baseLists.NewIDList(baseIDs.NewStringID("a"), baseIDs.NewStringID("b"))
	permissions := baseLists.NewIDList(baseIDs.NewStringID("c"), baseIDs.NewStringID("d"))
	maintainer := NewDocument(maintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(identityID)),
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(classificationID)),
		)),
		baseQualified.NewMutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedPropertiesProperty.GetKey(), maintainerProperties.ToListData()),
			baseProperties.NewMetaProperty(constantProperties.PermissionsProperty.GetKey(), permissions.ToListData()),
		)),
	)
	type args struct {
		identityID           ids.IdentityID
		classificationID     ids.ClassificationID
		maintainerProperties lists.IDList
		permissions          lists.IDList
	}
	tests := []struct {
		name string
		args args
		want documents.Maintainer
	}{
		{"+ve", args{identityID, classificationID, maintainerProperties, permissions}, NewMaintainerFromDocument(maintainer)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaintainer(tt.args.identityID, tt.args.classificationID, tt.args.maintainerProperties, tt.args.permissions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
