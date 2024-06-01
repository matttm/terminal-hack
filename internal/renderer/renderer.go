package renderer

import (
	"github.com/nsf/termbox-go"
)
func renderGrid(vpWidth, vpHeight int) {
	// cols := 3
	// colWidth := vpWidth / col
	// var vertChar rune = '|'

	const coldef = termbox.ColorDefault
	offset := 5
	guiWidth := vpWidth - 2*offset
	guiHeight := vpHeight - 2*offset
	x1 := offset
	y1 := offset
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

	containers := 3
	containerWidth := guiWidth / containers
	fill(x1+1*containerWidth, y1, 1, guiHeight, termbox.Cell{Ch: '|'})
	fill(x1+2*containerWidth, y1, 1, guiHeight, termbox.Cell{Ch: '|'})
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

