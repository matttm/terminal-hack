package symbol

import (
	"terminal_hack/internal/constants"

	"github.com/gdamore/tcell"
	"github.com/google/uuid"
)

type Symbol struct {
	Id    uuid.UUID
	Str   string
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
func (s *Symbol) FG() tcell.Color {
	if len(s.Runes) > 1 {
		return constants.WORD_FG
	} else {
		return constants.DUD_FG
	}
}
func (s *Symbol) BG() tcell.Color {
	if len(s.Runes) > 1 {
		return constants.WORD_BG
	} else {
		return constants.DUD_BG
	}
}
