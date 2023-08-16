package utilities

import (
	"strings"

	"github.com/AssetMantle/schema/go/data/constants"
)

func JoinListStrings(listStrings ...string) string {
	return strings.Join(listStrings, constants.ListStringSeparator)
}

func SplitListString(listString string) []string {
	return strings.Split(listString, constants.ListStringSeparator)
}

func SplitCompositeDataString(listString string, parts int) []string {
	return strings.SplitN(listString, constants.CompositeDataStringSeparator, parts)
}
func JoinCompositeDataStrings(listStrings ...string) string {
	return strings.Join(listStrings, constants.CompositeDataStringSeparator)
}
