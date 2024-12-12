// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import "github.com/AssetMantle/schema/ids"

type PutOrder interface {
	Order
	GetPutOrderID() ids.OrderID
}
