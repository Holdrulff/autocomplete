package autocomplete

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSnapshotDecoding(t *testing.T) {
	input := `{
		"source": "https://api.stackexchange.com/2.3/tags?site=stackoverflow",
		"attribution": "Stack Overflow data",
		"generated_at": "2026-08-11T16:56:24Z",
		"tags": [
			{
				"name": "javascript",
				"count": 100
			}
		]
	}`

	got, err := decodeSnapshot(strings.NewReader(input))
	if err != nil {
		t.Fatalf("decodeSnapshot() error = %v", err)
	}

	if got.Source == "" {
		t.Error("snapshot source is empty")
	}

	if got.GeneratedAt.IsZero() {
		t.Error("snapshot generation date is zero")
	}

	if gotCount, wantCount := len(got.Tags), 1; gotCount != wantCount {
		t.Fatalf("snapshot tags = %d; want %d", gotCount, wantCount)
	}

	if gotName, wantName := got.Tags[0].Name, "javascript"; gotName != wantName {
		t.Errorf("tag name = %q; want %q", gotName, wantName)
	}

	if got.Tags[0].Count == nil {
		t.Fatal("tag count is nil")
	}

	if gotScore, wantScore := *got.Tags[0].Count, 100; gotScore != wantScore {
		t.Errorf("tag count = %d; want %d", gotScore, wantScore)
	}
}

func TestDecodeSnapshotReturnsErrorForMalformedJSON(t *testing.T) {
	input := `{"source":`

	_, err := decodeSnapshot(strings.NewReader(input))

	if err == nil {
		t.Fatal("decodesnapshot() error = nil; want an error")
	}
}

func TestDecodeSnapshotReturnsErrorForInvalidMetadata(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name: "missing source",
			input: `{
				"attribution": "Stack Overflow data",
				"generated_at": "2026-08-11T16:56:24Z",
				"tags": [{"name": "go", "count": 100}]
			}`,
		},
		{
			name: "missing attribution",
			input: `{
				"source": "https://api.stackexchange.com",
				"generated_at": "2026-08-11T16:56:24Z",
				"tags": [{"name": "go", "count": 100}]
			}`,
		},
		{
			name: "missing generation date",
			input: `{
				"source": "https://api.stackexchange.com",
				"attribution": "Stack Overflow data",
				"tags": [{"name": "go", "count": 100}]
			}`,
		},
		{
			name: "empty tags",
			input: `{
				"source": "https://api.stackexchange.com",
				"attribution": "Stack Overflow data",
				"generated_at": "2026-08-11T16:56:24Z",
				"tags": []
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeSnapshot(strings.NewReader(tt.input))

			if err == nil {
				t.Fatal("decodeSnapshot() error = nil; want an error")
			}
		})
	}
}

func TestSnapshotSuggestionsConvertsTags(t *testing.T) {
	data := snapshot{
		Tags: []snapshotTag{
			{Name: "  javascript  ", Count: intPointer(100)},
			{Name: "go", Count: intPointer(90)},
		},
	}

	got, err := data.suggestions()
	if err != nil {
		t.Fatalf("suggestions() error = %v", err)
	}

	want := []Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "go", Score: 90},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("suggestions() = %#v; want %#v", got, want)
	}
}

func TestSnapshotSuggestionsReturnsIndexedError(t *testing.T) {
	tests := []struct {
		name string
		tag  snapshotTag
	}{
		{
			name: "empty name",
			tag:  snapshotTag{Name: "   ", Count: intPointer(100)},
		},
		{
			name: "negative count",
			tag:  snapshotTag{Name: "go", Count: intPointer(-1)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := snapshot{
				Tags: []snapshotTag{
					{Name: "javascript", Count: intPointer(100)},
					tt.tag,
				},
			}

			_, err := data.suggestions()
			if err == nil {
				t.Fatal("suggestions() error = nil; want an error")
			}

			if !strings.Contains(err.Error(), "tag at index 1") {
				t.Errorf("suggestions() error = %q; want index 1", err)
			}
		})
	}
}

func TestLoadCatalog(t *testing.T) {
	path := writeSnapshotTestFile(t, `{
		"source": "https://api.stackexchange.com",
		"attribution": "Stack Overflow data",
		"generated_at": "2026-08-11T16:56:24Z",
		"tags": [
			{"name": "javascript", "count": 100},
			{"name": "go", "count": 90}
		]
	}`)

	catalog, err := LoadCatalog(path)
	if err != nil {
		t.Fatalf("LoadCatalog() error = %v", err)
	}

	if got, want := catalog.Len(), 2; got != want {
		t.Errorf("Catalog.Len() = %d; want %d", got, want)
	}
}

func TestLoadCatalogReturnsErrorForMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.json")

	_, err := LoadCatalog(path)

	if err == nil {
		t.Fatal("LoadCatalog() error = nil; want an error")
	}
}

func TestSnapshotToTrieIntegration(t *testing.T) {
	path := writeSnapshotTestFile(t, `{
		"source": "https://api.stackexchange.com",
		"attribution": "Stack Overflow data",
		"generated_at": "2026-08-11T16:56:24Z",
		"tags": [
			{"name": "javascript", "count": 100},
			{"name": "java", "count": 90},
			{"name": "python", "count": 80}
		]
	}`)

	catalog, err := LoadCatalog(path)
	if err != nil {
		t.Fatalf("LoadCatalog() error = %v", err)
	}

	trie := NewTrieFromCatalog(catalog)

	got := searchWithoutError(t, trie, "java")

	want := []Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "java", Score: 90},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search() = %#v; want %#v", got, want)
	}
}

func intPointer(value int) *int {
	return &value
}

func TestSnapshotSuggestionsReturnsErrorForMissingCount(t *testing.T) {
	data := snapshot{
		Tags: []snapshotTag{
			{Name: "go"},
		},
	}

	_, err := data.suggestions()

	if err == nil {
		t.Fatal("suggestions() error = nil; want an error")
	}

	if !strings.Contains(err.Error(), "count is required") {
		t.Errorf(
			"suggestions() error = %q; want missing count message",
			err,
		)
	}
}

func TestDecodeSnapshotRejectsTrailingContent(t *testing.T) {
	validSnapshot := `{
		"source": "https://api.stackexchange.com",
		"attribution": "Stack Overflow data",
		"generated_at": "2026-08-11T16:56:24Z",
		"tags": [
			{"name": "go", "count": 100}
		]
	}`

	tests := []struct {
		name     string
		trailing string
	}{
		{
			name:     "second JSON object",
			trailing: `{"unexpected": true}`,
		},
		{
			name:     "malformed trailing content",
			trailing: `{invalid`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := validSnapshot + tt.trailing

			_, err := decodeSnapshot(strings.NewReader(input))

			if err == nil {
				t.Fatal("decodeSnapshot() error = nil; want an error")
			}
		})
	}
}
