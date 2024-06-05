package renderer

import (
	// "fmt"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/symbol"

	"github.com/nsf/termbox-go"
)

const fg = termbox.ColorGreen
const bg = termbox.ColorBlack

const offset_x = 5
const offset_y = 5

func RenderRectangle(x1, y1, vpWidth, vpHeight int) {
	// cols := 3
	// colWidth := vpWidth / col
	// var vertChar rune = '|'

	const coldef = termbox.ColorDefault
	guiWidth := vpWidth - 2*offset_x
	guiHeight := vpHeight - 2*offset_y
	x2 := x1 + guiWidth
	y2 := y1 + guiHeight
	termbox.SetCell(x1, y1, '┌', coldef, coldef)
	termbox.SetCell(x1, y2, '└', coldef, coldef)
	termbox.SetCell(x2, y1, '┐', coldef, coldef)
	termbox.SetCell(x2, y2, '┘', coldef, coldef)
	fill(x1+1, y1, guiWidth-1, 1, termbox.Cell{Ch: '─'})
	fill(x1+1, y2, guiWidth-1, 1, termbox.Cell{Ch: '─'})
	fill(x1, y1+1, 1, guiHeight-1, termbox.Cell{Ch: '│'})
	fill(x2, y1+1, 1, guiHeight-1, termbox.Cell{Ch: '│'})

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
func RenderSymbolsInContainer(x1, y1, vpWidth, vpHeight int, symbols [][]*symbol.Symbol) {
	// seen := map[int]bool{}
	offset_x := x1 + constants.INSET
	offset_y := y1 + constants.INSET
	for _, symRow := range symbols {
		for _, sym := range symRow {
			//	if seen[sym.Id] {
			//		continue
			//	}
			//	seen[sym.Id] = true
			if sym == nil {
				continue
			}
			for _, _rune := range sym.Runes {
				termbox.SetCell(_rune.X+offset_x, _rune.Y+offset_y, _rune.Ch, fg, bg)
			}
		}
	}
}
func ColorRune(s *symbol.Symbol, fg, bg termbox.Attribute) {
	for _, r := range s.Runes {
		termbox.SetCell(r.X, r.Y, r.Ch, fg, bg)
		err := termbox.Flush()
		if err != nil {
			panic(err)
		}
	}

}
