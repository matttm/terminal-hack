package player

import (
	"terminal_hack/internal/constants"
	"terminal_hack/internal/cursor"

	"github.com/google/uuid"
)

type Player struct {
	Id     uuid.UUID
	Cursor *cursor.Cursor
	Lives  uint
}

func CreatePlayer(addr int) *Player {
	id, _ := uuid.NewUUID()
	p := new(Player)
	p.Id = id
	p.Lives = constants.LIVES
	return p
}
func (p *Player) clone() *Player {
	_p := Player{
		Id:     p.Id,
		Cursor: p.Cursor.clone(),
		Lives:  p.Lives,
	}
	return _p
}
