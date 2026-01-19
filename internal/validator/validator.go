// Package validator manages the game logic for the terminal hacking game.
// It handles word selection, guess validation, and game state tracking.
package validator

import (
	"fmt"
	"math/rand"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/logger"
	"terminal_hack/internal/symbol"
)

// Validator represents the game master that manages the hacking game state.
// It tracks the player's remaining lives and the correct winning word.
type Validator struct {
	lives       int            // Number of remaining attempts
	winningWord *symbol.Symbol // The correct password to guess
}

// NewValidator creates a new game master instance.
// It randomly selects a winning word from the provided symbol grid
// and initializes the player's lives from constants.
func NewValidator(symbols [][]*symbol.Symbol) *Validator {
	c := new(Validator)
	c.lives = constants.LIVES
	c.winningWord = c.selectWinningWord(symbols)
	logger.Info("Validator initialized", "winningWord", c.winningWord.Str, "lives", c.lives)
	return c
}

// IsEnd checks if the game should end based on the player's selection.
// It returns true if the player won or ran out of lives, along with an appropriate message.
// For incorrect guesses, it returns false with feedback about matching characters and remaining lives.
func (c *Validator) IsEnd(s *symbol.Symbol) (bool, string) {
	logger.Debug("Checking selection", "selected", s.Str, "winning", c.winningWord.Str)
	win := c.winningWord.Id == s.Id
	if win {
		logger.Info("Player won!", "winningWord", c.winningWord.Str)
		return true, "Password accepted. Welcome back!"
	}
	c.lives -= 1
	if len(s.Str) <= 1 {
		logger.Debug("Dud removed", "remainingLives", c.lives)
		return false, "Dud removed"
	}
	fraction := c.findCommonCharacters(s.Str)
	logger.Info("Incorrect guess", "guessed", s.Str, "match", fraction, "remainingLives", c.lives)
	if c.lives == 0 {
		logger.Info("Player lost - out of lives", "winningWord", c.winningWord.Str)
		return true, fmt.Sprintf("Terminal locked. Winning word is %s", c.winningWord.Str)
	}
	return false, fmt.Sprintf("%s letters correct. %d lives remaining", fraction, c.lives)
}

// findCommonCharacters compares the guessed word with the winning word
// and returns a fraction string showing how many characters match at the same positions.
func (c *Validator) findCommonCharacters(s string) string {
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
func (c *Validator) selectWinningWord(syms [][]*symbol.Symbol) *symbol.Symbol {
	var s *symbol.Symbol = nil
	for s == nil || (s != nil && len(s.Runes) <= 1) {
		s = syms[rand.Intn(len(syms))][rand.Intn(len(syms[0]))]
	}
	return s
}
