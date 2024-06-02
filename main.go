package main

import (
	"terminal_hack/internal/container"
	"terminal_hack/internal/utilities"

	"github.com/nsf/termbox-go"
)

const fg = termbox.ColorGreen
const bg = termbox.ColorBlack

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	w, h := termbox.Size()
	words, _ := utilities.GetWordList(100)
	c := container.NewContainer(5, 5, h, w/3)
	for _, w := range words {
		c.InsertWord(w)
	}
	c.RenderContainer()
	c.RenderSymbols()

	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		}

	}
}
