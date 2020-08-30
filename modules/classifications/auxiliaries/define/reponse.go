/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryResponse struct {
	Success          bool     `json:"success"`
	Error            error    `json:"error"`
	ClassificationID types.ID `json:"classificationID"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}

func newAuxiliaryResponse(classificationID types.ID, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return auxiliaryResponse{
			Success:          false,
			Error:            error,
			ClassificationID: classificationID,
		}
	} else {
		return auxiliaryResponse{
			Success:          true,
			ClassificationID: classificationID,
		}
	}
}

func GetClassificationIDFromResponse(AuxiliaryResponse helpers.AuxiliaryResponse) (types.ID, error) {
	switch value := AuxiliaryResponse.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.ClassificationID, nil
		} else {
			return value.ClassificationID, value.GetError()
		}
	default:
		panic(errors.InvalidRequest)
	}
}
