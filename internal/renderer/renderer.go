package renderer

import (
	"terminal_hack/internal/constants"
	"terminal_hack/internal/symbol"

	"github.com/gdamore/tcell"
	"github.com/google/uuid"
)

const offset_x = 5
const offset_y = 5

func RenderRectangle(s tcell.Screen, x1, y1, vpWidth, vpHeight int) {

	guiWidth := vpWidth
	guiHeight := vpHeight
	x2 := x1 + guiWidth
	y2 := y1 + guiHeight
	var st tcell.Style = constants.GetSelectedStyle()
	fill(s, x1, y1-1, guiWidth+1, 1, st, '━')
	fill(s, x1, y2+1, guiWidth+1, 1, st, '━')
	fill(s, x1-1, y1, 1, guiHeight+1, st, '┃')
	fill(s, x2+1, y1, 1, guiHeight+1, st, '┃')
	s.SetCell(x1-1, y1-1, st, '┌')
	s.SetCell(x1-1, y2+1, st, '└')
	s.SetCell(x2+1, y1-1, st, '┐')
	s.SetCell(x2+1, y2+1, st, '┘')
}

func fill(s tcell.Screen, x, y, w, h int, st tcell.Style, r rune) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			s.SetCell(x+lx, y+ly, st, r)
		}
	}
}
func ClearRectangle(s tcell.Screen, x, y, w, h int) {
	var st tcell.Style = constants.GetEmptyStyle()
	fill(s, x, y, w, h, st, 'x')
}

// func drawHorizontalSegment(x1, y1, w int, cell termbox.Cell) { fill(x1, y1, w, 1, cell) }
// func drawVerticalSegment(x1, y1, h int, cell termbox.Cell)   { fill(x1, y1, 1, h, cell) }

// Function RenderSymbolsInContainer
// desc given a container with a symbols slice, render them to screen, bounded
//
//	by container bounds
func RenderSymbolsInContainer(s tcell.Screen, x1, y1, vpWidth, vpHeight int, symbols [][]*symbol.Symbol) {
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
				var st tcell.Style
				st = st.
					Foreground(sym.FG()).
					Background(sym.BG())
				s.SetCell(_rune.X, _rune.Y, st, _rune.Ch)
			}
		}
	}
}

// Function ColorRune
// desc change colors of given symbol
func ColorRune(screen tcell.Screen, s *symbol.Symbol, fg, bg tcell.Color) {
	for _, r := range s.Runes {
		var st tcell.Style
		st = st.
			Foreground(fg).
			Background(bg)
		screen.SetCell(r.X, r.Y, st, r.Ch)
	}

}

// Function WriteLine
// desc writes a line of text at given y value w provided fg/bg
// returns top-left corner of text's bounding box
//
//	useful for determining position for next message
func WriteLine(screen tcell.Screen, _x, _y int, w, h int, s string, fg, bg tcell.Color) (int, int) {
	runes := []rune(s)
	x, y := _x, _y
	for _, r := range runes {
		var st tcell.Style
		st = st.
			Foreground(fg).
			Background(bg)
		screen.SetCell(x, y, st, r)
		x++
		if x == _x+w-1 {
			x = _x
			y += 1
		}
	}
	return x, y
}
