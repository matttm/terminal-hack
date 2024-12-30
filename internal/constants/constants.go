package constants

import (
	"github.com/gdamore/tcell"
)

var WORD_FG tcell.Color = tcell.NewRGBColor(137, 254, 135)
var WORD_BG tcell.Color = tcell.NewRGBColor(0, 36, 4)

const SELECTED_FG = tcell.ColorGreen

var SELECTED_BG = WORD_BG

const DUD_FG = tcell.Color34

var DUD_BG = WORD_BG

const OFFSET = 3
const INSET = 1

const TEXT_PADDING = 2

const LIVES = 5
