package container

import "terminal_hack/internal/symbol"

type Container struct {
	id int
	// this is the top left point of container
	x1 int
	y1 int
	rows    int
	columns int
	tracking map[int]Symbol
}

func NewContainer(x1, y1, rows, columns int) *Container {
	c := new(Container)
	c.x1 = x1
	c.y1 = y1
	c.rows = rows
	c.columns = columns
	c.tracking = make(msp[int]Symbol)
	return c
}

func (c *Container) InsertWord(word string) bool {
	initialPosition := 0;
	// if there is an element in map, get biggest index
	if len(tracking) > 0 {
	}
	c.tracking = symbol.NewSymbol(word)
}

func (c *Container) IsFull() bool {
	return false
}
