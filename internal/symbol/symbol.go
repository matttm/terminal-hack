package symbol

import "golang.org/x/text/runes"

type Symbol struct {
	runes   []rune
	isGlyph bool // is a valid word
}

func NewSymbol(word string) *Symbol {
	s := new(Symbol)
	// since runes coming from an api. they xhould alreasdy be on hesp
	s.runes = []rune(word)
	s.isGlyph = len(s.runes) <= 1
	return s
}
