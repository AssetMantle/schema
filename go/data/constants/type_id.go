// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

// TODO use enum
var (
	AccAddressDataTypeID = baseIDs.NewStringID("A")
	AnyDataTypeID        = baseIDs.NewStringID("Y")
	BooleanDataTypeID    = baseIDs.NewStringID("B")
	DecDataTypeID        = baseIDs.NewStringID("D")
	HeightDataTypeID     = baseIDs.NewStringID("H")
	ListDataTypeID       = baseIDs.NewStringID("L")
	LinkedDataTypeID     = baseIDs.NewStringID("U")
	NumberDataTypeID     = baseIDs.NewStringID("N")
	StringDataTypeID     = baseIDs.NewStringID("S")
)
