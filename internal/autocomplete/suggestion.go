package autocomplete

import (
	"errors"
	"strings"
)

type Suggestion struct {
	Value string
	Score int
}

func NewSuggestion(value string, score int) (Suggestion, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return Suggestion{}, errors.New("suggestion value cannot be empty")
	}

	if score < 0 {
		return Suggestion{}, errors.New("suggestion score cannot be negative")
	}

	return Suggestion{
		Value: value,
		Score: score,
	}, nil
}
