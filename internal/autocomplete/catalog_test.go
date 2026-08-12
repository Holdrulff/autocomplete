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

func TestCatalogSuggestionsReturnsCopy(t *testing.T) {
	catalog := NewCatalog([]Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "go", Score: 90},
	})

	got := catalog.Suggestions()

	if gotCount, wantCount := len(got), 2; gotCount != wantCount {
		t.Fatalf("Suggestions() returned %d items; want %d", gotCount, wantCount)
	}

	got[0].Value = "changed"

	secondResult := catalog.Suggestions()

	if gotValue, wantValue := secondResult[0].Value, "javascript"; gotValue != wantValue {
		t.Errorf("stored suggestion = %q; want %q", gotValue, wantValue)
	}
}
