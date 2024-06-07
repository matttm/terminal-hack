package main

import (
	"fmt"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/utilities"

	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	w, h := termbox.Size()
	words, _ := utilities.GetWordList(100)
	words = append(words, utilities.GenerateRandomStrings(500)...)
	fmt.Print(words[0])
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	c := container.NewContainer(constants.OFFSET, constants.OFFSET, h, w/3)
	c.InsertWords(words)
	// c.RenderContainer()
	c.RenderSymbols()

	sym, err := c.GetSymbolAt(0, 0)
	if err != nil {
		panic(err)
	}
	cursor := cursor.InitializeCursor(c, 0, 0, sym)
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
			// x, y := cursor.X, cursor.Y
			cursor.ResetSymbol()
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeySpace:
				// cursor.Blink()
				break
			case termbox.KeyArrowUp:
				cursor.Displace(0, -1)
				break
			case termbox.KeyArrowDown:
				cursor.Displace(0, 1)
				break
			case termbox.KeyArrowLeft:
				cursor.Displace(-1, 0)
				break
			case termbox.KeyArrowRight:
				cursor.Displace(1, 0)
				break
			}
		}

	}
}
