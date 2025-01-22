package constants

import (
	"github.com/gdamore/tcell"
)

var WORD_FG tcell.Color = tcell.NewRGBColor(137, 254, 135)
var WORD_BG tcell.Color = tcell.ColorDefault //  tcell.NewRGBColor(0, 36, 2)

var EMPTY_FG tcell.Color = tcell.NewRGBColor(0, 36, 2)
var EMPTY_BG tcell.Color = tcell.ColorDefault //  tcell.NewRGBColor(0, 36, 2)

const SELECTED_FG = tcell.ColorGreen

var SELECTED_BG = WORD_BG

var DUD_FG = WORD_FG

var DUD_BG = WORD_BG

func GetEmptyStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(EMPTY_FG).
		Background(EMPTY_BG)
}
func GetWordStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(WORD_FG).
		Background(WORD_BG)
}
func GetDudStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(DUD_FG).
		Background(DUD_BG)
}
func GetSelectedStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(SELECTED_FG).
		Background(SELECTED_BG)
}

const OFFSET = 3
const INSET = 1

const TEXT_PADDING = 2

const LIVES = 5
