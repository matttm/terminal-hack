# terminal-hack

Status: stable, with occasional improvements.

## Overview

`terminal-hack` is a terminal-based game written in Go, inspired by the Fallout 3 terminal hacking mini‑game. A 2D grid is filled with words and single‑character "duds." Your goal is to find the winning word within a limited number of attempts.

Each incorrect guess reveals how many characters match **exactly** (same character and same index) with the winning word. Example: if the winning word is `BOOK` and you guess `BANG`, you have 1 correct character (`B` at index 0).

If you want the full retro vibe, try it in a CRT‑style terminal emulator like `cool-retro-term` (optional).

<img width="1728" alt="Screenshot 2025-01-22 at 5 16 23 PM" src="https://github.com/user-attachments/assets/5ad793ca-6147-46fb-abb5-56d03efec0a9" />

## Features

- Fallout‑style word grid with random "dud" characters.
- Cursor navigation across the grid.
- Limited lives with immediate feedback after each guess.
- File‑based logging for debugging (`./logs/`).
- Configurable layout and gameplay constants in `internal/constants`.

## Requirements

- Go `1.22+`
- A terminal that supports ANSI color

## Run

```bash
go run .
```

Build a binary instead:

```bash
go build -o terminal-hack .
./terminal-hack
```

## Controls

- `Arrow Keys` move the cursor
- `Enter` selects the current word/character
- `Esc` exits the game

## How It Works

1. A word list is loaded from a random file in `words/` and converted to uppercase.
2. A fixed number of words are inserted into the grid, with random single‑character duds filling the rest.
3. A winning word is randomly selected from the inserted words.
4. Each guess returns a match count (correct letters in the correct position).
5. The game ends when you guess the word or run out of lives.

## Design Notes

The UI is built on `tcell`, which provides a low‑level terminal drawing API. Because of that, cursor movement, selection, and redraw behavior are implemented manually.

## Architecture

```text
main.go (entry point + game loop)
├─ internal/cursor (cursor.go)
│  ├─ internal/renderer (renderer.go)
│  └─ internal/symbol (symbol.go, rune.go)
├─ internal/container (container.go, message-container.go)
│  ├─ internal/renderer (renderer.go)
│  ├─ internal/symbol (symbol.go, rune.go)
│  └─ internal/constants (constants.go)
├─ internal/validator (validator.go)
│  └─ internal/symbol (symbol.go, rune.go)
└─ internal/utilities (utilities.go)
   ├─ words/ (word list files)
   └─ internal/constants (constants.go)

Shared dependencies:
- github.com/gdamore/tcell (terminal UI)
- github.com/google/uuid (symbol ID generation)
```

### Package Responsibilities

| Package | File(s) | Purpose |
|---------|---------|---------|
| **main** | `main.go` | Entry point, initializes screen, runs game loop, handles events |
| **cursor** | `cursor.go` | Manages player cursor, navigation, and selection |
| **container** | `container.go`<br/>`message-container.go` | UI layout management, symbol grid storage, message display |
| **validator** | `validator.go` | Game logic, validates guesses, determines win/lose conditions |
| **renderer** | `renderer.go` | Low-level drawing functions for UI elements |
| **symbol** | `symbol.go`<br/>`rune.go` | Data structures for characters and words |
| **utilities** | `utilities.go` | Word loading, random generation, hex offset formatting |
| **constants** | `constants.go` | Configuration values, colors, styles, game settings |

### Data Flow

1. **Initialization**: `main.go` creates tcell screen and loads configuration
2. **Word Loading**: `utilities` loads word lists from `words/` directory
3. **Setup**: `container` creates UI containers and populates with symbols
4. **Game Logic**: `validator` randomly selects winning word
5. **Input Loop**: `cursor` handles user navigation via keyboard
6. **Rendering**: `renderer` draws all elements to screen
7. **Validation**: User presses Enter → `validator` checks guess → result displayed
8. **Loop**: Game continues until win, lose, or exit

## Configuration

Most gameplay and layout settings live in `internal/constants/constants.go`:

- `LIVES` (attempts)
- `DEFAULT_WORD_COUNT`
- `DEFAULT_WORD_LENGTH`
- Layout and spacing values (container sizes, offsets)

## Logging

Every run writes a timestamped log file to `./logs/`. See `LOGGING.md` for details and examples.

## Troubleshooting

- If the UI looks cramped, enlarge your terminal window.
- If colors look off, ensure your terminal supports true‑color ANSI rendering.
- If the game doesn’t start, run `go env` to confirm you’re using Go `1.22+`.

## Authors

-   Matt Maloney : matttm

## Contribute

If you want to contribute, open an issue or send a message with your idea.
