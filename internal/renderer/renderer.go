package renderer

import (
	// "fmt"
	"terminal_hack/internal/symbol"

	"github.com/gdamore/tcell/termbox"
	"github.com/google/uuid"
)

const fg = termbox.ColorGreen
const bg = termbox.ColorBlack

const offset_x = 5
const offset_y = 5

func RenderRectangle(x1, y1, vpWidth, vpHeight int) {

	const coldef = termbox.ColorDefault
	guiWidth := vpWidth
	guiHeight := vpHeight
	x2 := x1 + guiWidth
	y2 := y1 + guiHeight
	fill(x1+1, y1, guiWidth, 1, termbox.Cell{Ch: '━'})
	fill(x1+1, y2+1, guiWidth, 1, termbox.Cell{Ch: '━'})
	fill(x1, y1+1, 1, guiHeight, termbox.Cell{Ch: '┃'})
	fill(x2+1, y1+1, 1, guiHeight, termbox.Cell{Ch: '┃'})
	termbox.SetCell(x1, y1, '┌', coldef, coldef)
	termbox.SetCell(x1, y2+1, '└', coldef, coldef)
	termbox.SetCell(x2+1, y1, '┐', coldef, coldef)
	termbox.SetCell(x2+1, y2+1, '┘', coldef, coldef)
	termbox.Flush()
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}
func ClearRectangle(x, y, w, h int) {
	fill(x, y, w, h, termbox.Cell{Ch: 'x', Fg: termbox.ColorBlack, Bg: termbox.ColorBlack})
}
func drawHorizontalSegment(x1, y1, w int, cell termbox.Cell) { fill(x1, y1, w, 1, cell) }
func drawVerticalSegment(x1, y1, h int, cell termbox.Cell)   { fill(x1, y1, 1, h, cell) }

// Function RenderSymbolsInContainer
// desc given a container with a symbols slice, render them to screen, bounded
//
//	by container bounds
func RenderSymbolsInContainer(x1, y1, vpWidth, vpHeight int, symbols [][]*symbol.Symbol) {
	seen := make(map[uuid.UUID]bool)
	for _, symRow := range symbols {
		for _, sym := range symRow {
			if sym == nil {
				continue
			}
			if seen[sym.Id] {
				continue
			}
			seen[sym.Id] = true
			for _, _rune := range sym.Runes {
				termbox.SetCell(_rune.X, _rune.Y, _rune.Ch, sym.FG(), sym.BG())
			}
		}
	}
}

// Function ColorRune
// desc change colors of given symbol
func ColorRune(s *symbol.Symbol, fg, bg termbox.Attribute) {
	for _, r := range s.Runes {
		termbox.SetCell(r.X, r.Y, r.Ch, fg, bg)
		err := termbox.Flush()
		if err != nil {
			panic(err)
		}
	}

}

// Function WriteLine
// desc writes a line of text at given y value w provided fg/bg
// returns top-left corner of text's bounding box
//
//	useful for determining position for next message
func WriteLine(_x, _y int, w, h int, s string, fg, bg termbox.Attribute) (int, int) {
	runes := []rune(s)
	x, y := _x, _y
	for _, r := range runes {
		termbox.SetCell(x, y, r, fg, bg)
		x++
		if x == _x+w-1 {
			x = _x
			y += 1
		}
	}
	return _x, _y
}
