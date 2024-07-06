package messages

import "terminal_hack/internal/player"

type AddPlayer struct {
	srcId  uint32
	dstId  uint32
	player player.Player
}
type PlayerRoster struct{}
