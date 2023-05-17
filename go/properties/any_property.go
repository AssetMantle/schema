package properties

type AnyProperty interface {
	IsAnyProperty()
	Get() Property
	Property
}
