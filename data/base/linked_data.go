// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"cosmossdk.io/math"
	"fmt"
	"github.com/AssetMantle/schema/data"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/data/utilities"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"net/url"
	"strings"
)

var _ data.LinkedData = (*LinkedData)(nil)

func (linkedData *LinkedData) ValidateBasic() error {
	if err := linkedData.GetResourceID().ValidateBasic(); err != nil {
		return err
	}

	if err := linkedData.GetExtensionID().ValidateBasic(); err != nil {
		return err
	}

	if linkedData.GetServiceEndpoint() != "" || linkedData.GetExtensionID().Get() != "" {
		if !dataConstants.IsValidFileExtension(linkedData.GetExtensionID().Get()) {
			return fmt.Errorf("invalid extension ID: %s", linkedData.GetExtensionID().Get())
		}

		if URL, err := url.Parse(linkedData.GetServiceEndpoint()); err != nil || URL.Scheme == "" || URL.Host == "" {
			if err != nil {
				return fmt.Errorf("invalid service endpoint: %s : %s", linkedData.GetServiceEndpoint(), err.Error())
			} else {
				return fmt.Errorf("invalid service endpoint: %s : scheme or host is missing", linkedData.GetServiceEndpoint())
			}
		}
	}

	return nil
}
func (linkedData *LinkedData) ToAnyListableData() data.AnyListableData {
	return &AnyListableData{
		Impl: &AnyListableData_LinkedData{
			LinkedData: linkedData,
		}}
}
func (linkedData *LinkedData) Compare(listableData data.ListableData) int {
	return bytes.Compare(linkedData.Bytes(), listableData.ToAnyListableData().Get().(*LinkedData).Bytes())
}
func (linkedData *LinkedData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(linkedData)
}
func (linkedData *LinkedData) GetBondWeight() math.Int {
	return dataConstants.LinkedDataWeight
}
func (linkedData *LinkedData) AsString() string {
	return utilities.JoinCompositeDataStrings(linkedData.GetResourceID().AsString(), linkedData.GetExtensionID().AsString(), linkedData.GetServiceEndpoint())
}
func (*LinkedData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeLinkedData(), nil
	}

	dataStringList := utilities.SplitCompositeDataString(dataString, 3)
	if len(dataStringList) != 3 {
		return PrototypeLinkedData(), fmt.Errorf("linked data is either missing resource ID, extension ID, or service endpoint: %s", dataString)
	}

	resourceID, err := baseIDs.PrototypeHashID().FromString(dataStringList[0])
	if err != nil {
		return PrototypeLinkedData(), fmt.Errorf("invalid resource ID: %s: %s", dataStringList[0], err.Error())
	}

	extensionID, err := baseIDs.PrototypeStringID().FromString(dataStringList[1])
	if err != nil {
		return PrototypeLinkedData(), fmt.Errorf("invalid extension ID: %s: %s", dataStringList[1], err.Error())
	}

	linkedData := NewLinkedData(resourceID.(ids.HashID), extensionID.(ids.StringID), dataStringList[2])
	if err := linkedData.ValidateBasic(); err != nil {
		return PrototypeLinkedData(), fmt.Errorf("invalid linked data: %s: %s", dataString, err.Error())
	}

	return linkedData, nil
}
func (linkedData *LinkedData) Bytes() []byte {
	return bytes.Join([][]byte{linkedData.GetResourceID().Bytes(), linkedData.GetExtensionID().Bytes(), []byte(linkedData.GetServiceEndpoint())}, dataConstants.ListBytesSeparator)
}
func (linkedData *LinkedData) GetTypeID() ids.StringID {
	return dataConstants.LinkedDataTypeID
}
func (linkedData *LinkedData) ZeroValue() data.Data {
	return PrototypeLinkedData()
}
func (linkedData *LinkedData) GenerateHashID() ids.HashID {
	if linkedData.Compare(linkedData.ZeroValue().(data.ListableData)) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(linkedData.Bytes())
}
func (linkedData *LinkedData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_LinkedData{
			LinkedData: linkedData,
		}}
}
func (linkedData *LinkedData) GetResourceID() ids.HashID {
	return linkedData.ResourceID
}
func (linkedData *LinkedData) GetExtensionID() ids.StringID {
	return linkedData.ExtensionID
}
func (linkedData *LinkedData) GetServiceEndpoint() string {
	return linkedData.ServiceEndpoint
}

func PrototypeLinkedData() data.LinkedData {
	return NewLinkedData(baseIDs.PrototypeHashID(), baseIDs.PrototypeStringID(), "")
}

func NewLinkedData(resourceID ids.HashID, extensionID ids.StringID, serviceEndpoint string) data.LinkedData {
	return &LinkedData{
		ResourceID:      resourceID.(*baseIDs.HashID),
		ExtensionID:     extensionID.(*baseIDs.StringID),
		ServiceEndpoint: serviceEndpoint,
	}
}
