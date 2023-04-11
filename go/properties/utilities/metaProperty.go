package utilities

import (
	"strings"

	"github.com/AssetMantle/schema/x/data/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/AssetMantle/schema/x/properties"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	"github.com/AssetMantle/schema/x/properties/constants"
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
