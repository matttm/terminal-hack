package coordinator

import (
	"math/rand"

	"github.com/nsf/termbox-go"

	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/utilities"
)

type Coordinator struct {
	width      int
	height     int
	players    []int // slice of ids
	containers int
	carnie     *carnie.Carnie
	renderer   []*container.Container
}

func Initialize(_containers int) *Coordinator {
	c := new(Coordinator)
	c.ConstructBoard(_containers)
	return c
}
func (c *Coordinator) ConstructBoard(_containers int) {
	err := termbox.Init()
	c.containers = _containers
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	w, h := termbox.Size()
	c.width = w
	c.height = h

	words, _ := utilities.GetWordList(125)
	words = append(words, utilities.GenerateRandomStrings(500)...)

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	// subtract 1 so rooom is left for output
	for i := 0; i < c.containers-1; i++ {
	}
	_container := container.NewContainer(constants.OFFSET, constants.OFFSET, h-2*constants.OFFSET, w/3)
	out := container.NewContainer(2*constants.OFFSET+w/3, constants.OFFSET, h-2*constants.OFFSET, w/3)
	_container.InsertWords(words)
	carnie := carnie.NewCarnie(c.GetSymbols())

	c.RenderContainer()
	out.RenderContainer()
	c.RenderSymbols()
}
