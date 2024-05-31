package symbol

import "golang.org/x/text/runes"

type Symbol struct {
	runes []rune
}

func NewSymbol(word string) *Symbol {
	s := new(Symbol)
	// since runes coming from an api. they xhould alreasdy be on hesp
	s.runes = []rune(word)
	return s
}
func (s *Symbol) IsGlyph() bool {
	return len(s.runes) <= 1
}
func (s *Symbol) Length() int {
	return len(s.runes)
}
