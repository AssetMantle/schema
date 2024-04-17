package base

import (
	"bytes"
	"fmt"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/constants"
	"github.com/AssetMantle/schema/go/qualified"
	"strings"
)

var _ ids.OrderID = (*OrderID)(nil)

func (orderID *OrderID) ValidateBasic() error {
	if orderID == nil {
		return fmt.Errorf("order ID is empty")
	}

	return orderID.HashID.ValidateBasic()
}
func (orderID *OrderID) GetTypeID() ids.StringID {
	return NewStringID(constants.OrderIDType)
}
func (orderID *OrderID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeOrderID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeOrderID(), err
	} else {
		orderID := &OrderID{HashID: hashID.(*HashID)}
		if err := orderID.ValidateBasic(); err != nil {
			return PrototypeOrderID(), err
		}

		return orderID, nil
	}
}
func (orderID *OrderID) AsString() string {
	return orderID.HashID.AsString()
}
func (orderID *OrderID) Bytes() []byte {
	return orderID.HashID.IDBytes
}
func (orderID *OrderID) IsOrderID() {}
func (orderID *OrderID) Compare(id ids.ID) int {
	return bytes.Compare(orderID.Bytes(), id.ToAnyID().Get().(*OrderID).Bytes())
}
func (orderID *OrderID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_OrderID{
			OrderID: orderID,
		},
	}
}

func NewOrderID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.OrderID {
	return &OrderID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeOrderID() ids.OrderID {
	return &OrderID{
		HashID: PrototypeHashID().(*HashID),
	}
}
