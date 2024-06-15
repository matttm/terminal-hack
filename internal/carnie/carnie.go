package carnie

import (
	"fmt"
	"math/rand"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/symbol"
)

type Carnie struct {
	lives       int
	winningWord *symbol.Symbol
}

func NewCarnie(symbols [][]*symbol.Symbol) *Carnie {
	c := new(Carnie)
	c.lives = constants.LIVES
	c.winningWord = c.selectWinningWord(symbols)
	return c
}

func (c *Carnie) IsWinner(s *symbol.Symbol) (bool, string) {
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
		return false, fmt.Sprintf("Come back when you get some money buddy. Winning word is %s", c.winningWord.Str)
	}
	return false, fmt.Sprintf("%s letters correct. %d lives remaining", fraction, c.lives)
}
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
func (c *Carnie) selectWinningWord(syms [][]*symbol.Symbol) *symbol.Symbol {
	var s *symbol.Symbol = nil
	for s == nil || (s != nil && len(s.Runes) <= 1) {
		s = syms[rand.Intn(len(syms))][rand.Intn(len(syms[0]))]
	}
	return s
}
