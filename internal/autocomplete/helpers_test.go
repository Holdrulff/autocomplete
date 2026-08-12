package autocomplete

import (
	"os"
	"path/filepath"
	"testing"
)

func writeSnapshotTestFile(t *testing.T, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "snapshot.json")

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("os.WriteFile() error = %v", err)
	}

	return path
}

func searchWithoutError(t *testing.T, trie *Trie, prefix string) []Suggestion {
	t.Helper()

	results, err := trie.Search(prefix)
	if err != nil {
		t.Fatalf("Search(%q) error = %v", prefix, err)
	}

	return results
}
