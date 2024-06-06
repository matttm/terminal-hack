package symbol

import "github.com/google/uuid"

type Symbol struct {
	Id    uuid.UUID
	Runes []Rune
}

func NewSymbol() *Symbol {
	s := new(Symbol)
	s.Id = uuid.New()
	// since runes coming from an api. they xhould alreasdy be on hesp
	s.Runes = []Rune{}
	return s
}
func (s *Symbol) IsGlyph() bool {
	return len(s.Runes) <= 1
}
func (s *Symbol) Length() int {
	return len(s.Runes)
}
func (s *Symbol) InsertRune(r Rune) {
	s.Runes = append(s.Runes, r)
}
