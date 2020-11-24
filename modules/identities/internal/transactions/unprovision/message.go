/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unprovision

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/common"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type message struct {
	From       sdkTypes.AccAddress `json:"from" valid:"required~required field from missing matches(^commit[a-z0-9]{39}$)~invalid field from"`
	To         sdkTypes.AccAddress `json:"to" valid:"required~required field to missing matches(^commit[a-z0-9]{39}$)~invalid field to"`
	IdentityID types.ID            `json:"identityID" valid:"required~required field identityID missing"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(common.Codec.MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, message{})
}
func messageFromInterface(msg sdkTypes.Msg) message {
	switch value := msg.(type) {
	case message:
		return value
	default:
		return message{}
	}
}
func messagePrototype() helpers.Message {
	return message{}
}
func newMessage(from sdkTypes.AccAddress, to sdkTypes.AccAddress, identityID types.ID) sdkTypes.Msg {
	return message{
		From:       from,
		To:         to,
		IdentityID: identityID,
	}
}