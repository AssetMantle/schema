// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import "regexp"

var IsValidStringData = regexp.MustCompile(`^[\pL\pM\pZ\pS\pN\pP]{0,280}$`).MatchString
var IsValidFileExtension = regexp.MustCompile(`(?:jpg|jpeg|png|gif|bmp|svg|mp4|avi|mov|wmv|flv|mp3|wav|ogg|aac|pdf)$`).MatchString
