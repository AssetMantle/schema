// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/schema/ids"
)

type LinkedData interface {
	ListableData
	GetResourceID() ids.HashID
	GetExtensionID() ids.StringID
	GetServiceEndpoint() string
}
