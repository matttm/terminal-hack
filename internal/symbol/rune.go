package symbol

// Rune represents a single character with its position on the screen.
type Rune struct {
	X  int  // Screen X coordinate
	Y  int  // Screen Y coordinate
	Ch rune // The character itself
}
