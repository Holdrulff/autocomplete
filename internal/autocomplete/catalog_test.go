package autocomplete

import "testing"

func TestCatalog(t *testing.T) {
	catalog := NewCatalog([]Suggestion{
		{Value: "go", Score: 100},
		{Value: "react", Score: 90},
	})

	if got, want := catalog.Len(), 2; got != want {
		t.Errorf("Catalog.Len() = %d, want %d", got, want)
	}
}

func TestNewCatalogCopiesSuggestions(t *testing.T) {
	source := []Suggestion{
		{Value: "go", Score: 100},
	}

	catalog := NewCatalog(source)
	source[0].Value = "changed"

	if got, want := catalog.suggestions[0].Value, "go"; got != want {
		t.Errorf("Catalog stored value = %q, want %q", got, want)
	}
}
