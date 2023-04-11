// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package proto

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/schema/x/data"
	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/documents"
	baseDocuments "github.com/AssetMantle/schema/x/documents/base"
	"github.com/AssetMantle/schema/x/errors"
	baseErrors "github.com/AssetMantle/schema/x/errors/base"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/AssetMantle/schema/x/lists"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	"github.com/AssetMantle/schema/x/properties"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	"github.com/AssetMantle/schema/x/qualified"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	"github.com/AssetMantle/schema/x/traits"
	typesSchema "github.com/AssetMantle/schema/x/types"
	baseTypes "github.com/AssetMantle/schema/x/types/base"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {

	data.RegisterLegacyAminoCodec(legacyAmino)
	baseData.RegisterLegacyAminoCodec(legacyAmino)

	documents.RegisterLegacyAminoCodec(legacyAmino)
	baseDocuments.RegisterLegacyAminoCodec(legacyAmino)

	errors.RegisterLegacyAminoCodec(legacyAmino)
	baseErrors.RegisterLegacyAminoCodec(legacyAmino)

	ids.RegisterLegacyAminoCodec(legacyAmino)
	baseIDs.RegisterLegacyAminoCodec(legacyAmino)

	lists.RegisterLegacyAminoCodec(legacyAmino)
	baseLists.RegisterLegacyAminoCodec(legacyAmino)

	properties.RegisterLegacyAminoCodec(legacyAmino)
	baseProperties.RegisterLegacyAminoCodec(legacyAmino)

	qualified.RegisterLegacyAminoCodec(legacyAmino)
	baseQualified.RegisterLegacyAminoCodec(legacyAmino)

	traits.RegisterLegacyAminoCodec(legacyAmino)

	typesSchema.RegisterLegacyAminoCodec(legacyAmino)
	baseTypes.RegisterLegacyAminoCodec(legacyAmino)
}
