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
	renderer.ColorRune(c.Selection, c.Selection.FG(), c.Selection.BG())
}
func (c *Cursor) getBlinkStateColors() (termbox.Attribute, termbox.Attribute) {
	c.blinkStatus = !c.blinkStatus
	c1 := constants.SELECTED_FG
	c2 := constants.SELECTED_BG
	if c.blinkStatus {
		return c1, c2
	} else {
		return c2, c1
	}
}
func (c *Cursor) Displace(x, y int) {
	c.mu.Lock()
	c._displace(x, y)
	c.mu.Unlock()
}
func (c *Cursor) _displace(x, y int) {
	if !c.container.IsPointInContainer(c.X+x, c.Y+y) {
		return
	}
	c.X += x
	c.Y += y
	tmp, _ := c.container.GetSymbolAt(c.X, c.Y)
	if c.Selection.Id == tmp.Id {
		c._displace(x, y)
	}
	c.Selection = tmp
}
func (c *Cursor) GetSelectedSymbol() *symbol.Symbol {
	sym, _ := c.container.GetSymbolAt(c.X, c.Y)
	return sym
}
