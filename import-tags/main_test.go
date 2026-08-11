package main

import "testing"

func TestNormalizedAndDeduplicate(t *testing.T) {
	input := []tag{
		{Name: " Javascript ", Count: 100},
		{Name: "javascript", Count: 120},
		{Name: "Go", Count: 90},
		{Name: "  ", Count: 50},
	}

	got := normalizeAndDeduplicate(input)
	if gotCount, wantCount := len(got), 2; gotCount != wantCount {
		t.Fatalf("total tags = %d; wanted %d", gotCount, wantCount)
	}

	gotByName := make(map[string]tag)

	for _, currentTag := range got {
		gotByName[currentTag.Name] = currentTag
	}

	if gotByName["javascript"].Count != 120 {
		t.Errorf(
			"contagem de javascript = %d; quero 120",
			gotByName["javascript"].Count,
		)
	}

	if gotByName["go"].Count != 90 {
		t.Errorf(
			"contagem de go = %d; quero 90",
			gotByName["go"].Count,
		)
	}
}
