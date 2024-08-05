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
type GameBoard struct {
	Symbols [][]symbol.Symbol
}
type PlayerMove struct {
	SrcId  uint32
	DstId  uint32
	Player player.Player
}
type GameMessage struct {
	MessageType uint32
	Data        interface{}
}

const (
	AddPlayerType  = iota
	PlayerMoveType = iota
	GameBoardType  = iota
)
