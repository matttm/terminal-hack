package symbol

import "golang.org/x/text/runes"

type Symbol struct {
	runes   []rune
	isGlyph bool // is a valid word
}

func NewSymbol(runes []rune) *Symbol {
	s := new(Symbol)
	// since runes coming from an api. they xhould alreasdy be on hesp
	s.runes = runes
	s.isGlyph = len(runes) <= 1
	return s
}
