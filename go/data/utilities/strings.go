package utilities

import (
	"strings"

	"github.com/AssetMantle/schema/go/data/constants"
)

// TODO write testcase for empty and singular input
func JoinListStrings(listStrings ...string) string {
	return strings.Join(listStrings, constants.ListStringSeparator)
}

func SplitListString(listString string) []string {
	return strings.Split(listString, constants.ListStringSeparator)
}

func SplitNListString(listString string, n int) []string {
	return strings.SplitN(listString, constants.ListStringSeparator, n)
}
