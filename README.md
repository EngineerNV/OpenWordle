# OpenWordle

A Go implementation of the popular Wordle word-guessing game.

## 📖 Overview

OpenWordle is a recreation of Wordle, featuring both backend game logic and frontend interface. Players have 6 attempts to guess a 5-letter word, with colored feedback provided after each guess.

For detailed information about the project goals and architecture, see [PROBLEM_STATEMENT.md](PROBLEM_STATEMENT.md).

## 🎮 How to Play

1. Guess a 5-letter word
2. Receive feedback:
   - **Green**: Correct letter in correct position
   - **Yellow**: Correct letter in wrong position
   - **Gray**: Letter not in word
3. You have 6 attempts to find the word

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher
- Git (for cloning the repository)

### Installation

```bash
# Clone the repository
git clone https://github.com/EngineerNV/OpenWordle.git
cd OpenWordle

# Initialize Go modules
go mod tidy

# Build the project
go build -o openwordle
```

### Running the Game

```bash
# Run the game
./openwordle
```

Or directly with Go:

```bash
go run .
```

## 📁 Project Structure

```
OpenWordle/
├── game_board/          # Game board and block logic
│   └── block.go         # BlockRow and BoardBlock structures
├── word_manager/        # Word dictionary and validation
│   └── mapWord.go       # Word management functions
├── open_wordle.go       # Main entry point
├── words.txt            # Dictionary of 5-letter words (5,755 words)
├── go.mod               # Go module definition
├── README.md            # This file
└── PROBLEM_STATEMENT.md # Detailed project goals and architecture
```

## 🛠️ Development

### Building

```bash
# Build the project
go build -v

# Build for specific platforms
GOOS=windows GOARCH=amd64 go build -o openwordle.exe
GOOS=linux GOARCH=amd64 go build -o openwordle
GOOS=darwin GOARCH=amd64 go build -o openwordle
```

### Code Formatting

```bash
# Format all Go files
go fmt ./...

# Check formatting
gofmt -l .
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests verbosely
go test -v ./...
```

### Code Quality

```bash
# Run go vet for static analysis
go vet ./...

# Install and run golint (if available)
go install golang.org/x/lint/golint@latest
golint ./...
```

## 📦 Packages

### game_board

Handles the visual representation and game logic for the board.

- `BoardBlock`: Represents a single letter block with character and color
- `BlockRow`: Represents a row of 5 blocks (one guess attempt)
- `CheckGuess()`: Validates a guess against the answer and updates colors

### word_manager

Manages the word dictionary and validation.

- Word loading from `words.txt`
- Word validation
- Random word selection

## 🔧 Current Implementation Status

### ✅ Completed
- Basic project structure
- Block and BlockRow data structures
- Simple guess checking logic
- Word dictionary file (5,755 words)

### 🚧 In Progress
- Word manager implementation
- Complete game state management
- User input handling
- Frontend interface

### 📋 To Do
- Unit tests for all packages
- Integration tests
- Error handling improvements
- Game statistics tracking
- Web interface
- API endpoints

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Coding Guidelines

- Follow Go conventions and idioms
- Use meaningful variable and function names (PascalCase for exported, camelCase for unexported)
- Add comments for exported functions and types
- Write tests for new functionality
- Run `go fmt` before committing

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎯 Future Enhancements

- [ ] Daily word challenges
- [ ] Statistics tracking (win rate, guess distribution)
- [ ] Multiple difficulty levels
- [ ] Multiplayer support
- [ ] Web-based UI with REST API
- [ ] Command-line interface with colors
- [ ] Support for different word lengths (4-8 letters)
- [ ] Hard mode (revealed hints must be used)

## 📚 Resources

- [Original Wordle](https://www.nytimes.com/games/wordle/index.html)
- [Go Documentation](https://go.dev/doc/)
- [Go Project Layout](https://github.com/golang-standards/project-layout)

## 👥 Authors

- **EngineerNV** - Initial work

## 🙏 Acknowledgments

- Inspired by Josh Wardle's original Wordle game
- Word list sourced from common 5-letter word dictionaries
