package symbol

type Symbol struct {
	Runes []rune
}

func NewSymbol(word string) *Symbol {
	s := new(Symbol)
	// since runes coming from an api. they xhould alreasdy be on hesp
	s.Runes = []rune(word)
	return s
}
func (s *Symbol) IsGlyph() bool {
	return len(s.Runes) <= 1
}
func (s *Symbol) Length() int {
	return len(s.Runes)
}
