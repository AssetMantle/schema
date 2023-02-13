package utilities

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
)

func GetZeroValueDataFromID(typeID ids.StringID) data.Data {
	if zeroDataValue, err := ReadData(joinDataTypeAndValueStrings(typeID.AsString(), "")); err != nil {
		panic(err)
	} else {
		return zeroDataValue
	}
}
