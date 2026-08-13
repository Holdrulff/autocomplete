package autocomplete

import (
	"errors"
)

var ErrNilTrie = errors.New("trie must not be nil")

type Service struct {
	trie *Trie
}

func NewService(trie *Trie) (*Service, error) {
	if trie == nil {
		return nil, ErrNilTrie
	}

	return &Service{
		trie: trie,
	}, nil
}

func (s *Service) Search(prefix string) ([]Suggestion, error) {
	return s.trie.Search(prefix)
}
