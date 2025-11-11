// Package symbol defines the Symbol type which represents a word or character
// in the terminal hacking game grid.
package symbol

import (
	"terminal_hack/internal/constants"

	"github.com/gdamore/tcell"
	"github.com/google/uuid"
)

// Symbol represents a word or single character in the game grid.
// Multi-rune symbols are selectable words, while single-rune symbols are "duds" (filler).
type Symbol struct {
	Id    uuid.UUID // Unique identifier for the symbol
	Str   string    // The complete string representation
	Runes []Rune    // Individual runes with their screen positions
}

// NewSymbol creates a new symbol with a unique ID and empty rune list.
func NewSymbol() *Symbol {
	s := new(Symbol)
	s.Id = uuid.New()
	// since runes coming from an api. they xhould alreasdy be on hesp
	s.Runes = []Rune{}
	return s
}

// IsGlyph returns true if this symbol is a single character (a "dud"), false if it's a word.
func (s *Symbol) IsGlyph() bool {
	return len(s.Runes) <= 1
}

// Length returns the number of runes in this symbol.
func (s *Symbol) Length() int {
	return len(s.Runes)
}

// InsertRune adds a rune to this symbol's rune list.
func (s *Symbol) InsertRune(r Rune) {
	s.Runes = append(s.Runes, r)
}

// FG returns the appropriate foreground color based on whether this is a word or dud.
func (s *Symbol) FG() tcell.Color {
	if len(s.Runes) > 1 {
		return constants.WORD_FG
	} else {
		return constants.DUD_FG
	}
}

// BG returns the appropriate background color based on whether this is a word or dud.
func (s *Symbol) BG() tcell.Color {
	if len(s.Runes) > 1 {
		return constants.WORD_BG
	} else {
		return constants.DUD_BG
	}
}
