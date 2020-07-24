package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	splits := mapper.NewSplits(transactionKeeper.mapper, context).Fetch(message.SplitID)
	split := splits.Get(message.SplitID)
	if split == nil {
		return constants.EntityNotFound
	}
	//if !split.CanBurn(types.NewHeight(context.BlockHeight())) {
	//	return constants.DeletionNotAllowed
	//}
	splits.Remove(split)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, _ []interface{}) helpers.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
