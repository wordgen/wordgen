package wordgen

import (
	"strings"
	"testing"
	"time"

	"golang.org/x/text/cases"
)

func TestGenerate(t *testing.T) {
	g := NewGenerator()
	g.Words = []string{"apple", "banana", "cherry"}
	g.Count = 3
	g.Separator = "-"

	testCases := []struct {
		casing   string
		expected func(string) string
	}{
		{"upper", cases.Upper(g.Language).String},
		{"lower", cases.Lower(g.Language).String},
		{"title", cases.Title(g.Language).String},
	}

	for _, tc := range testCases {
		g.Casing = tc.casing
		result, err := g.Generate()
		if err != nil {
			t.Errorf("Generate() returned an error: %v", err)
		}

		expectedWords := strings.Split(result, "-")
		if len(expectedWords) != 3 {
			t.Errorf("Generate() expected 3 words, got %d", len(expectedWords))
		}

		for _, word := range strings.Split(result, g.Separator) {
			expectedWord := tc.expected(word)
			if word != expectedWord {
				t.Errorf("Generate() with casing %s failed: got %s, want %s", tc.casing, word, expectedWord)
			}
		}
	}
}

func TestGeneratePerformance(t *testing.T) {
	g := NewGenerator()
	g.Words = []string{"wordgen"}
	g.Casing = "title"
	g.Separator = "."

	tests := []struct {
		name  string
		count int
	}{
		{"10k words", 10000},
		{"100k words", 100000},
		{"1m words", 1000000},
		{"10m words", 10000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g.Count = tt.count
			start := time.Now()
			_, err := g.Generate()
			if err != nil {
				t.Errorf("Generate() returned an error: %v", err)
			}

			duration := time.Since(start)
			t.Logf("%s took %v", tt.name, duration)
		})
	}
}
