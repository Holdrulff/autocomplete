package autocomplete

import "testing"

func TestNewSuggestion(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		score   int
		want    Suggestion
		wantErr bool
	}{
		{
			name:  "creates a valid suggestion",
			value: "golang",
			score: 100,
			want: Suggestion{
				Value: "golang",
				Score: 100,
			},
		},
		{
			name:  "trims surrounding whitespace",
			value: "  react  ",
			score: 80,
			want: Suggestion{
				Value: "react",
				Score: 80,
			},
		},
		{
			name:    "rejects an empty value",
			value:   "",
			score:   10,
			wantErr: true,
		},
		{
			name:    "rejects whitespace only value",
			value:   "   ",
			score:   10,
			wantErr: true,
		},
		{
			name:    "rejects a negative score",
			value:   "docker",
			score:   -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSuggestion(tt.value, tt.score)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewSuggestion() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if got != tt.want {
				t.Errorf("NewSuggestion() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
