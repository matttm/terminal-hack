package container

import (
	"terminal_hack/internal/word"
)

type Container struct {
	rows    int
	columns int
	words   []word.Word
}

func NewContainer(rows, columns int) *Container {
	c := new(Container)
	c.rows = rows
	c.columns = columns
	return c
}

func (c *Container) InsertWord(word string) bool {
	return true
}

func (c *Container) IsFull() bool {
	return false
}
