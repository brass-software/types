package types

type Metadata struct {
	ID         string
	Name       string
	Doc        string
	Discussion Collection[*Comment]
}
