package main

import (
	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/utilities"

	"github.com/gdamore/tcell/termbox"
	"math/rand"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	w, h := termbox.Size()
	words, _ := utilities.GetWordList(125)
	words = append(words, utilities.GenerateRandomStrings(500)...)

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	// TODO: put container and offset into a "hex-panel"

	x1, y1, dy, dx := constants.OFFSET, constants.OFFSET, h-2*constants.OFFSET, w/6
	c := container.NewContainer(x1, y1, dx, dy)
	offsetColumns := container.NewContainer(x1+dx, y1, 8, dx)
	out := container.NewContainer(2*constants.OFFSET+w/3, constants.OFFSET, h-2*constants.OFFSET, w/3)
	c.InsertWords(words)
	carnie := carnie.NewCarnie(c.GetSymbols())

	c.RenderContainer()
	offsetColumns.RenderContainer()
	out.RenderContainer()
	c.RenderSymbols()
	//
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
			case termbox.KeyEnter:
				_, winStr := carnie.IsWinner(cursor.GetSelectedSymbol())
				out.WriteLine(winStr)
				break
			}
		}

	}
}
