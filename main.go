package main

import (
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/utilities"

	"time"

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
	words, _ := utilities.GetWordList(1000)
	c := container.NewContainer(5, 5, h, w/3)
	for _, w := range words {
		c.InsertWord(w)
	}
	c.RenderContainer()
	c.RenderSymbols()

	sym, err := c.GetSymbolAt(0)
	if err != nil {
		panic(err)
	}
	cursor := cursor.InitializeCursor(6, 6, sym)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				// fmt.Println("Blink")
				cursor.Blink()
			}
		}
	}()
	defer ticker.Stop()

	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeySpace:
				// cursor.Blink()
				break
			}
		}

	}
}
