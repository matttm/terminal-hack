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
	guiWidth := vpWidth - 2
	guiHeight := vpHeight - 2
	x1 := 1
	y1 := 1
	x2 := x1 + guiWidth
	y2 := y1 + guiHeight
	termbox.SetCell(x1, y1, '┌', coldef, coldef)
	termbox.SetCell(x1, y2, '└', coldef, coldef)
	termbox.SetCell(x2, y1, '┐', coldef, coldef)
	termbox.SetCell(x2, y2, '┘', coldef, coldef)
	fill(x1, y1, guiWidth, 1, termbox.Cell{Ch: '─'})
	fill(x1, y2, guiWidth, 1, termbox.Cell{Ch: '─'})
	fill(x1, y1, 1, guiHeight, termbox.Cell{Ch: '│'})
	fill(x2, y1, 1, guiHeight, termbox.Cell{Ch: '│'})
	termbox.Flush()
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}
func drawHorizontalSegment(x1, y1, w int, cell termbox.Cell) { fill(x1, y1, w, 1, cell) }
func drawVerticalSegment(x1, y1, h int, cell termbox.Cell)   { fill(x1, y1, 1, h, cell) }

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
