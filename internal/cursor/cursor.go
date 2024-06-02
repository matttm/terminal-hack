package cursor

import "terminal_hack/internal/symbol"

type Cursor struct {
	X         int
	Y         int
	Selection symbol.Symbol
}

func InitializeCursor(x, y int, symbol *symbol.Symbol) *Cursor {
	c := new(Cursor)
	c.X = x
	c.Y = y
	return c
}
