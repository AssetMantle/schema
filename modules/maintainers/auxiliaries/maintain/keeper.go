// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"context"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	maintainerID := baseIDs.NewMaintainerID(constansts.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.IdentityID)),
		)))

	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(maintainerID))

	Mappable := maintainers.Get(key.NewKey(maintainerID))
	if Mappable == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}
	maintainer := mappable.GetMaintainer(Mappable)

	for _, maintainedProperty := range auxiliaryRequest.MaintainedMutables.GetMutablePropertyList().GetList() {
		if !maintainer.MaintainsProperty(maintainedProperty.GetID()) {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
