package word

import "terminal_hack/internal/character"

type Word struct {
	runes []character.Character
	isGlyph bool  // is a valid word
}

func NewWord(runes []rune) *Word {
	word := new(Word)
	word.isGlyph = len(runes) <= 1
	for i, w :=range runes {
		
	}
	return word
}
