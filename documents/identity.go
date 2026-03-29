// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Identity interface {

	// TODO removal for expired identity
	// GetExpiry returns the expiry property of an Identity
	// * If the property is not found, it returns a default value and not nil
	GetExpiry() types.Height

	// GetAuthentication returns the authentication property of an Identity
	// * If the property is not found, it returns a default value and not nil
	GetAuthentication() data.ListData

	IsProvisioned(sdkTypes.AccAddress) bool
	GetProvisionedAddressCount() math.Int
	ProvisionAddress(...sdkTypes.AccAddress) Identity
	UnprovisionAddress(...sdkTypes.AccAddress) Identity

	// RWA Compliance methods
	// GetComplianceTier returns the compliance tier (0=anonymous, 1=basic, 2=accredited, 3=institutional, 4=custodian)
	GetComplianceTier() math.Int
	// GetJurisdiction returns the ISO-3166 jurisdiction code
	GetJurisdiction() string
	// GetAccreditationExpiry returns the block height at which accreditation expires
	GetAccreditationExpiry() types.Height
	// IsSanctionsCleared returns whether this identity has passed sanctions screening
	IsSanctionsCleared() bool

	Document
}
