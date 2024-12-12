// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"strings"

	"github.com/AssetMantle/schema/lists/constants"
)

func SplitPropertyListString(propertyListString string) []string {
	return strings.Split(propertyListString, constants.PropertyListStringSeparator)
}
