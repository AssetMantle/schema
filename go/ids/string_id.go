package ids

type StringID interface {
	ID
	Get() string
}
