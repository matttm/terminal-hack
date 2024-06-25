package coordinator

import (
	"math/rand"
	"timd"

	"github.com/nsf/termbox-go"

	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/utilities"
)

type Coordinator struct {
	width           int
	height          int
	players         []int // slice of ids
	containersCount int
	carnie          *carnie.Carnie
	containers      []*container.Container
}

func Initialize(_containers int) *Coordinator {
	c := new(Coordinator)
	c.ConstructBoard(_containers)
	return c
}
func (c *Coordinator) ConstructBoard(_containers int) {
	err := termbox.Init()
	c.containersCount = _containers
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
	for i := 0; i < c.containersCount-1; i++ {
	}
	_container := container.NewContainer(constants.OFFSET, constants.OFFSET, h-2*constants.OFFSET, w/3)
	out := container.NewContainer(2*constants.OFFSET+w/3, constants.OFFSET, h-2*constants.OFFSET, w/3)
	_container.InsertWords(words)
	c.carnie := carnie.NewCarnie(_container.GetSymbols())

	_container.RenderContainer()
	out.RenderContainer()
	_container.RenderSymbols()
	c.containers[0] = _container

	c.initializeCursor()
}

func (c *Coordinator) initializeCursor() {
	sym, err := c.containers[0].GetSymbolAt(0, 0)
	if err != nil {
		panic(err)
	}
	c.cursor := cursor.InitializeCursor(c, 0, 0, sym)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				// fmt.Println("Blink")
				c.cursor.Blink()
			}
		}
	}()
	defer ticker.Stop()

	termbox.Flush()
}
