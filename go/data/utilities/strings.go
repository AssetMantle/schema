package utilities

import (
	"strings"

	"github.com/AssetMantle/schema/go/data/constants"
)

func JoinListDataStrings(listDataStrings ...string) string {
	return strings.Join(listDataStrings, constants.ListDataStringSeparator)
}

func SplitListDataString(listDataString string) []string {
	return strings.Split(listDataString, constants.ListDataStringSeparator)
}

func SplitCompositeDataString(listString string, parts int) []string {
	return strings.SplitN(listString, constants.CompositeDataStringSeparator, parts)
}
func JoinCompositeDataStrings(listStrings ...string) string {
	return strings.Join(listStrings, constants.CompositeDataStringSeparator)
}
