package player

import (
	"terminal_hack/internal/cursor"

	"github.com/google/uuid"
)

type Player struct {
	id     uuid.UUID
	cursor cursor.Cursor
	lives  uint
}
