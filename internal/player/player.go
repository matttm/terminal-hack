package player

import (
	"terminal_hack/internal/cursor"

	"github.com/google/uuid"
)

type Player struct {
	Id     uuid.UUID
	Cursor *cursor.Cursor
	Lives  uint
}
func createPlayer(addr int, cursor *cursor.Cursor) *Player {
	id, _ := uuid.NewUUID()
	p := Player(
		id,
		cursor
	)
	return p
}
