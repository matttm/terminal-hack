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

func renderGrid(vpWidth, vpHeight int) {
	// cols := 3
	// colWidth := vpWidth / col
	// var vertChar rune = '|'

	const coldef = termbox.ColorDefault
	midy := vpHeight / 2
	midx := (vpWidth - 30) / 2
	termbox.SetCell(midx-1, midy, '│', coldef, coldef)
	termbox.SetCell(midx+vpWidth, midy, '│', coldef, coldef)
	termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
	termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
	termbox.SetCell(midx+vpWidth, midy-1, '┐', coldef, coldef)
	termbox.SetCell(midx+vpWidth, midy+1, '┘', coldef, coldef)
	fill(midx, midy-1, vpWidth, 1, termbox.Cell{Ch: '─'})
	fill(midx, midy+1, vpWidth, 1, termbox.Cell{Ch: '─'})
	termbox.Flush()
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
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
