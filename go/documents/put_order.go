package documents

import "github.com/AssetMantle/schema/go/ids"

type PutOrder interface {
	Order
	GetPutOrderID() ids.OrderID
}
