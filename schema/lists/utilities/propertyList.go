package utilities

import (
	"github.com/AssetMantle/schema/schema/lists"
	"github.com/AssetMantle/schema/schema/lists/base"
	"github.com/AssetMantle/schema/schema/properties"
	"github.com/AssetMantle/schema/schema/properties/utilities"
	stringUtilities "github.com/AssetMantle/schema/utilities/string"
)

func ReadMetaPropertyList(metaPropertiesString string) (lists.PropertyList, error) {
	var Properties []properties.Property

	metaProperties := stringUtilities.SplitListString(metaPropertiesString)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, err := utilities.ReadMetaProperty(metaPropertyString)
			if err != nil {
				return nil, err
			}

			Properties = append(Properties, metaProperty)
		}
	}

	return base.NewPropertyList(Properties...), nil
}
