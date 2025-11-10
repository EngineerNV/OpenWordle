package word_manager

import (
	"os"
	"testing"
)

func TestNewWordManager(t *testing.T) {
	wm := NewWordManager()

	if wm == nil {
		t.Error("NewWordManager returned nil")
	}

	if wm.IsLoaded() {
		t.Error("Expected IsLoaded to be false before loading words")
	}

	if wm.GetWordCount() != 0 {
		t.Errorf("GetWordCount = %d; want 0 before loading", wm.GetWordCount())
	}
}

func TestLoadWords(t *testing.T) {
	// Create a temporary test file
	tmpFile, err := os.CreateTemp("", "test_words_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write test words
	testWords := "hello\nworld\ntests\nvalid\nwords\n"
	if _, err := tmpFile.WriteString(testWords); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	wm := NewWordManager()
	err = wm.LoadWords(tmpFile.Name())

	if err != nil {
		t.Errorf("LoadWords failed: %v", err)
	}

	if !wm.IsLoaded() {
		t.Error("Expected IsLoaded to be true after loading")
	}

	if wm.GetWordCount() != 5 {
		t.Errorf("GetWordCount = %d; want 5", wm.GetWordCount())
	}
}

func TestLoadWords_NonExistentFile(t *testing.T) {
	wm := NewWordManager()
	err := wm.LoadWords("nonexistent_file.txt")

	if err == nil {
		t.Error("Expected error when loading non-existent file")
	}

	if wm.IsLoaded() {
		t.Error("Expected IsLoaded to be false after failed load")
	}
}

func TestIsValidWord(t *testing.T) {
	// Create a temporary test file
	tmpFile, err := os.CreateTemp("", "test_words_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write test words
	testWords := "hello\nworld\n"
	if _, err := tmpFile.WriteString(testWords); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	wm := NewWordManager()
	wm.LoadWords(tmpFile.Name())

	tests := []struct {
		word  string
		valid bool
	}{
		{"HELLO", true},
		{"hello", true},
		{"HeLLo", true},
		{"WORLD", true},
		{"TESTS", false},
		{"XYZQW", false},
		{"HI", false},    // Too short
		{"TOOLONG", false}, // Too long
	}

	for _, tt := range tests {
		result := wm.IsValidWord(tt.word)
		if result != tt.valid {
			t.Errorf("IsValidWord(%s) = %v; want %v", tt.word, result, tt.valid)
		}
	}
}

func TestIsValidWord_NotLoaded(t *testing.T) {
	wm := NewWordManager()

	if wm.IsValidWord("HELLO") {
		t.Error("Expected IsValidWord to return false when no words are loaded")
	}
}

func TestGetRandomWord(t *testing.T) {
	// Create a temporary test file
	tmpFile, err := os.CreateTemp("", "test_words_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write test words
	testWords := "hello\nworld\ntests\n"
	if _, err := tmpFile.WriteString(testWords); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	wm := NewWordManager()
	wm.LoadWords(tmpFile.Name())

	word := wm.GetRandomWord()

	if word == "" {
		t.Error("GetRandomWord returned empty string")
	}

	if len(word) != 5 {
		t.Errorf("GetRandomWord returned word of length %d; want 5", len(word))
	}

	// Verify the word is one of our test words
	validWords := []string{"HELLO", "WORLD", "TESTS"}
	found := false
	for _, validWord := range validWords {
		if word == validWord {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("GetRandomWord returned unexpected word: %s", word)
	}
}

func TestGetRandomWord_NotLoaded(t *testing.T) {
	wm := NewWordManager()

	word := wm.GetRandomWord()

	if word != "" {
		t.Errorf("GetRandomWord = %s; want empty string when not loaded", word)
	}
}

func TestLoadWords_OnlyFiveLetterWords(t *testing.T) {
	// Create a temporary test file with various length words
	tmpFile, err := os.CreateTemp("", "test_words_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Mix of different length words
	testWords := "hi\nhello\nworld\ntoolongword\nok\ntests\n"
	if _, err := tmpFile.WriteString(testWords); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	wm := NewWordManager()
	err = wm.LoadWords(tmpFile.Name())

	if err != nil {
		t.Errorf("LoadWords failed: %v", err)
	}

	// Should only load 5-letter words: hello, world, tests
	if wm.GetWordCount() != 3 {
		t.Errorf("GetWordCount = %d; want 3 (only 5-letter words)", wm.GetWordCount())
	}
}
