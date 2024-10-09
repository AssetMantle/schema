package data

import "cosmossdk.io/math"

type NumberData interface {
	ListableData
	Get() math.Int
}
