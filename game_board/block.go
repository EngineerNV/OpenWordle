package game_board

// Color represents the feedback color for a block in Wordle:
// - "GRAY" for incorrect letter
// - "GREEN" for correct letter in correct position
// - "YELLOW" for correct letter in wrong position

// BoardBlock represents a single block in the Wordle game board.
// It contains a character and its associated color based on the guess feedback.
type BoardBlock struct {
	Char  string
	Color string
}

// BlockRow represents a row of 5 blocks, corresponding to a 5-letter word guess.
type BlockRow struct {
	Blocks [5]BoardBlock // Fixed: Changed from [4] to [5] to match 5-letter words
}

// FillBlocks fills all blocks in the row with the specified color and word.
// The word parameter should be a 5-letter string.
func (br *BlockRow) FillBlocks(color string, word string) {
	if len(word) < 5 {
		// Handle words shorter than 5 letters
		for i := 0; i < len(word) && i < 5; i++ {
			br.Blocks[i].Color = color
			br.Blocks[i].Char = string(word[i])
		}
		return
	}

	for i := 0; i < 5; i++ {
		br.Blocks[i].Color = color
		br.Blocks[i].Char = string(word[i])
	}
}

// FillBlock fills a single block at the specified index with a color and letter.
// The index should be between 0 and 4 (inclusive).
func (br *BlockRow) FillBlock(index int, color string, letter string) {
	if index >= 0 && index < 5 {
		br.Blocks[index].Color = color
		br.Blocks[index].Char = letter
	}
}

// CheckGuess compares a guess against the answer and updates the block colors accordingly.
// Returns true if the guess matches the answer perfectly, false otherwise.
//
// Color logic:
// - GREEN: Letter is correct and in the right position
// - YELLOW: Letter is in the word but in the wrong position (to be implemented)
// - GRAY: Letter is not in the word (to be implemented)
func (br *BlockRow) CheckGuess(answer string, guess string) bool {
	if len(answer) != 5 || len(guess) != 5 {
		return false
	}

	// Check for perfect match
	if answer == guess {
		br.FillBlocks("GREEN", guess)
		return true
	}

	// Initialize all blocks as GRAY
	br.FillBlocks("GRAY", guess)

	// First pass: Mark exact matches (GREEN)
	answerChars := []rune(answer)
	guessChars := []rune(guess)
	used := make([]bool, 5) // Track which answer positions have been matched

	for i := 0; i < 5; i++ {
		if answerChars[i] == guessChars[i] {
			br.FillBlock(i, "GREEN", string(guess[i]))
			used[i] = true
		}
	}

	// Second pass: Mark letters in wrong position (YELLOW)
	for i := 0; i < 5; i++ {
		if br.Blocks[i].Color == "GREEN" {
			continue // Skip already matched letters
		}

		// Check if this letter exists elsewhere in the answer
		for j := 0; j < 5; j++ {
			if !used[j] && guessChars[i] == answerChars[j] {
				br.FillBlock(i, "YELLOW", string(guess[i]))
				used[j] = true
				break
			}
		}
	}

	return false
}
