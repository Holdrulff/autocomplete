package autocomplete

import "testing"

func TestNewTrie(t *testing.T) {
	trie := NewTrie()

	if trie.root == nil {
		t.Fatal("Trie root is nil")
	}

	if trie.root.children == nil {
		t.Fatal("Trie root children map is nil")
	}

	if got, want := len(trie.root.children), 0; got != want {
		t.Errorf("root children = %d; want %d", got, want)
	}
}

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Suggestion{Value: "go", Score: 100})
	gNode, exists := trie.root.children['g']
	if !exists {
		t.Fatal("root does not contain g")
	}

	_, exists = gNode.children['o']
	if !exists {
		t.Fatal("g node does not contain o")
	}
}

func TestTrieInsertSharesPrefix(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Suggestion{Value: "reactjs", Score: 100})
	trie.Insert(Suggestion{Value: "react-native", Score: 90})

	currentNode := trie.root

	for _, character := range "react" {
		childNode, exists := currentNode.children[character]

		if !exists {
			t.Fatalf("node does not contain %q", character)
		}

		currentNode = childNode
	}

	if got, want := len(currentNode.children), 2; got != want {
		t.Errorf("children after react = %d; want %d", got, want)
	}

	if _, exists := currentNode.children['j']; !exists {
		t.Error("react prefix does not branch to j")
	}

	if _, exists := currentNode.children['-']; !exists {
		t.Error("react prefix does not branch to -")
	}
}

func TestTrieInsertStoresSuggestionAtEnd(t *testing.T) {
	trie := NewTrie()

	want := Suggestion{
		Value: "go",
		Score: 100,
	}

	trie.Insert(want)

	endNode := trie.root.children['g'].children['o']

	if !endNode.isTerminal {
		t.Fatal("end node is not terminal")
	}

	if got := endNode.suggestion; got != want {
		t.Errorf("end node suggestion = %#v; want %#v", got, want)
	}
}

func TestTrieInsertCachesSuggestionForEveryPrefix(t *testing.T) {
	trie := NewTrie()

	want := Suggestion{
		Value: "go",
		Score: 100,
	}

	trie.Insert(want)

	nodes := []*trieNode{
		trie.root,
		trie.root.children['g'],
		trie.root.children['g'].children['o'],
	}

	for _, node := range nodes {
		if got, wantCount := len(node.suggestions), 1; got != wantCount {
			t.Fatalf("cached suggestions = %d; want %d", got, wantCount)
		}

		if got := node.suggestions[0]; got != want {
			t.Fatalf("cached suggestion = %#v; want %#v", got, want)
		}
	}
}
