package graph

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/Holdrulff/autocomplete/internal/autocomplete"
)

func TestAutocompleteResolverReturnsSuggestions(t *testing.T) {
	trie := autocomplete.NewTrie()
	trie.Insert(autocomplete.Suggestion{
		Value: "javascript",
		Score: 100,
	})
	trie.Insert(autocomplete.Suggestion{
		Value: "java",
		Score: 80,
	})

	service, err := autocomplete.NewService(trie)
	if err != nil {
		t.Fatalf("NewService() error = %v; want nil", err)
	}

	resolver, err := NewResolver(service)
	if err != nil {
		t.Fatalf("NewResolver() error = %v; want nil", err)
	}

	got, err := resolver.Query().Autocomplete(
		context.Background(),
		"java",
	)
	if err != nil {
		t.Fatalf("Autocomplete() error = %v; want nil", err)
	}

	want := []autocomplete.Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "java", Score: 80},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Autocomplete() = %#v; want %#v", got, want)
	}
}

func TestNewResolverRejectsNilService(t *testing.T) {
	resolver, err := NewResolver(nil)

	if !errors.Is(err, ErrNilService) {
		t.Fatalf("NewResolver() error = %v; want %v", err, ErrNilService)
	}

	if resolver != nil {
		t.Errorf("NewResolver() = %#v; want nil", resolver)
	}
}

func TestAutocompleteResolverPropagatesPrefixTooLongError(t *testing.T) {
	trie := autocomplete.NewTrie()

	service, err := autocomplete.NewService(trie)
	if err != nil {
		t.Fatalf("NewService() error = %v; want nil", err)
	}

	resolver, err := NewResolver(service)
	if err != nil {
		t.Fatalf("NewResolver() error = %v; want nil", err)
	}

	prefix := strings.Repeat("a", autocomplete.MaxPrefixLength+1)

	results, err := resolver.Query().Autocomplete(
		context.Background(),
		prefix,
	)

	if !errors.Is(err, autocomplete.ErrPrefixTooLong) {
		t.Fatalf(
			"Autocomplete() error = %v; want %v",
			err,
			autocomplete.ErrPrefixTooLong,
		)
	}

	if results != nil {
		t.Errorf("Autocomplete() results = %#v; want nil", results)
	}
}
