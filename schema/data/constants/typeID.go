// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	baseIDs "github.com/AssetMantle/schema/schema/ids/base"
)

// TODO use enum
var (
	AccAddressDataTypeID = baseIDs.NewStringID("A")
	BooleanDataTypeID    = baseIDs.NewStringID("B")
	DecDataTypeID        = baseIDs.NewStringID("D")
	HeightDataTypeID     = baseIDs.NewStringID("H")
	IDDataTypeID         = baseIDs.NewStringID("I")
	ListDataTypeID       = baseIDs.NewStringID("L")
	NumberDataTypeID     = baseIDs.NewStringID("N")
	StringDataTypeID     = baseIDs.NewStringID("S")
)
