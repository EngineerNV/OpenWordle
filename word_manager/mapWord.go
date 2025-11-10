package word_manager

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

// WordManager handles loading and managing the dictionary of valid words
type WordManager struct {
	words       []string
	wordsLoaded bool
}

// NewWordManager creates a new WordManager instance
func NewWordManager() *WordManager {
	return &WordManager{
		words:       make([]string, 0),
		wordsLoaded: false,
	}
}

// LoadWords loads words from the specified file path.
// Each line in the file should contain one word.
// Words are converted to uppercase for consistency.
func (wm *WordManager) LoadWords(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wm.words = make([]string, 0)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		// Only include 5-letter words
		if len(word) == 5 {
			wm.words = append(wm.words, strings.ToUpper(word))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	wm.wordsLoaded = true
	return nil
}

// IsValidWord checks if a word exists in the dictionary.
// The word is compared case-insensitively.
func (wm *WordManager) IsValidWord(word string) bool {
	if !wm.wordsLoaded || len(word) != 5 {
		return false
	}

	upperWord := strings.ToUpper(word)
	for _, w := range wm.words {
		if w == upperWord {
			return true
		}
	}
	return false
}

// GetRandomWord returns a random word from the dictionary.
// Returns an empty string if no words are loaded.
func (wm *WordManager) GetRandomWord() string {
	if !wm.wordsLoaded || len(wm.words) == 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(wm.words))
	return wm.words[randomIndex]
}

// GetWordCount returns the number of words loaded in the dictionary
func (wm *WordManager) GetWordCount() int {
	return len(wm.words)
}

// IsLoaded returns whether words have been successfully loaded
func (wm *WordManager) IsLoaded() bool {
	return wm.wordsLoaded
}
