// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/schema/schema/data"
	baseData "github.com/AssetMantle/schema/schema/data/base"
	"github.com/AssetMantle/schema/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/schema/documents/base"
	"github.com/AssetMantle/schema/schema/errors"
	baseErrors "github.com/AssetMantle/schema/schema/errors/base"
	"github.com/AssetMantle/schema/schema/helpers"
	"github.com/AssetMantle/schema/schema/ids"
	baseIDs "github.com/AssetMantle/schema/schema/ids/base"
	"github.com/AssetMantle/schema/schema/lists"
	baseLists "github.com/AssetMantle/schema/schema/lists/base"
	"github.com/AssetMantle/schema/schema/properties"
	baseProperties "github.com/AssetMantle/schema/schema/properties/base"
	"github.com/AssetMantle/schema/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/schema/qualified/base"
	"github.com/AssetMantle/schema/schema/traits"
	typesSchema "github.com/AssetMantle/schema/schema/types"
	baseTypes "github.com/AssetMantle/schema/schema/types/base"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {

	data.RegisterLegacyAminoCodec(legacyAmino)
	baseData.RegisterLegacyAminoCodec(legacyAmino)

	documents.RegisterLegacyAminoCodec(legacyAmino)
	baseDocuments.RegisterLegacyAminoCodec(legacyAmino)

	errors.RegisterLegacyAminoCodec(legacyAmino)
	baseErrors.RegisterLegacyAminoCodec(legacyAmino)

	helpers.RegisterLegacyAminoCodec(legacyAmino)

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
