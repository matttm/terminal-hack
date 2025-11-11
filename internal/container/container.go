// Package container provides rectangular UI containers for organizing and displaying
// content in the terminal hacking game. Containers manage symbols, rendering, and text layout.
package container

import (
	"fmt"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/renderer"
	"terminal_hack/internal/symbol"

	"github.com/gdamore/tcell"
	// "terminal_hack/internal/utilities"
)

// Container represents a rectangular region on the screen that can hold and display symbols.
// It manages a 2D grid of symbols with automatic word wrapping and boundary checking.
type Container struct {
	id      int                // Unique identifier for the container
	x1      int                // Top-left X coordinate
	y1      int                // Top-left Y coordinate
	rows    int                // Number of rows in the container
	columns int                // Number of columns in the container
	size    int                // Current count of characters (cannot exceed rows * columns)
	symbols [][]*symbol.Symbol // 2D grid of symbols
	s       tcell.Screen       // Screen instance for rendering
}

// NewContainer creates a new container with the specified dimensions and position.
func NewContainer(s tcell.Screen, x1, y1, rows, columns int) *Container {
	c := new(Container)
	c.x1 = x1
	c.y1 = y1
	c.rows = rows
	c.columns = columns
	c.s = s
	c.symbols = make([][]*symbol.Symbol, rows)

	for i := 0; i < rows; i++ {
		c.symbols[i] = make([]*symbol.Symbol, columns)
	}
	c.size = 0
	return c
}

// InsertWords inserts multiple words into the container starting from the top-left position.
// Words are placed sequentially with automatic wrapping to the next row when a row is full.
func (c *Container) InsertWords(words []string) {
	x := 0
	y := 0
	for _, w := range words {
		if y >= c.rows || x >= c.columns {
			return
		}
		x, y = c.InsertWord(x, y, w)
		if x == -1 {
			fmt.Print(x, " ", y, "(", c.rows, " ", c.columns, " ", c.size, ")")
		}
	}
}

// InsertWord inserts a single word into the container at the specified position.
// It returns the next available position (x, y) after the word, or (-1, -1) if there's insufficient space.
func (c *Container) InsertWord(x, y int, word string) (int, int) {
	// check to see containrer has enough room for word
	offset_x := c.x1 + constants.INSET
	offset_y := c.y1 + constants.INSET
	if c.rows*c.columns < c.size+len(word) {
		return -1, -1
	}
	s := symbol.NewSymbol()
	s.Str = word
	for _, r := range []rune(word) {
		if y >= c.rows || x >= c.columns {
			panic("x/y is out-of-bounds")
		}
		c.symbols[y][x] = s
		s.InsertRune(symbol.Rune{X: x + offset_x, Y: y + offset_y, Ch: r})
		c.size += 1
		x += 1
		if x >= c.columns {
			x = 0
			y += 1
		}
	}
	return x, y
}

// RemainingCapacity returns the number of characters that can still be added to the container.
func (c *Container) RemainingCapacity() int {
	return c.rows*c.columns - c.size
}

// RenderSymbols renders all symbols in the container to the screen.
func (c *Container) RenderSymbols() {
	renderer.RenderSymbolsInContainer(c.s, c.x1, c.y1, c.columns, c.rows, c.symbols)
}

// RenderContainer renders the container's border/frame to the screen.
func (c *Container) RenderContainer() error {
	renderer.RenderRectangle(c.s, c.x1, c.y1, c.columns, c.rows)
	return nil
}

// GetSymbolAt retrieves the symbol at the specified grid coordinates within the container.
// Returns nil if the coordinates are out of bounds.
func (c *Container) GetSymbolAt(x, y int) (*symbol.Symbol, error) {
	if x > c.columns || y > c.rows {
		return nil, nil
	}
	return c.symbols[y][x], nil
}

// IsPointInContainer checks if the given coordinates are within the container's bounds.
func (c *Container) IsPointInContainer(x, y int) bool {
	return x >= 0 && y >= 0 && x < c.columns && y < c.rows
}

// GetSymbols returns the entire 2D grid of symbols in the container.
func (c *Container) GetSymbols() [][]*symbol.Symbol {
	return c.symbols
}

// WriteLineAtPosition writes text at the given position within the container.
// Returns the screen coordinates after writing the text.
func (c *Container) WriteLineAtPosition(pos, lines int, s string) (int, int) {
	y := c.y1 + 1 + pos
	x, y := renderer.WriteLine(c.s, c.x1, y, c.columns, c.rows, s, constants.DUD_FG, constants.DUD_BG)
	return x, y
}

// ClearContainer clears all content from the container.
func (c *Container) ClearContainer() {
	renderer.ClearRectangle(c.s, c.x1, c.y1, c.columns, c.rows)
}
