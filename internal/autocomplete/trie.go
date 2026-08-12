package autocomplete

import (
	"sort"
	"strings"
)

type trieNode struct {
	children    map[rune]*trieNode
	suggestion  Suggestion
	suggestions []Suggestion
	isTerminal  bool
}

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{
			children: make(map[rune]*trieNode),
		},
	}
}

func (t *Trie) Insert(suggestion Suggestion) {

	normalizedValue := strings.ToLower(strings.TrimSpace(suggestion.Value))

	if normalizedValue == "" {
		return
	}

	currentNode := t.root
	currentNode.addSuggestion(suggestion)

	for _, character := range normalizedValue {
		childNode, exists := currentNode.children[character]

		if !exists {
			childNode = &trieNode{
				children: make(map[rune]*trieNode),
			}

			currentNode.children[character] = childNode
		}

		currentNode = childNode
		currentNode.addSuggestion(suggestion)
	}

	currentNode.suggestion = suggestion
	currentNode.isTerminal = true
}

const maxSuggestions = 20

func (n *trieNode) addSuggestion(suggestion Suggestion) {
	normalizedValue := strings.ToLower(strings.TrimSpace(suggestion.Value))
	replaced := false

	for index, existingSuggestion := range n.suggestions {
		existingValue := strings.ToLower(
			strings.TrimSpace(existingSuggestion.Value),
		)

		if existingValue == normalizedValue {
			n.suggestions[index] = suggestion
			replaced = true
			break
		}
	}

	if !replaced {
		n.suggestions = append(n.suggestions, suggestion)
	}

	sort.Slice(n.suggestions, func(i, j int) bool {
		left := n.suggestions[i]
		right := n.suggestions[j]

		if left.Score == right.Score {
			leftValue := strings.ToLower(strings.TrimSpace(left.Value))
			rightValue := strings.ToLower(strings.TrimSpace(right.Value))

			if leftValue == rightValue {
				return left.Value < right.Value
			}

			return leftValue < rightValue
		}

		return left.Score > right.Score
	})

	if len(n.suggestions) > maxSuggestions {
		n.suggestions = n.suggestions[:maxSuggestions]
	}
}

func (t *Trie) findNode(prefix string) *trieNode {
	currentNode := t.root

	for _, character := range prefix {
		childNode, exists := currentNode.children[character]
		if !exists {
			return nil
		}

		currentNode = childNode
	}

	return currentNode
}

func (t *Trie) Search(prefix string) []Suggestion {
	prefix = strings.ToLower(strings.TrimSpace(prefix))
	if prefix == "" {
		return []Suggestion{}
	}

	node := t.findNode(prefix)
	if node == nil {
		return []Suggestion{}
	}

	results := make([]Suggestion, len(node.suggestions))
	copy(results, node.suggestions)

	return results
}

func NewTrieFromCatalog(catalog Catalog) *Trie {
	trie := NewTrie()

	for _, suggestion := range catalog.Suggestions() {
		trie.Insert(suggestion)
	}

	return trie
}
