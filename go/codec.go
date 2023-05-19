// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/errors"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/qualified"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	typesSchema "github.com/AssetMantle/schema/go/types"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {

	data.RegisterLegacyAminoCodec(legacyAmino)
	baseData.RegisterLegacyAminoCodec(legacyAmino)

	documents.RegisterLegacyAminoCodec(legacyAmino)
	baseDocuments.RegisterLegacyAminoCodec(legacyAmino)

	errors.RegisterLegacyAminoCodec(legacyAmino)

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
