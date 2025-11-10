# Usage Examples

This document provides practical examples of using the OpenWordle game logic.

## Basic Game Flow

### 1. Initialize Word Manager

```go
package main

import (
    "fmt"
    "github.com/EngineerNV/OpenWordle/word_manager"
)

func main() {
    wm := word_manager.NewWordManager()
    
    // Load the word dictionary
    err := wm.LoadWords("words.txt")
    if err != nil {
        fmt.Printf("Error loading words: %v\n", err)
        return
    }
    
    fmt.Printf("Loaded %d words\n", wm.GetWordCount())
}
```

### 2. Validate User Input

```go
func validateGuess(wm *word_manager.WordManager, guess string) bool {
    // Check if the word is exactly 5 letters
    if len(guess) != 5 {
        fmt.Println("Word must be exactly 5 letters")
        return false
    }
    
    // Check if the word exists in dictionary
    if !wm.IsValidWord(guess) {
        fmt.Println("Not a valid word")
        return false
    }
    
    return true
}
```

### 3. Check a Guess

```go
import "github.com/EngineerNV/OpenWordle/game_board"

func checkGuess(answer, guess string) game_board.BlockRow {
    br := game_board.BlockRow{}
    br.CheckGuess(answer, guess)
    return br
}

// Display the result
func displayResult(br game_board.BlockRow) {
    for i, block := range br.Blocks {
        symbol := " "
        switch block.Color {
        case "GREEN":
            symbol = "🟩"
        case "YELLOW":
            symbol = "🟨"
        case "GRAY":
            symbol = "⬜"
        }
        fmt.Printf("%s ", symbol)
    }
    fmt.Println()
}
```

### 4. Complete Game Example

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    
    "github.com/EngineerNV/OpenWordle/game_board"
    "github.com/EngineerNV/OpenWordle/word_manager"
)

func main() {
    // Initialize
    wm := word_manager.NewWordManager()
    if err := wm.LoadWords("words.txt"); err != nil {
        fmt.Printf("Error loading words: %v\n", err)
        return
    }
    
    // Select target word
    targetWord := wm.GetRandomWord()
    maxAttempts := 6
    attempts := 0
    
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to OpenWordle!")
    fmt.Println("Guess the 5-letter word. You have 6 attempts.")
    fmt.Println()
    
    for attempts < maxAttempts {
        fmt.Printf("Attempt %d/%d: ", attempts+1, maxAttempts)
        
        // Read user input
        input, _ := reader.ReadString('\n')
        guess := strings.TrimSpace(strings.ToUpper(input))
        
        // Validate input
        if len(guess) != 5 {
            fmt.Println("Please enter exactly 5 letters")
            continue
        }
        
        if !wm.IsValidWord(guess) {
            fmt.Println("Not a valid word")
            continue
        }
        
        // Check guess
        br := game_board.BlockRow{}
        isCorrect := br.CheckGuess(targetWord, guess)
        
        // Display result
        for _, block := range br.Blocks {
            symbol := "⬜"
            if block.Color == "GREEN" {
                symbol = "🟩"
            } else if block.Color == "YELLOW" {
                symbol = "🟨"
            }
            fmt.Printf("%s ", symbol)
        }
        fmt.Printf(" %s\n", guess)
        
        attempts++
        
        // Check if won
        if isCorrect {
            fmt.Printf("\n🎉 Congratulations! You guessed it in %d attempts!\n", attempts)
            return
        }
    }
    
    fmt.Printf("\n😞 Game Over! The word was: %s\n", targetWord)
}
```

## Advanced Usage

### Custom Game Modes

#### Hard Mode (Use All Revealed Hints)

```go
type HardModeValidator struct {
    knownGreen  map[int]string  // Position -> Letter
    knownYellow map[string]bool // Letters that must be included
}

func (hmv *HardModeValidator) validateGuess(guess string, previousResult game_board.BlockRow) bool {
    // Check all green letters are in the same position
    for pos, letter := range hmv.knownGreen {
        if string(guess[pos]) != letter {
            fmt.Printf("Position %d must be %s\n", pos+1, letter)
            return false
        }
    }
    
    // Check all yellow letters are included
    for letter := range hmv.knownYellow {
        if !strings.Contains(guess, letter) {
            fmt.Printf("Guess must contain %s\n", letter)
            return false
        }
    }
    
    return true
}
```

#### Daily Challenge Mode

```go
import (
    "crypto/md5"
    "fmt"
    "time"
)

