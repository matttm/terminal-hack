// Package renderer provides low-level rendering functions for drawing UI elements
// to the terminal screen in the hacking game.
package renderer

import (
	"terminal_hack/internal/constants"
	"terminal_hack/internal/symbol"

	"github.com/gdamore/tcell"
	"github.com/google/uuid"
)

// RenderRectangle draws a bordered rectangle on the screen at the specified position and size.
// Uses Unicode box-drawing characters for the border.
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

// fill fills a rectangular area with the specified rune and style.
func fill(s tcell.Screen, x, y, w, h int, st tcell.Style, r rune) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			s.SetCell(x+lx, y+ly, st, r)
		}
	}
}

// ClearRectangle clears a rectangular area by filling it with the empty style.
func ClearRectangle(s tcell.Screen, x, y, w, h int) {
	var st tcell.Style = constants.GetEmptyStyle()
	fill(s, x, y, w, h, st, 'x')
}

// RenderSymbolsInContainer renders all symbols within a container to the screen.
// It efficiently tracks which symbols have been rendered to avoid duplicates,
// as symbols may span multiple grid cells.
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

// ColorRune changes the foreground and background colors of all runes in a symbol.
// Used for highlighting, selection, and cursor effects.
func ColorRune(screen tcell.Screen, s *symbol.Symbol, fg, bg tcell.Color) {
	for _, r := range s.Runes {
		var st tcell.Style
		st = st.
			Foreground(fg).
			Background(bg)
		screen.SetCell(r.X, r.Y, st, r.Ch)
	}

}

// WriteLine writes a line of text at the specified position with the given colors.
// Text automatically wraps to the next line when reaching the container width.
// Returns the position (x, y) after the last character written.
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
