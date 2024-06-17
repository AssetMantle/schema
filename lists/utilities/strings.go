package utilities

import (
	"strings"

	"github.com/AssetMantle/schema/lists/constants"
)

func SplitPropertyListString(propertyListString string) []string {
	return strings.Split(propertyListString, constants.PropertyListStringSeparator)
}
