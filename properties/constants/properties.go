// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"cosmossdk.io/math"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/properties/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

var (
	AuthenticationProperty             = base.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())
	BondAmountProperty                 = base.NewMetaProperty(baseIDs.NewStringID("bondAmount"), baseData.NewNumberData(math.ZeroInt()))
	BondRateProperty                   = base.NewMetaProperty(baseIDs.NewStringID("bondRate"), baseData.NewNumberData(math.ZeroInt()))
	BurnEnabledProperty                = base.NewMetaProperty(baseIDs.NewStringID("burnEnabled"), baseData.NewBooleanData(false))
	BurnHeightProperty                 = base.NewMetaProperty(baseIDs.NewStringID("burnHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	DenomProperty                      = base.NewMetaProperty(baseIDs.NewStringID("denom"), baseData.NewStringData(""))
	DefineEnabledProperty              = base.NewMetaProperty(baseIDs.NewStringID("defineEnabled"), baseData.NewBooleanData(false))
	DeputizeAllowedProperty            = base.NewMetaProperty(baseIDs.NewStringID("deputizeAllowed"), baseData.NewBooleanData(false))
	CreationHeightProperty             = base.NewMetaProperty(baseIDs.NewStringID("creationHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	ExchangeRateProperty               = base.NewMetaProperty(baseIDs.NewStringID("exchangeRate"), baseData.NewDecData(math.LegacySmallestDec()).ZeroValue())
	ExpiryHeightProperty               = base.NewMetaProperty(baseIDs.NewStringID("expiryHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	ExecutionHeightProperty            = base.NewMetaProperty(baseIDs.NewStringID("executionHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	LockHeightProperty                 = base.NewMetaProperty(baseIDs.NewStringID("lockHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	IdentityIDProperty                 = base.NewMetaProperty(baseIDs.NewStringID("identityID"), baseData.NewIDData(baseIDs.PrototypeIdentityID()))
	IssueEnabledProperty               = base.NewMetaProperty(baseIDs.NewStringID("issueEnabled"), baseData.NewBooleanData(false))
	MaintainedClassificationIDProperty = base.NewMetaProperty(baseIDs.NewStringID("maintainedClassificationID"), baseData.NewIDData(baseIDs.PrototypeClassificationID()))
	MaintainedPropertiesProperty       = base.NewMetaProperty(baseIDs.NewStringID("maintainedProperties"), baseData.NewListData())
	MakerIDProperty                    = base.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(baseIDs.PrototypeIdentityID()))
	MakerAssetIDProperty               = base.NewMetaProperty(baseIDs.NewStringID("makerAssetID"), baseData.NewIDData(baseIDs.PrototypeAssetID()))
	MakerSplitProperty                 = base.NewMetaProperty(baseIDs.NewStringID("makerSplit"), baseData.NewNumberData(math.ZeroInt()))
	MaxPropertyCountProperty           = base.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(math.ZeroInt()))
	MaxProvisionAddressCountProperty   = base.NewMetaProperty(baseIDs.NewStringID("maxProvisionAddressCount"), baseData.NewNumberData(math.ZeroInt()))
	MaxOrderLifeProperty               = base.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MintEnabledProperty                = base.NewMetaProperty(baseIDs.NewStringID("mintEnabled"), baseData.NewBooleanData(false))
	ModuleNameProperty                 = base.NewMetaProperty(baseIDs.NewStringID("moduleName"), baseData.NewIDData(baseIDs.PrototypeStringID()))
	NameProperty                       = base.NewMetaProperty(baseIDs.NewStringID("name"), baseData.NewIDData(baseIDs.PrototypeStringID()))
	PermissionsProperty                = base.NewMetaProperty(baseIDs.NewStringID("permissions"), baseData.NewListData())
	PutEnabledProperty                 = base.NewMetaProperty(baseIDs.NewStringID("putEnabled"), baseData.NewBooleanData(false))
	RenumerateEnabledProperty          = base.NewMetaProperty(baseIDs.NewStringID("renumerateEnabled"), baseData.NewBooleanData(false))
	RevealEnabledProperty              = base.NewMetaProperty(baseIDs.NewStringID("revealEnabled"), baseData.NewBooleanData(false))
	TakerIDProperty                    = base.NewMetaProperty(baseIDs.NewStringID("takerID"), baseData.NewIDData(baseIDs.PrototypeIdentityID()))
	TakerAssetIDProperty               = base.NewMetaProperty(baseIDs.NewStringID("takerAssetID"), baseData.NewIDData(baseIDs.PrototypeAssetID()))
	TakerSplitProperty                 = base.NewMetaProperty(baseIDs.NewStringID("takerSplit"), baseData.NewNumberData(math.ZeroInt()))
	TransferEnabledProperty            = base.NewMetaProperty(baseIDs.NewStringID("transferEnabled"), baseData.NewBooleanData(false))
	SupplyProperty                     = base.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewNumberData(math.OneInt()))
	WrapAllowedCoinsProperty           = base.NewMetaProperty(baseIDs.NewStringID("wrapAllowedCoins"), baseData.NewListData())
	UnwrapAllowedCoinsProperty         = base.NewMetaProperty(baseIDs.NewStringID("unwrapAllowedCoins"), baseData.NewListData())
	QuashEnabledProperty               = base.NewMetaProperty(baseIDs.NewStringID("quashEnabled"), baseData.NewBooleanData(false))

	// RWA Compliance Properties
	ComplianceTierProperty             = base.NewMetaProperty(baseIDs.NewStringID("complianceTier"), baseData.NewNumberData(math.ZeroInt()))              // 0=anonymous, 1=basic KYC, 2=accredited, 3=institutional, 4=custodian
	JurisdictionProperty               = base.NewMetaProperty(baseIDs.NewStringID("jurisdiction"), baseData.NewStringData(""))                             // ISO-3166 country code
	AccreditationExpiryProperty        = base.NewMetaProperty(baseIDs.NewStringID("accreditationExpiry"), baseData.NewHeightData(baseTypes.NewHeight(-1))) // block height when accreditation expires
	SanctionsClearedProperty           = base.NewMetaProperty(baseIDs.NewStringID("sanctionsCleared"), baseData.NewBooleanData(false))                     // whether sanctions screening passed
	SanctionsClearedHeightProperty     = base.NewMetaProperty(baseIDs.NewStringID("sanctionsClearedHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1))) // block height of last sanctions check
	VerificationProviderProperty       = base.NewMetaProperty(baseIDs.NewStringID("verificationProvider"), baseData.NewIDData(baseIDs.PrototypeIdentityID())) // identity of KYC provider
)
