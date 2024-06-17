package documents

import "github.com/AssetMantle/schema/ids"

type PutOrder interface {
	Order
	GetPutOrderID() ids.OrderID
}
