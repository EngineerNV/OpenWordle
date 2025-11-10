# OpenWordle - Problem Statement

## Project Overview

OpenWordle is a recreation of the popular Wordle game, implementing both backend game logic and frontend interface for educational and entertainment purposes.

## Game Description

Wordle is a word-guessing game where:
- Players have 6 attempts to guess a 5-letter word
- After each guess, the game provides feedback using colored blocks:
  - **Green**: Letter is correct and in the right position
  - **Yellow**: Letter is in the word but in the wrong position
  - **Gray**: Letter is not in the word at all
- The goal is to guess the correct word within the allowed attempts

## Project Goals

1. **Backend Implementation**: Create robust game logic in Go that handles:
   - Word validation against a dictionary
   - Guess checking and feedback generation
   - Game state management
   - Win/loss conditions

2. **Frontend Interface**: Develop a user interface that displays:
   - Game board with colored blocks
   - Current guesses and remaining attempts
   - Input validation and error messages
   - Game results

3. **Code Quality**: Ensure the codebase is:
   - Well-structured and modular
   - Following Go best practices and conventions
   - Properly documented
   - Easy to extend and improve

## Current Implementation Status

### Completed Components

1. **Game Board Package** (`game_board/`):
   - `BoardBlock`: Represents a single letter block with color state
   - `BlockRow`: Represents a row of 5 blocks (one guess)
   - Basic guess checking logic

2. **Word Manager Package** (`word_manager/`):
   - Placeholder for dictionary management (to be implemented)

3. **Word List** (`words.txt`):
   - Contains 5,755 five-letter words for the game dictionary

### Components to Implement/Improve

1. **Word Manager**:
   - Load words from `words.txt`
   - Validate user guesses against the dictionary
   - Select random words for the game
   - Handle word comparison logic

2. **Game Controller**:
   - Manage game state (attempts remaining, current guesses, etc.)
   - Coordinate between word manager and game board
   - Handle win/loss conditions
   - Provide game statistics

3. **Frontend Interface**:
   - Display game board visually
   - Accept user input
   - Show colored feedback
   - Display game messages and results

4. **Testing**:
   - Unit tests for game logic
   - Integration tests for complete game flow
   - Edge case handling

## Technical Requirements

- **Language**: Go 1.18 or higher
- **Architecture**: Modular package structure
- **Code Style**: Follow standard Go conventions
- **Documentation**: Clear comments and usage examples
- **Build System**: Standard Go toolchain (`go build`, `go test`)

## Future Enhancements

- Multiple difficulty levels
- Statistics tracking (win rate, guess distribution)
- Daily word challenges
- Multiplayer support
- Web-based interface
- API for programmatic access
- Support for different word lengths

## Success Criteria

The project will be considered successfully initialized when:
1. All Go code builds without errors
2. Code follows Go naming conventions and best practices
3. Basic game logic is functional and testable
4. Documentation clearly explains how to run and extend the project
5. Repository structure supports future development
