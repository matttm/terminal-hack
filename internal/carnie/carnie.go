// Package carnie manages the game logic for the terminal hacking game.
// It handles word selection, guess validation, and game state tracking.
package carnie

import (
	"fmt"
	"math/rand"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/symbol"
)

// Carnie represents the game master that manages the hacking game state.
// It tracks the player's remaining lives and the correct winning word.
type Carnie struct {
	lives       int            // Number of remaining attempts
	winningWord *symbol.Symbol // The correct password to guess
}

// NewCarnie creates a new game master instance.
// It randomly selects a winning word from the provided symbol grid
// and initializes the player's lives from constants.
func NewCarnie(symbols [][]*symbol.Symbol) *Carnie {
	c := new(Carnie)
	c.lives = constants.LIVES
	c.winningWord = c.selectWinningWord(symbols)
	return c
}

// IsEnd checks if the game should end based on the player's selection.
// It returns true if the player won or ran out of lives, along with an appropriate message.
// For incorrect guesses, it returns false with feedback about matching characters and remaining lives.
func (c *Carnie) IsEnd(s *symbol.Symbol) (bool, string) {
	win := c.winningWord.Id == s.Id
	if win {
		return true, "You won"
	}
	c.lives -= 1
	if len(s.Str) <= 1 {
		return false, "Dud removed"
	}
	fraction := c.findCommonCharacters(s.Str)
	if c.lives == 0 {
		return true, fmt.Sprintf("Come back when you get some money buddy. Winning word is %s", c.winningWord.Str)
	}
	return false, fmt.Sprintf("%s letters correct. %d lives remaining", fraction, c.lives)
}

// findCommonCharacters compares the guessed word with the winning word
// and returns a fraction string showing how many characters match at the same positions.
func (c *Carnie) findCommonCharacters(s string) string {
	s1 := []rune(s)
	s2 := []rune(c.winningWord.Str)
	if len(s1) != len(s2) {
		return "error"
	}
	// assuming same len rn
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			count += 1
		}
	}
	return fmt.Sprintf("%d/%d", count, len(s1))
}

// selectWinningWord randomly selects a valid word (not a single character)
// from the symbol grid to be the winning password.
func (c *Carnie) selectWinningWord(syms [][]*symbol.Symbol) *symbol.Symbol {
	var s *symbol.Symbol = nil
	for s == nil || (s != nil && len(s.Runes) <= 1) {
		s = syms[rand.Intn(len(syms))][rand.Intn(len(syms[0]))]
	}
	return s
}
