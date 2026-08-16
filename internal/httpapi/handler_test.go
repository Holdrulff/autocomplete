package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/Holdrulff/autocomplete/graph"
	"github.com/Holdrulff/autocomplete/internal/autocomplete"
)

func TestHandlerExecutesAutocompleteQuery(t *testing.T) {
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

	resolver, err := graph.NewResolver(service)
	if err != nil {
		t.Fatalf("NewResolver() error = %v; want nil", err)
	}

	handler, err := NewHandler(resolver)
	if err != nil {
		t.Fatalf("NewHandler() error = %v; want nil", err)
	}

	requestBody := `{
		"query": "query { autocomplete(prefix: \"java\") { value score } }"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/query",
		strings.NewReader(requestBody),
	)
	request.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf(
			"status code = %d; want %d",
			responseRecorder.Code,
			http.StatusOK,
		)
	}

	var response struct {
		Data struct {
			Autocomplete []autocomplete.Suggestion `json:"autocomplete"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.NewDecoder(responseRecorder.Body).Decode(&response); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if len(response.Errors) != 0 {
		t.Fatalf("graphql errors = %#v; want none", response.Errors)
	}

	want := []autocomplete.Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "java", Score: 80},
	}

	if !reflect.DeepEqual(response.Data.Autocomplete, want) {
		t.Errorf(
			"autocomplete = %#v; want %#v",
			response.Data.Autocomplete,
			want,
		)
	}
}

func TestNewHandlerRejectsNilResolver(t *testing.T) {
	handler, err := NewHandler(nil)

	if !errors.Is(err, ErrNilResolver) {
		t.Fatalf("NewHandler() error = %v; want %v", err, ErrNilResolver)
	}

	if handler != nil {
		t.Errorf("NewHandler() = %#v; want nil", handler)
	}
}

func TestHandlerReturnsGraphQLErrorForPrefixTooLong(t *testing.T) {
	trie := autocomplete.NewTrie()

	service, err := autocomplete.NewService(trie)
	if err != nil {
		t.Fatalf("NewService() error = %v; want nil", err)
	}

	resolver, err := graph.NewResolver(service)
	if err != nil {
		t.Fatalf("NewResolver() error = %v; want nil", err)
	}

	handler, err := NewHandler(resolver)
	if err != nil {
		t.Fatalf("NewHandler() error = %v; want nil", err)
	}

	prefix := strings.Repeat(
		"a",
		autocomplete.MaxPrefixLength+1,
	)

	requestBody := `{"query":"query { autocomplete(prefix: \"` +
		prefix +
		`\") { value score } }"}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/query",
		strings.NewReader(requestBody),
	)
	request.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf(
			"status code = %d; want %d",
			responseRecorder.Code,
			http.StatusOK,
		)
	}

	var response struct {
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.NewDecoder(responseRecorder.Body).Decode(&response); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if len(response.Errors) == 0 {
		t.Fatal("graphql errors = empty; want prefix-too-long error")
	}

	if !strings.Contains(
		response.Errors[0].Message,
		autocomplete.ErrPrefixTooLong.Error(),
	) {
		t.Errorf(
			"graphql error = %q; want it to contain %q",
			response.Errors[0].Message,
			autocomplete.ErrPrefixTooLong.Error(),
		)
	}
}
