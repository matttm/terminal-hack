package carnie

import (
	"fmt"
	"math/rand"
	"terminal_hack/internal/symbol"
)

type Carnie struct {
	lives       int
	winningWord *symbol.Symbol
}

func NewCarnie(symbols [][]*symbol.Symbol) *Carnie {
	c := new(Carnie)
	c.lives = 3
	c.winningWord = c.selectWinningWord(symbols)
	return c
}

func (c *Carnie) IsWinner(s *symbol.Symbol) bool {
	win := c.winningWord.Id == s.Id
	if win {
		panic("You won")
	}
	c.lives -= 1
	fmt.Println("You selected ", s.Str)
	if len(s.Str) <= 1 {
		return false
	}
	fraction := c.findCommonCharacters(s.Str)
	if c.lives == 0 {
		panic("Come back when you get some money buddy")
	}
	fmt.Println(fraction, " letters correct")
	return win
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
	for s == nil || (s != nil && len(s.Str) <= 1) {
		s = syms[rand.Intn(len(syms))][rand.Intn(len(syms[0]))]
	}
	return s
}
