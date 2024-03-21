package types

type Collection[T any] struct {
	Items   map[string]T
	Indexes map[string]map[string]string
}
