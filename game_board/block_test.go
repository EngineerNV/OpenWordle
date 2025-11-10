package game_board

import "testing"

func TestBlockRowFillBlocks(t *testing.T) {
	br := BlockRow{}
	br.FillBlocks("GREEN", "HELLO")

	// Check that all blocks are filled correctly
	for i := 0; i < 5; i++ {
		if br.Blocks[i].Color != "GREEN" {
			t.Errorf("Block %d color = %s; want GREEN", i, br.Blocks[i].Color)
		}
	}

	expectedChars := []string{"H", "E", "L", "L", "O"}
	for i := 0; i < 5; i++ {
		if br.Blocks[i].Char != expectedChars[i] {
			t.Errorf("Block %d char = %s; want %s", i, br.Blocks[i].Char, expectedChars[i])
		}
	}
}

func TestBlockRowCheckGuess_PerfectMatch(t *testing.T) {
	br := BlockRow{}
	answer := "HELLO"
	guess := "HELLO"

	result := br.CheckGuess(answer, guess)

	if !result {
		t.Error("Expected perfect match to return true")
	}

	// All blocks should be green
	for i := 0; i < 5; i++ {
		if br.Blocks[i].Color != "GREEN" {
			t.Errorf("Block %d color = %s; want GREEN for perfect match", i, br.Blocks[i].Color)
		}
	}
}

func TestBlockRowCheckGuess_NoMatch(t *testing.T) {
	br := BlockRow{}
	answer := "HELLO"
	guess := "WXYZQ"

	result := br.CheckGuess(answer, guess)

	if result {
		t.Error("Expected no match to return false")
	}

	// All blocks should be gray (no matching letters)
	for i := 0; i < 5; i++ {
		if br.Blocks[i].Color != "GRAY" {
			t.Errorf("Block %d color = %s; want GRAY for no match", i, br.Blocks[i].Color)
		}
	}
}

func TestBlockRowCheckGuess_PartialMatch(t *testing.T) {
	br := BlockRow{}
	answer := "HELLO"
	guess := "HELPS"

	result := br.CheckGuess(answer, guess)

	if result {
		t.Error("Expected partial match to return false")
	}

	// H should be green (correct position)
	if br.Blocks[0].Color != "GREEN" {
		t.Errorf("Block 0 color = %s; want GREEN (H in correct position)", br.Blocks[0].Color)
	}

	// E should be green (correct position)
	if br.Blocks[1].Color != "GREEN" {
		t.Errorf("Block 1 color = %s; want GREEN (E in correct position)", br.Blocks[1].Color)
	}

	// L should be green (correct position)
	if br.Blocks[2].Color != "GREEN" {
		t.Errorf("Block 2 color = %s; want GREEN (L in correct position)", br.Blocks[2].Color)
	}

	// P should be gray (not in word)
	if br.Blocks[3].Color != "GRAY" {
		t.Errorf("Block 3 color = %s; want GRAY (P not in word)", br.Blocks[3].Color)
	}

	// S should be gray (not in word)
	if br.Blocks[4].Color != "GRAY" {
		t.Errorf("Block 4 color = %s; want GRAY (S not in word)", br.Blocks[4].Color)
	}
}

func TestBlockRowCheckGuess_YellowMatch(t *testing.T) {
	br := BlockRow{}
	answer := "HELLO"
	guess := "LEMON"

	result := br.CheckGuess(answer, guess)

	if result {
		t.Error("Expected partial match to return false")
	}

	// L should be yellow (in word but wrong position)
	if br.Blocks[0].Color != "YELLOW" {
		t.Errorf("Block 0 color = %s; want YELLOW (L in word but wrong position)", br.Blocks[0].Color)
	}

	// E should be green (correct position)
	if br.Blocks[1].Color != "GREEN" {
		t.Errorf("Block 1 color = %s; want GREEN (E in correct position)", br.Blocks[1].Color)
	}
}

func TestBlockRowFillBlock(t *testing.T) {
	br := BlockRow{}

	br.FillBlock(2, "YELLOW", "A")

	if br.Blocks[2].Color != "YELLOW" {
		t.Errorf("Block 2 color = %s; want YELLOW", br.Blocks[2].Color)
	}

	if br.Blocks[2].Char != "A" {
		t.Errorf("Block 2 char = %s; want A", br.Blocks[2].Char)
	}
}

func TestBlockRowCheckGuess_InvalidLength(t *testing.T) {
	br := BlockRow{}
	answer := "HELLO"
	guess := "HI"

	result := br.CheckGuess(answer, guess)

	if result {
		t.Error("Expected invalid length to return false")
	}
}
