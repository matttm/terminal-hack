package container

import (
	"terminal_hack/internal/constants"
	"terminal_hack/internal/renderer"
	"terminal_hack/internal/symbol"
	// "terminal_hack/internal/utilities"
)

type Container struct {
	id int
	// this is the top left point of container
	x1      int
	y1      int
	rows    int
	columns int
	size    int // the amount of characters in this container (cannot exceed R*C)
	symbols [][]*symbol.Symbol
}

func NewContainer(x1, y1, rows, columns int) *Container {
	c := new(Container)
	c.x1 = x1
	c.y1 = y1
	c.rows = rows
	c.columns = columns
	c.symbols = make([][]*symbol.Symbol, rows)

	for i := 0; i < rows; i++ {
		c.symbols[i] = make([]*symbol.Symbol, columns)
	}
	c.size = 0
	return c
}
func (c *Container) InsertWords(words []string) {
	x := 0
	y := 0
	for _, w := range words {
		if y >= c.rows || x >= c.columns {
			return
		}
		x, y = c.InsertWord(x, y, w)
	}
}

func (c *Container) InsertWord(x, y int, word string) (int, int) {
	// check to see containrer has enough room for word
	offset_x := c.x1 + constants.INSET
	offset_y := c.y1 + constants.INSET
	if c.RemainingCapacity() < c.size+len(word) {
		return -1, -1
	}
	s := symbol.NewSymbol()
	for _, r := range []rune(word) {
		if y >= c.rows || x >= c.columns {
			panic("y is out-of-bounds")
		}
		c.symbols[y][x] = s
		s.InsertRune(symbol.Rune{X: x + offset_x, Y: y + offset_y, Ch: r})
		x += 1
		if x >= c.columns {
			x = 0
			y += 1
		}
	}
	// if there is an element in map, get biggest index
	c.size += len(word)
	return x, y
}
func (c *Container) RemainingCapacity() int {
	return c.rows*c.columns - c.size
}
func reset() {}
func (c *Container) RenderSymbols() {
	renderer.RenderSymbolsInContainer(c.x1, c.y1, c.columns, c.rows, c.symbols)
}
func (c *Container) RenderContainer() error {
	renderer.RenderRectangle(c.x1, c.y1, c.columns, c.rows)
	return nil
}

//	func (c *Container) findSymbolAtCoordinates(x, y int) (*symbol.Symbol, error) {
//		x, y = removeOffset(x, y)
//		encodedCoordinate := y*c.columns + x
//		index := utilities.binarySearch(c.startIndices, 0, len(c.startIndices), encodedCoordinate)
//		return c.tracking[index], nil
//	}
func removeOffset(x, y int) (int, int) {
	return x, y
}
func (c *Container) GetSymbolAt(x, y int) (*symbol.Symbol, error) {
	return c.symbols[y][x], nil
}
