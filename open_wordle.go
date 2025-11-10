package main

import (
	"fmt"
	"github.com/EngineerNV/OpenWordle/game_board"
	"github.com/EngineerNV/OpenWordle/word_manager"
)

func main() {
	fmt.Println("=== OpenWordle - A Wordle Clone in Go ===")
	fmt.Println()

	// Initialize the word manager
	wm := word_manager.NewWordManager()

	// Load words from the dictionary
	err := wm.LoadWords("words.txt")
	if err != nil {
		fmt.Printf("Error loading words: %v\n", err)
		fmt.Println("Make sure words.txt is in the current directory.")
		return
	}

	fmt.Printf("Loaded %d words from dictionary\n", wm.GetWordCount())
	fmt.Println()

	// Get a random target word
	targetWord := wm.GetRandomWord()
	fmt.Printf("Random word selected (for demo): %s\n", targetWord)
	fmt.Println()

	// Example 1: Perfect match
	fmt.Println("Example 1: Testing perfect match")
	br1 := game_board.BlockRow{}
	guess1 := targetWord
	result1 := br1.CheckGuess(targetWord, guess1)
	fmt.Printf("  Guess: %s | Answer: %s | Perfect Match: %v\n", guess1, targetWord, result1)
	for i, block := range br1.Blocks {
		fmt.Printf("  Block %d: %s [%s]\n", i, block.Char, block.Color)
	}
	fmt.Println()

	// Example 2: Partial match
	fmt.Println("Example 2: Testing partial match")
	br2 := game_board.BlockRow{}

	// Validate the test guess
	testGuess := "HELLO"
	if !wm.IsValidWord(testGuess) {
		// Try to find a valid word from our dictionary
		testGuess = wm.GetRandomWord()
	}

	result2 := br2.CheckGuess(targetWord, testGuess)
	fmt.Printf("  Guess: %s | Answer: %s | Perfect Match: %v\n", testGuess, targetWord, result2)
	for i, block := range br2.Blocks {
		fmt.Printf("  Block %d: %s [%s]\n", i, block.Char, block.Color)
	}
	fmt.Println()

	// Example 3: Word validation
	fmt.Println("Example 3: Word validation")
	fmt.Printf("  Is 'HELLO' a valid word? %v\n", wm.IsValidWord("HELLO"))
	fmt.Printf("  Is 'XYZQW' a valid word? %v\n", wm.IsValidWord("XYZQW"))
	fmt.Println()

	fmt.Println("=== Demo Complete ===")
	fmt.Println("This is the initial implementation.")
	fmt.Println("The game logic is functional and ready for further development!")
}
