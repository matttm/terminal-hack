package container

import (
	"terminal_hack/internal/renderer"
	"terminal_hack/internal/symbol"
	// "terminal_hack/internal/utilities"
)

type Container struct {
	id int
	// this is the top left point of container
	x1           int
	y1           int
	rows         int
	columns      int
	size         int   // the amount of characters in this container (cannot exceed R*C)
	startIndices []int // where the value is the starting position of that word (r*C + c)
	tracking     map[int]*symbol.Symbol
}

func NewContainer(x1, y1, rows, columns int) *Container {
	c := new(Container)
	c.x1 = x1
	c.y1 = y1
	c.rows = rows
	c.columns = columns
	c.size = 0
	c.startIndices = []int{}
	c.tracking = make(map[int]*symbol.Symbol)
	return c
}

func (c *Container) InsertWord(word string) bool {
	nextPosition := 0
	// check to see containrer has enough room for word
	if c.RemainingCapacity() < len(word) {
		return false
	}
	// if there is an element in map, get biggest index
	if c.size > 0 {
		lastPosition := c.startIndices[len(c.startIndices)-1] // starting position of last word
		lastSymbol := c.tracking[lastPosition]
		nextPosition = lastPosition + lastSymbol.Length()
	}
	c.startIndices = append(c.startIndices, nextPosition)
	c.tracking[nextPosition] = symbol.NewSymbol(word)
	c.size += len(word)
	return true
}
func (c *Container) RemainingCapacity() int {
	return c.rows*c.columns - c.size
}
func reset() {}
func (c *Container) RenderSymbols() {
	renderer.RenderSymbolsInContainer(c.x1, c.y1, c.columns, c.rows, c.tracking)
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
func (c *Container) GetSymbolAt(position int) (*symbol.Symbol, error) {
	if position == 0 {
		return c.tracking[0], nil
	}
	for _, i := range c.startIndices {
		if position > i {
			return c.tracking[i-1], nil
		}
	}
	return nil, nil
}
