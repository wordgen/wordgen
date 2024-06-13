package wordgen

import (
	"strings"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	gen := NewGenerator()

	// Verify default values
	if len(gen.Words) != 0 {
		t.Errorf("Expected wordlist to be empty, got %d words", len(gen.Words))
	}
	if gen.Count != 1 {
		t.Errorf("Expected Count to be 1, got %d", gen.Count)
	}
	if gen.Casing != "" {
		t.Errorf("Expected Casing to be empty, got %s", gen.Casing)
	}
	if gen.Separator != " " {
		t.Errorf("Expected Separator to be ' ', got %s", gen.Separator)
	}
}

func TestGenerateWithEmptyWordlist(t *testing.T) {
	gen := NewGenerator()

	// Generate words with empty wordlist
	_, err := gen.Generate()
	if err == nil {
		t.Errorf("Expected error for empty wordlist, got nil")
	}
}

func TestGenerate(t *testing.T) {
	gen := NewGenerator()
	gen.Words = []string{"apple", "banana", "cherry"}

	// Generate words using the generator
	result, err := gen.Generate()
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	// Verify the generated word count
	resultWords := strings.Split(result, gen.Separator)
	if len(resultWords) != gen.Count {
		t.Errorf("Expected %d words, got %d", gen.Count, len(resultWords))
	}
}

func TestGenerateWithSettings(t *testing.T) {
	gen := NewGenerator()
	gen.Words = []string{"apple", "banana", "cherry"}

	// Test with Count set to 3
	gen.Count = 3
	result, err := gen.Generate()
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
	resultWords := strings.Split(result, gen.Separator)
	if len(resultWords) != gen.Count {
		t.Errorf("Expected %d words, got %d", gen.Count, len(resultWords))
	}

	// Test with Casing set to "upper"
	gen.Casing = "upper"
	result, err = gen.Generate()
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
	for _, word := range strings.Split(result, gen.Separator) {
		if word != strings.ToUpper(word) {
			t.Errorf("Expected word in upper case, got %s", word)
		}
	}

	// Test with Separator set to ", "
	gen.Separator = ", "
	result, err = gen.Generate()
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
	if !strings.Contains(result, gen.Separator) {
		t.Errorf("Expected separator %q, but result was %q", gen.Separator, result)
	}
}
