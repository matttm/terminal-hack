package container

import "terminal_hack/internal/symbol"

type Container struct {
	id int
	// this is the top left point of container
	x1           int
	y1           int
	rows         int
	columns      int
	startIndices []int // where the value is the starting position of that word (r*C + c)
	tracking     map[int]*symbol.Symbol
}

func NewContainer(x1, y1, rows, columns int) *Container {
	c := new(Container)
	c.x1 = x1
	c.y1 = y1
	c.rows = rows
	c.columns = columns
	c.startIndices = make([]int, 20)
	c.tracking = make(map[int]*symbol.Symbol)
	return c
}

func (c *Container) InsertWord(word string) bool {
	nextPosition := 0
	// if there is an element in map, get biggest index
	if len(c.startIndices) > 0 {
		lastPosition := c.startIndices[len(c.startIndices)-1] // starting position of last word
		lastSymbol := c.tracking[lastPosition]
		nextPosition = lastPosition + lastSymbol.Length()
	}
	c.startIndices = append(c.startIndices, nextPosition)
	c.tracking[nextPosition] = symbol.NewSymbol(word)
}

func (c *Container) IsFull() bool {
	return false
}
