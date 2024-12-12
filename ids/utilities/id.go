// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"strings"
)

const IDSeparator = "."

// TODO write testcase for empty and singular input
func JoinIDStrings(idStrings ...string) string {
	return strings.Join(idStrings, IDSeparator)
}

func SplitCompositeIDString(idString string) []string {
	return strings.SplitN(idString, IDSeparator, 2)
}
