// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/ids"
)

type Maintainer interface {
	GetIdentityID() ids.IdentityID
	GetMaintainedClassificationID() ids.ClassificationID
	GetMaintainedProperties() data.ListData
	GetPermissions() data.ListData

	IsPermitted(permissionID ids.StringID) bool
	MaintainsProperty(ids.PropertyID) bool

	Document
}
