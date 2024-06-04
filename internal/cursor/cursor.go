package cursor

import (
	"terminal_hack/internal/renderer"
	"terminal_hack/internal/symbol"

	"github.com/nsf/termbox-go"
)

type Cursor struct {
	X           int
	Y           int
	Selection   *symbol.Symbol
	blinkStatus bool
}

func InitializeCursor(x, y int, symbol *symbol.Symbol) *Cursor {
	c := new(Cursor)
	c.X = x
	c.Y = y
	c.Selection = symbol
	c.blinkStatus = false
	return c
}
func (c *Cursor) Blink() {
	c1, c2 := c.getBlinkStateColors()
	renderer.ColorRune(c.X, c.Y, c.Selection, c1, c2)
}
func (c *Cursor) getBlinkStateColors() (termbox.Attribute, termbox.Attribute) {
	c.blinkStatus = !c.blinkStatus
	c1 := termbox.ColorGreen
	c2 := termbox.ColorBlack
	if c.blinkStatus {
		return c1, c2
	} else {
		return c2, c1
	}
}
