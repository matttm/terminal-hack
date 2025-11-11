// Package cursor manages the player's cursor for navigating and selecting
// symbols in the terminal hacking game.
package cursor

import (
	"sync"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/renderer"
	"terminal_hack/internal/symbol"

	"github.com/gdamore/tcell"
)

// Cursor represents the player's selection cursor that can move around the container grid.
// It handles cursor movement, blinking animation, and symbol selection.
type Cursor struct {
	X           int                  // Current X position in the container grid
	Y           int                  // Current Y position in the container grid
	Selection   *symbol.Symbol       // Currently selected symbol
	blinkStatus bool                 // Current blink state (on/off)
	mu          sync.Mutex           // Mutex for thread-safe cursor operations
	container   *container.Container // Container the cursor operates within
	s           tcell.Screen         // Screen instance for rendering
}

// InitializeCursor creates a new cursor at the specified position with the given initial symbol.
func InitializeCursor(s tcell.Screen, container *container.Container, x, y int, symbol *symbol.Symbol) *Cursor {
	c := new(Cursor)
	c.container = container
	c.X = x
	c.Y = y
	c.Selection = symbol
	c.blinkStatus = false
	c.s = s
	return c
}

// Blink toggles the cursor's visual state between highlighted and normal.
// This is typically called on a timer to create a blinking effect.
func (c *Cursor) Blink() {
	c1, c2 := c.getBlinkStateColors()
	c.mu.Lock()
	renderer.ColorRune(c.s, c.Selection, c1, c2)
	c.mu.Unlock()
}

// ResetSymbol resets the currently selected symbol to its default colors.
func (c *Cursor) ResetSymbol() {
	renderer.ColorRune(c.s, c.Selection, c.Selection.FG(), c.Selection.BG())
}

// getBlinkStateColors toggles the blink state and returns the appropriate colors.
func (c *Cursor) getBlinkStateColors() (tcell.Color, tcell.Color) {
	c.blinkStatus = !c.blinkStatus
	c1 := constants.SELECTED_FG
	c2 := constants.SELECTED_BG
	if c.blinkStatus {
		return c1, c2
	} else {
		return c2, c1
	}
}

// Displace moves the cursor by the specified offset (x, y) with thread safety.
// Movement is bounded by the container and skips to the next different symbol.
func (c *Cursor) Displace(x, y int) {
	c.mu.Lock()
	c._displace(x, y)
	c.mu.Unlock()
}

// _displace is the internal implementation of cursor displacement.
// It recursively moves the cursor until it lands on a different symbol,
// ensuring the cursor doesn't stay on the same multi-character word.
func (c *Cursor) _displace(x, y int) {
	if !c.container.IsPointInContainer(c.X+x, c.Y+y) {
		return
	}
	c.X += x
	c.Y += y
	tmp, _ := c.container.GetSymbolAt(c.X, c.Y)
	if tmp == nil {
		return
	}
	if c.Selection.Id == tmp.Id {
		c._displace(x, y)
	}
	c.Selection = tmp
}

// GetSelectedSymbol returns the symbol currently under the cursor.
func (c *Cursor) GetSelectedSymbol() *symbol.Symbol {
	sym, _ := c.container.GetSymbolAt(c.X, c.Y)
	return sym
}
