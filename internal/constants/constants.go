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

// Screen layout constants

// CONTAINER_PADDING is the padding between containers
const CONTAINER_PADDING = 2

// HEX_COLUMN_WIDTH is the width of the hexadecimal address column
const HEX_COLUMN_WIDTH = 8

// HEX_COLUMN_PADDING is the trailing padding for hex addresses
const HEX_COLUMN_PADDING = 2

// SCREEN_HEIGHT_REDUCTION reduces the available screen height for containers
const SCREEN_HEIGHT_REDUCTION = 12

// SCREEN_WIDTH_DIVISOR is used to calculate container width from screen width
const SCREEN_WIDTH_DIVISOR = 6

// GAME_AREA_VERTICAL_OFFSET is the additional vertical offset for the game area
const GAME_AREA_VERTICAL_OFFSET = 4

// LIVES_CONTAINER_HEIGHT is the height of the lives display container
const LIVES_CONTAINER_HEIGHT = 4

// ESC_CONTAINER_HEIGHT is the height of the escape instruction container
const ESC_CONTAINER_HEIGHT = 1

// ESC_CONTAINER_VERTICAL_OFFSET is the offset for the escape instruction container
const ESC_CONTAINER_VERTICAL_OFFSET = 2

// Game configuration constants

// DEFAULT_WORD_COUNT is the default number of words in the game
const DEFAULT_WORD_COUNT = 25

// DEFAULT_WORD_LENGTH is the default length of words in the game
const DEFAULT_WORD_LENGTH = 5

// HEX_ADDRESS_INCREMENT is the increment between hex addresses
const HEX_ADDRESS_INCREMENT = 8

// CURSOR_BLINK_INTERVAL_MS is the cursor blink interval in milliseconds
const CURSOR_BLINK_INTERVAL_MS = 500

// Renderer constants

// RENDERER_OFFSET_X is the X offset for rendering
const RENDERER_OFFSET_X = 5

// RENDERER_OFFSET_Y is the Y offset for rendering
const RENDERER_OFFSET_Y = 5

// UI Text positioning constants

// LIVES_TITLE_ROW is the row for the title in the lives container
const LIVES_TITLE_ROW = 0

// LIVES_SUBTITLE_ROW is the row for the subtitle in the lives container
const LIVES_SUBTITLE_ROW = 1

// LIVES_COUNT_ROW is the row for the remaining lives count
const LIVES_COUNT_ROW = 3

// UI_TEXT_COLUMN is the standard column for UI text
const UI_TEXT_COLUMN = 1
