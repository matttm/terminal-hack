package carnie

import "terminal_hack/internal/symbol"

type Carnie struct {
	lives       int
	winningWord *symbol.Symbol
}

func NewCarnie(symbols [][]*symbol.Symbol) *Carnie {
	c := new(Carnie)
	c.lives = 3
	c.winningWord = symbols[0][0]
	return c
}

func (c *Carnie) IsWinner(s *symbol.Symbol) bool {
	win := c.winningWord.Id == s.Id
	if win {
		panic("You won")
	}
	c.lives -= 1
	if c.lives == 0 {
		panic("Come back when you get some money buddy")
	}
	return win
}
