package main

import (
	"time"
	// "unicode/utf8"

	// "github.com/mattn/go-runewidth"
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
	renderGrid(w, h)

	// go bgthread()

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

func bgthread() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		w, h := termbox.Size()
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				termbox.SetFg(x, y, termbox.ColorGreen)
			}
		}
		termbox.Flush()
		<-ticker.C // wait for ticker
	}
}
