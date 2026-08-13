package autocomplete

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestServiceSearchReturnsTrieSuggestions(t *testing.T) {
	trie := NewTrie()
	trie.Insert(Suggestion{Value: "javascript", Score: 100})
	trie.Insert(Suggestion{Value: "java", Score: 80})

	service := newServiceWithoutError(t, trie)

	got, err := service.Search("java")
	if err != nil {
		t.Fatalf("Search() error = %#v; want nil", err)
	}

	want := []Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "java", Score: 80},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search() = %#v; want %#v", got, want)
	}
}

func TestServiceSearchPropagatesTrieError(t *testing.T) {
	trie := NewTrie()
	service := newServiceWithoutError(t, trie)
	prefix := strings.Repeat("a", MaxPrefixLength+1)

	results, err := service.Search(prefix)

	if !errors.Is(err, ErrPrefixTooLong) {
		t.Fatalf("Search() error = %v; want %v", err, ErrPrefixTooLong)
	}

	if results != nil {
		t.Errorf("Search() results = %#v; want nil", results)
	}
}

func TestNewServiceRejectsNilTrie(t *testing.T) {
	service, err := NewService(nil)

	if !errors.Is(err, ErrNilTrie) {
		t.Fatalf("NewService() error = %v; want %v", err, ErrNilTrie)
	}

	if service != nil {
		t.Errorf("NewService() = %#v; want nil", service)
	}
}

func TestServiceSearchReturnsEmptySliceForBlankPrefix(t *testing.T) {
	trie := NewTrie()
	service := newServiceWithoutError(t, trie)

	results, err := service.Search("   ")

	if err != nil {
		t.Fatalf("Search() error = %v, want nil", err)
	}

	if results == nil {
		t.Fatal("Search() results = nil; want empty slice")
	}

	if len(results) != 0 {
		t.Errorf("Search() returned %d suggestions; want 0", len(results))
	}
}
