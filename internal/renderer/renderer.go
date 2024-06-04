package renderer

import (
	// "fmt"
	"terminal_hack/internal/symbol"

	"github.com/nsf/termbox-go"
)

const fg = termbox.ColorGreen
const bg = termbox.ColorBlack

func RenderRectangle(x1, y1, vpWidth, vpHeight int) {
	// cols := 3
	// colWidth := vpWidth / col
	// var vertChar rune = '|'

	const coldef = termbox.ColorDefault
	offset := 5
	guiWidth := vpWidth - 2*offset
	guiHeight := vpHeight - 2*offset
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
func RenderSymbolsInContainer(x1, y1, vpWidth, vpHeight int, symbols map[int]*symbol.Symbol) {
	offset := 5
	cols := vpWidth - 2*offset - 1
	rows := vpHeight - 2*offset - 1
	offset_x := x1 + 1
	offset_y := y1 + 1
	// shrink row/col to be in container
	// rows -= 2
	// cols -= 2
	// TODO: move the startingInddices[] out of container
	for position, sym := range symbols {
		r := (position / cols) + offset_y
		c := (position % cols) + offset_x
		for _, _rune := range sym.Runes {
			// the following checks if the container is full
			if r-offset_y >= rows {
				_rune = 'X' // return
			}
			termbox.SetCell(c, r, _rune, fg, bg)
			c += 1
			// the below conditipn is c - 2 to remove offset
			if c-offset_x >= cols {
				c = offset_x
				r += 1
			}
		}
	}
}
func ColorRune(x, y int, s *symbol.Symbol, fg, bg termbox.Attribute) {
	// TODO: get the right starting coordintes
	// var cell termbox.Cell
	_x := x
	_y := y
	// fmt.Print(s.Runes)
	for _, r := range s.Runes {
		termbox.SetCell(_y, _x, rune(r), fg, bg)
		err := termbox.Flush()
		if err != nil {
			panic(err)
		}
		// TODO: just get one cell to flicker for now
		// TODO move this logic to func
		_y += 1
		// todo get # of cols
	}

}
