package data

type ListableData interface {
	ToAnyListableData() AnyListableData
	Compare(ListableData) int
	Data
}
