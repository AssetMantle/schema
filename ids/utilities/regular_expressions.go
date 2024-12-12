// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import "regexp"

var IsValidStringID = regexp.MustCompile(`^[A-Za-z0-9_]{0,30}$`).MatchString
