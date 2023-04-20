package utilities

import (
	dataUtilities "github.com/AssetMantle/schema/go/data/utilities"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/properties/utilities"
)

func ReadMetaPropertyList(metaPropertiesString string) (lists.PropertyList, error) {
	var Properties []properties.Property

	metaProperties := dataUtilities.SplitListString(metaPropertiesString)
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
