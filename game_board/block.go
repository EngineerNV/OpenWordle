package block

// Color of block like Wordle
// Gray for bad letter,
// Green for good letter and placement,
// Yellow for good letter bad placement

// Scenarios to code for
// User has all of the letter right in their attempt
// User has some letters right in their attempt
// User has no letters right

type boardBlock struct {
	char, color string
}

// Row of Blocks - 5 letter word
type blockRow struct {
	col [4]boardBlock
}

func (br blockRow) fill_blocks(color string, word string) {
	for i := 0; i < 5; i++ {
		br.col[i].color = color
		br.col[i].char = string(word[i])
	}
}

func (br blockRow) fill_block(index int, color string, letter string) {
	br.col[index].color = color
}

func (br blockRow) checkGuess(ans string, guess string) bool {
	perfectAnswer := true
	if ans == guess {
		br.fill_blocks("GREEN", guess)
		return perfectAnswer
	}

	for i := 0; i < 5; i++ {
		if ans[i] == guess[i] { // Green find
			br.fill_block(i, "GREEN", string(guess[i]))
		}
	}
	return false
}