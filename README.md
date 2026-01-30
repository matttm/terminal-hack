# terminal-hack

Development: Main development completed with intermittent improvements

## Description

This is a terminal based game written in Go, which is inspired by terminal hacking in Fallout 3. In this game, a 2d grid is filled with single 'dud' characters and words. The objective is to determine which word is the winning word in a limited number of guesses. Each incorrect guess will tell you how many of the selected letters are correct. A letter is correct if it has the same index and actual character as in the winning word. For example. if the winning word is `book` and selected word is `bang`, the selected word has one correct character.

Running this inside of `cool-retro-term` is suggested because it looks sick. Just search `cool-retro-term` on github, download the latest relase, build this from source, then run this in the new terminal.

Below is a run in `cool-retro-term`:

<img width="1728" alt="Screenshot 2025-01-22 at 5 16 23 PM" src="https://github.com/user-attachments/assets/5ad793ca-6147-46fb-abb5-56d03efec0a9" />

## Design

This game is created with a very low-level drawing library called Termbox. Because of this, I had to program all navigation/animation functionaliy.

## Architecture

```mermaid
graph TD
    Main[main.go<br/>Entry Point & Game Loop]
    
    Main --> Cursor[internal/cursor<br/>cursor.go]
    Main --> Container[internal/container<br/>container.go & message-container.go]
    Main --> Validator[internal/validator<br/>validator.go]
    Main --> Utilities[internal/utilities<br/>utilities.go]
    
    Cursor --> Renderer[internal/renderer<br/>renderer.go]
    Cursor --> Symbol[internal/symbol<br/>symbol.go & rune.go]
    
    Container --> Renderer
    Container --> Symbol
    Container --> Constants[internal/constants<br/>constants.go]
    
    Validator --> Symbol
    
    Utilities --> Words[words/<br/>Word List Files]
    Utilities --> Constants
    
    Renderer --> Symbol
    Renderer --> Constants
    
    Main -.-> Tcell[github.com/gdamore/tcell<br/>Terminal UI Library]
    Renderer -.-> Tcell
    Cursor -.-> Tcell
    
    Symbol -.-> UUID[github.com/google/uuid<br/>Symbol ID Generation]
    
    style Main fill:#4a90e2,stroke:#2e5c8a,color:#fff
    style Tcell fill:#90ee90,stroke:#228b22,color:#000
    style UUID fill:#90ee90,stroke:#228b22,color:#000
    style Words fill:#ffa07a,stroke:#ff6347,color:#000
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

## Authors

-   Matt Maloney : matttm

## Contribute

If you want to contribute, just send me a message.
