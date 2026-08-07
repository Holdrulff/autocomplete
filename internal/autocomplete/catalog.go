package autocomplete

type Catalog struct {
	suggestions []Suggestion
}

func NewCatalog(suggestions []Suggestion) Catalog {
	copied := make([]Suggestion, len(suggestions))
	copy(copied, suggestions)

	return Catalog{
		suggestions: copied,
	}
}

func (c Catalog) Len() int {
	return len(c.suggestions)
}
