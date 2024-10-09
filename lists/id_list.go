// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/ids"
)

type IDList interface {
	GetList() []ids.AnyID
	Search(ids.ID) (index int, found bool)
	// Add ignores ids of different type
	Add(...ids.ID) IDList
	Remove(...ids.ID) IDList
	ValidateBasic() error
	ToListData() data.ListData
}
