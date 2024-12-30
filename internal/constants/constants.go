package constants

import (
	"github.com/gdamore/tcell"
)

const WORD_FG = tcell.Color40

var WORD_BG tcell.Color = tcell.NewRGBColor(0, 36, 4)

const SELECTED_FG = tcell.ColorGreen

var SELECTED_BG = WORD_BG

const DUD_FG = tcell.Color34

var DUD_BG = WORD_BG

const OFFSET = 3
const INSET = 1

const TEXT_PADDING = 2

const LIVES = 5
