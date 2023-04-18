package utilities

import (
	"strings"

	"github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
)

func ReadMetaProperty(metaPropertyString string) (properties.MetaProperty, error) {
	propertyIDString, dataString := SplitMetaProperty(metaPropertyString)
	if propertyIDString != "" {
		data, err := base.PrototypeAnyData().FromString(dataString)
		if err != nil {
			return nil, err
		}

		return baseProperties.NewMetaProperty(baseIDs.NewStringID(propertyIDString), data), nil
	}

	return nil, errorConstants.IncorrectFormat.Wrapf("propertyID missing from metaPropertyString %s", metaPropertyString)
}

func SplitMetaProperty(metaPropertyString string) (propertyIDString, dataString string) {
	if propertyIDAndData := strings.Split(metaPropertyString, constants.PropertyIDAndDataSeparator); len(propertyIDAndData) < 2 {
		return "", ""
	} else {
		return strings.TrimSpace(propertyIDAndData[0]), propertyIDAndData[1]
	}

}
