package types

type Collection[T any] struct {
	Items   []CollectionItem
	Indexes map[string]map[string]int
}
