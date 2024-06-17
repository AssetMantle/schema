package data

type AnyListableData interface {
	ListableData
	Get() Data
	IsAnyListableData()
}
