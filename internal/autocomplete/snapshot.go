package autocomplete

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type snapshotTag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type snapshot struct {
	Source      string        `json:"source"`
	Attribution string        `json:"attribution"`
	GeneratedAt time.Time     `json:"generated_at"`
	Tags        []snapshotTag `json:"tags"`
}

func decodeSnapshot(reader io.Reader) (snapshot, error) {
	var data snapshot

	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return snapshot{}, fmt.Errorf("decode snapshot: %w", err)
	}

	if err := data.validate(); err != nil {
		return snapshot{}, fmt.Errorf("validate snapshot: %w", err)
	}

	return data, nil
}

func (s snapshot) validate() error {
	if strings.TrimSpace(s.Source) == "" {
		return errors.New("snapshot source cannot be empty")
	}

	if strings.TrimSpace(s.Attribution) == "" {
		return errors.New("snapshot attribution cannot be empty")
	}

	if s.GeneratedAt.IsZero() {
		return errors.New("snapshot generation date cannot be empty")
	}

	if len(s.Tags) == 0 {
		return errors.New("snapshot tags cannot be empty")
	}

	return nil
}

func (s snapshot) suggestions() ([]Suggestion, error) {
	suggestions := make([]Suggestion, 0, len(s.Tags))

	for index, currentTag := range s.Tags {
		suggestion, err := NewSuggestion(currentTag.Name, currentTag.Count)
		if err != nil {
			return nil, fmt.Errorf("tag at index %d: %w", index, err)
		}

		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}

func LoadCatalog(path string) (Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return Catalog{}, fmt.Errorf("open snapshot: %w", err)
	}
	defer file.Close()

	data, err := decodeSnapshot(file)
	if err != nil {
		return Catalog{}, err
	}

	suggestions, err := data.suggestions()
	if err != nil {
		return Catalog{}, fmt.Errorf("convert snapshot tags: %w", err)
	}
	return NewCatalog(suggestions), nil
}
