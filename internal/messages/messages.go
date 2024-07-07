package messages

import (
	"terminal_hack/internal/player"
	"terminal_hack/internal/symbol"
)

type AddPlayer struct {
	SrcId  uint32
	DstId  uint32
	Player player.Player
}
type PlayerRoster struct {
	SrcId     uint32
	DstId     uint32
	Players   []player.Player
	GameBoard [][]symbol.Symbol
}
type PlayerMove struct {
	SrcId  uint32
	DstId  uint32
	Player player.Player
}
