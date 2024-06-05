// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	typesSchema "github.com/AssetMantle/schema/types"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	data.RegisterLegacyAminoCodec(legacyAmino)
	baseData.RegisterLegacyAminoCodec(legacyAmino)

	documents.RegisterLegacyAminoCodec(legacyAmino)
	baseDocuments.RegisterLegacyAminoCodec(legacyAmino)

	ids.RegisterLegacyAminoCodec(legacyAmino)
	baseIDs.RegisterLegacyAminoCodec(legacyAmino)

	lists.RegisterLegacyAminoCodec(legacyAmino)
	baseLists.RegisterLegacyAminoCodec(legacyAmino)

	properties.RegisterLegacyAminoCodec(legacyAmino)
	baseProperties.RegisterLegacyAminoCodec(legacyAmino)

	qualified.RegisterLegacyAminoCodec(legacyAmino)
	baseQualified.RegisterLegacyAminoCodec(legacyAmino)

	typesSchema.RegisterLegacyAminoCodec(legacyAmino)
	baseTypes.RegisterLegacyAminoCodec(legacyAmino)
}
