package autocomplete

import (
	"reflect"
	"strconv"
	"testing"
)

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

func TestTrieCachesSuggestionByScoreThenName(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Suggestion{Value: "redux", Score: 100})
	trie.Insert(Suggestion{Value: "react-native", Score: 90})
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	want := []Suggestion{
		{Value: "reactjs", Score: 100},
		{Value: "redux", Score: 100},
		{Value: "react-native", Score: 90},
	}

	if got := trie.root.suggestions; !reflect.DeepEqual(got, want) {
		t.Errorf("root suggestions = %#v; want %#v", got, want)
	}
}

func TestTrieLimitsCachedSuggestionsToTwenty(t *testing.T) {
	trie := NewTrie()

	for score := 1; score <= 21; score++ {
		trie.Insert(Suggestion{
			Value: "technology-" + strconv.Itoa(score),
			Score: score,
		})
	}

	got := trie.root.suggestions

	if gotCount, wantCount := len(got), 20; gotCount != wantCount {
		t.Fatalf("cached suggestions = %d; want %d", gotCount, wantCount)
	}

	if gotScore, wantScore := got[0].Score, 21; gotScore != wantScore {
		t.Errorf("first suggestion score = %d, want %d", gotScore, wantScore)
	}

	if gotScore, wantScore := got[19].Score, 2; gotScore != wantScore {
		t.Errorf("lastsuggestion score = %d, want %d", gotScore, wantScore)
	}

	for _, suggestion := range got {
		if suggestion.Score == 1 {
			t.Error("cache contains the lowest-ranked suggestion")
		}
	}
}

func TestTrieFindNodeReturnsPrefixNode(t *testing.T) {
	trie := NewTrie()
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	got := trie.findNode("reac")

	if got == nil {
		t.Fatal("findNode returned nil for an existing prefix")
	}

	if gotCount, wantCount := len(got.suggestions), 1; gotCount != wantCount {
		t.Errorf("cached suggestions = %d; want %d", gotCount, wantCount)
	}
}

func TestTrieFindNodeReturnsNilForUnknownPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	if got := trie.findNode("vue"); got != nil {
		t.Errorf("findNode returned %#v for an unknows prefix; want nil", got)
	}
}

func TestTrieSearchReturnsRankedSuggestions(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Suggestion{Value: "reactjs", Score: 100})
	trie.Insert(Suggestion{Value: "react-native", Score: 80})
	trie.Insert(Suggestion{Value: "redux", Score: 90})

	want := []Suggestion{
		{Value: "reactjs", Score: 100},
		{Value: "react-native", Score: 80},
	}

	got := trie.Search("reac")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search() = %#v; want %#v", got, want)
	}
}

func TestTrieSearchReturnsEmptySliceForUnknownPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	got := trie.Search("vue")

	if got == nil {
		t.Fatal("Search() returned nil; want empty slice")
	}

	if len(got) != 0 {
		t.Errorf("Search() returned %d suggestions; want 0", len(got))
	}
}

func TestTrieSearchReturnsCopyOfCachedSuggestions(t *testing.T) {
	trie := NewTrie()
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	firstResult := trie.Search("reac")
	firstResult[0].Value = "changed"

	secondResult := trie.Search("reac")

	if got, want := secondResult[0].Value, "reactjs"; got != want {
		t.Errorf("Search() cached value = %q; want %q", got, want)
	}
}

func TestTrieSearchNormalizesPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	tests := []struct {
		name   string
		prefix string
	}{
		{
			name:   "uppercase",
			prefix: "REAC",
		},
		{
			name:   "surrounding spaces",
			prefix: "  reac  ",
		},
		{
			name:   "uppercase and surrounding spaces",
			prefix: "  REAC  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trie.Search(tt.prefix)

			if gotCount, wantCount := len(got), 1; gotCount != wantCount {
				t.Errorf(
					"Search(%q) returned %d suggestions; want %d",
					tt.prefix,
					gotCount,
					wantCount,
				)
			}
		})
	}
}

func TestTrieSearchReturnsEmptySliceForBlankPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	tests := []string{
		"",
		"   ",
	}

	for _, prefix := range tests {
		t.Run(prefix, func(t *testing.T) {
			got := trie.Search(prefix)

			if got == nil {
				t.Fatal("Search() returned nil; want empty slice")
			}

			if len(got) != 0 {
				t.Errorf(
					"Search(%q) returned %d suggestions; want 0",
					prefix,
					len(got),
				)
			}
		})
	}
}

func TestTrieInsertNormalizesSuggestionPath(t *testing.T) {
	trie := NewTrie()

	want := Suggestion{
		Value: "ReactJS",
		Score: 100,
	}

	trie.Insert(want)

	got := trie.Search("reac")

	if gotCount, wantCount := len(got), 1; gotCount != wantCount {
		t.Fatalf("Search() returned %d suggestions; want %d", gotCount, wantCount)
	}

	if got[0] != want {
		t.Errorf("Search() = %#v; want %#v", got[0], want)
	}
}

func TestTrieInsertIgnoresBlankSuggestion(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Suggestion{Value: "", Score: 100})
	trie.Insert(Suggestion{Value: "   ", Score: 90})

	if got := len(trie.root.suggestions); got != 0 {
		t.Errorf("root cached suggestions = %d; want 0", got)
	}

	if got := len(trie.root.children); got != 0 {
		t.Errorf("root children = %d; want 0", got)
	}

	if trie.root.isTerminal {
		t.Error("root became terminal after inserting blank suggestions")
	}
}

func TestTrieInsertReplacesDuplicateSuggestion(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Suggestion{Value: "reactjs", Score: 80})
	trie.Insert(Suggestion{Value: "redux", Score: 90})
	trie.Insert(Suggestion{Value: "ReactJS", Score: 100})

	got := trie.Search("r")

	want := []Suggestion{
		{Value: "ReactJS", Score: 100},
		{Value: "redux", Score: 90},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search() = %#v; want %#v", got, want)
	}
}

func TestTrieRanksEqualScoresCaseInsensitively(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Suggestion{Value: "Redux", Score: 100})
	trie.Insert(Suggestion{Value: "reactjs", Score: 100})

	want := []Suggestion{
		{Value: "reactjs", Score: 100},
		{Value: "Redux", Score: 100},
	}

	got := trie.Search("r")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search() = %#v; want %#v", got, want)
	}
}

func TestNewTrieFromCatalog(t *testing.T) {
	catalog := NewCatalog([]Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "java", Score: 90},
		{Value: "python", Score: 80},
	})

	trie := NewTrieFromCatalog(catalog)

	got := trie.Search("java")

	want := []Suggestion{
		{Value: "javascript", Score: 100},
		{Value: "java", Score: 90},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search() = %#v; want %#v", got, want)
	}
}
