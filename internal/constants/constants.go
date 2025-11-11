// Package constants defines color schemes, styling, and game configuration
// constants used throughout the terminal hacking game.
package constants

import (
	"github.com/gdamore/tcell"
)

// Color definitions for different UI elements

// WORD_FG is the foreground color for selectable words
var WORD_FG tcell.Color = tcell.NewRGBColor(137, 254, 135)

// WORD_BG is the background color for selectable words
var WORD_BG tcell.Color = tcell.ColorDefault //  tcell.NewRGBColor(0, 36, 2)

// EMPTY_FG is the foreground color for empty spaces
var EMPTY_FG tcell.Color = tcell.NewRGBColor(0, 36, 2)

// EMPTY_BG is the background color for empty spaces
var EMPTY_BG tcell.Color = tcell.ColorDefault //  tcell.NewRGBColor(0, 36, 2)

// SELECTED_FG is the foreground color for the currently selected character
const SELECTED_FG = tcell.ColorGreen

// SELECTED_BG is the background color for the currently selected character
var SELECTED_BG = WORD_BG

// DUD_FG is the foreground color for non-word characters (duds)
var DUD_FG = WORD_FG

// DUD_BG is the background color for non-word characters (duds)
var DUD_BG = WORD_BG

// GetEmptyStyle returns the tcell style for empty/background spaces.
func GetEmptyStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(EMPTY_FG).
		Background(EMPTY_BG)
}

// GetWordStyle returns the tcell style for selectable words.
func GetWordStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(WORD_FG).
		Background(WORD_BG)
}

// GetDudStyle returns the tcell style for non-word characters (duds).
func GetDudStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(DUD_FG).
		Background(DUD_BG)
}

// GetSelectedStyle returns the tcell style for the currently selected character.
func GetSelectedStyle() tcell.Style {
	var st tcell.Style
	return st.Foreground(SELECTED_FG).
		Background(SELECTED_BG)
}

// Layout and game configuration constants

// OFFSET is the vertical offset from the top of the screen for the game area
const OFFSET = 3

// INSET is the padding inside containers
const INSET = 1

// TEXT_PADDING is the padding around text elements
const TEXT_PADDING = 2

// LIVES is the number of attempts the player has to guess the correct word
const LIVES = 5