func getDailyWord(wm *word_manager.WordManager) string {
    // Use date as seed for reproducible daily word
    today := time.Now().Format("2006-01-02")
    hash := md5.Sum([]byte(today))
    
    // Use hash to select word
    wordCount := wm.GetWordCount()
    if wordCount == 0 {
        return ""
    }
    
    index := int(hash[0]) % wordCount
    // Note: Would need to add GetWordByIndex method to WordManager
    return wm.GetRandomWord() // Simplified for this example
}
```

### Statistics Tracking

```go
type GameStats struct {
    GamesPlayed  int
    GamesWon     int
    CurrentStreak int
    MaxStreak    int
    GuessDistribution [7]int // Index 0 unused, 1-6 for attempts
}

func (stats *GameStats) recordWin(attempts int) {
    stats.GamesPlayed++
    stats.GamesWon++
    stats.CurrentStreak++
    
    if stats.CurrentStreak > stats.MaxStreak {
        stats.MaxStreak = stats.CurrentStreak
    }
    
    if attempts >= 1 && attempts <= 6 {
        stats.GuessDistribution[attempts]++
    }
}

func (stats *GameStats) recordLoss() {
    stats.GamesPlayed++
    stats.CurrentStreak = 0
}

func (stats *GameStats) winRate() float64 {
    if stats.GamesPlayed == 0 {
        return 0.0
    }
    return float64(stats.GamesWon) / float64(stats.GamesPlayed) * 100
}
```

### Color Formatting for Terminal

```go
const (
    ColorReset  = "\033[0m"
    ColorGreen  = "\033[42m" // Green background
    ColorYellow = "\033[43m" // Yellow background
    ColorGray   = "\033[100m" // Gray background
)

func displayColoredResult(br game_board.BlockRow) {
    for _, block := range br.Blocks {
        color := ColorGray
        if block.Color == "GREEN" {
            color = ColorGreen
        } else if block.Color == "YELLOW" {
            color = ColorYellow
        }
        
        fmt.Printf("%s %s %s", color, block.Char, ColorReset)
    }
    fmt.Println()
}
```

## Testing

### Test Your Game Logic

```go
package main

import (
    "testing"
    "github.com/EngineerNV/OpenWordle/game_board"
)

func TestGameScenario(t *testing.T) {
    answer := "HELLO"
    
    testCases := []struct {
        guess  string
        expected [5]string // Expected colors
    }{
        {
            guess: "HELPS",
            expected: [5]string{"GREEN", "GREEN", "GREEN", "GRAY", "GRAY"},
        },
        {
            guess: "LEMON",
            expected: [5]string{"YELLOW", "GREEN", "GRAY", "YELLOW", "GRAY"},
        },
    }
    
    for _, tc := range testCases {
        br := game_board.BlockRow{}
        br.CheckGuess(answer, tc.guess)
        
        for i, expectedColor := range tc.expected {
            if br.Blocks[i].Color != expectedColor {
                t.Errorf("Guess %s, position %d: got %s, want %s",
                    tc.guess, i, br.Blocks[i].Color, expectedColor)
            }
        }
    }
}
```

## Building and Deploying

### Build for Multiple Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o openwordle-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o openwordle-macos

# Windows
GOOS=windows GOARCH=amd64 go build -o openwordle.exe

# ARM (e.g., Raspberry Pi)
GOOS=linux GOARCH=arm GOARM=7 go build -o openwordle-arm
```

### Create a Release Build

```bash
# Build with optimizations
go build -ldflags="-s -w" -o openwordle

# This reduces binary size by stripping debug information
```

## Next Steps

1. Add web interface using HTML/CSS/JavaScript
2. Create REST API endpoints for multiplayer
3. Implement user authentication
4. Add social sharing features
5. Create mobile app using the Go backend

See [PROBLEM_STATEMENT.md](PROBLEM_STATEMENT.md) for more ideas and future enhancements.
