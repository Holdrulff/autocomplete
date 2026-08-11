package autocomplete

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
	currentNode := t.root
	currentNode.addSuggestion(suggestion)

	for _, character := range suggestion.Value {
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

func (n *trieNode) addSuggestion(suggestion Suggestion) {
	n.suggestions = append(n.suggestions, suggestion)
}
