package cursor

import (
	"sync"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/renderer"
	"terminal_hack/internal/symbol"

	"github.com/nsf/termbox-go"
)

type Cursor struct {
	X           int
	Y           int
	Selection   *symbol.Symbol
	blinkStatus bool
	mu          sync.Mutex
	container   *container.Container
}

func InitializeCursor(container *container.Container, x, y int, symbol *symbol.Symbol) *Cursor {
	c := new(Cursor)
	c.container = container
	c.X = x
	c.Y = y
	c.Selection = symbol
	c.blinkStatus = false
	return c
}
func (c *Cursor) Blink() {
	c1, c2 := c.getBlinkStateColors()
	c.mu.Lock()
	renderer.ColorRune(c.Selection, c1, c2)
	c.mu.Unlock()
}
func (c *Cursor) ResetSymbol() {
	renderer.ColorRune(c.Selection, constants.FG, constants.BG)
}
func (c *Cursor) getBlinkStateColors() (termbox.Attribute, termbox.Attribute) {
	c.blinkStatus = !c.blinkStatus
	c1 := constants.FG
	c2 := constants.BG
	if c.blinkStatus {
		return c1, c2
	} else {
		return c2, c1
	}
}
func (c *Cursor) Displace(x, y int) {
	c.mu.Lock()
	c.X += x
	c.Y += y
	c.Selection, _ = c.container.GetSymbolAt(c.X, c.Y)
	c.mu.Unlock()
}
