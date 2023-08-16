// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"sort"

	"github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
)

var _ lists.IDList = (*IDList)(nil)

func (idList *IDList) ValidateBasic() error {
	if len(idList.AnyIDs) > 0 {
		expectedIDType := idList.AnyIDs[0].GetTypeID()

		for _, id := range idList.AnyIDs {
			if err := id.ValidateBasic(); err != nil {
				return err
			}

			if id.GetTypeID().Compare(expectedIDType) != 0 {
				return constants.IncorrectFormat.Wrapf("id list contains unexpected id type %s", id.GetTypeID().AsString())
			}
		}
	}
	return nil
}
func (idList *IDList) GetList() []ids.AnyID {
	returnIDList := make([]ids.AnyID, len(idList.AnyIDs))

	for i, listable := range idList.AnyIDs {
		returnIDList[i] = listable
	}

	return returnIDList
}
func (idList *IDList) Search(id ids.ID) (index int, found bool) {
	updatedList := idList.sort().(*IDList)
	index = sort.Search(
		len(updatedList.AnyIDs),
		func(i int) bool {
			return updatedList.AnyIDs[i].Compare(id) >= 0
		},
	)

	if index < len(updatedList.AnyIDs) && updatedList.AnyIDs[index].Compare(id) == 0 {
		return index, true
	}

	return index, false
}
func (idList *IDList) Add(ids ...ids.ID) lists.IDList {
	updatedList := idList.sort().(*IDList)
	for _, listable := range ids {
		if index, found := updatedList.Search(listable); !found {
			updatedList.AnyIDs = append(updatedList.AnyIDs, listable.ToAnyID().(*baseIDs.AnyID))
			copy(updatedList.AnyIDs[index+1:], updatedList.AnyIDs[index:])
			updatedList.AnyIDs[index] = listable.ToAnyID().(*baseIDs.AnyID)
		}
	}
	return updatedList
}
func (idList *IDList) Remove(ids ...ids.ID) lists.IDList {
	updatedList := idList.sort().(*IDList)

	for _, listable := range ids {
		if index, found := updatedList.Search(listable); found {
			updatedList.AnyIDs = append(updatedList.AnyIDs[:index], updatedList.AnyIDs[index+1:]...)
		}
	}

	return updatedList
}
func (idList *IDList) sort() lists.IDList {
	sort.Slice(idList.AnyIDs, func(i, j int) bool {
		return idList.AnyIDs[i].Compare(idList.AnyIDs[j]) < 0
	})
	return idList
}

func NewIDList(ids ...ids.ID) lists.IDList {
	sort.Slice(ids, func(i, j int) bool {
		return ids[i].Compare(ids[j]) < 0
	})
	return (&IDList{}).Add(ids...)
}
