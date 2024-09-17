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
	MessageType string
	Data        interface{}
}

const (
	GameMessageTopic = "GAME_MESSAGE"
	AddPlayerType    = "ADD_PLAYER"
	PlayerMoveType   = "PLAYEER_MOVE"
	GameBoardType    = "GAMEBOARD_STATE"
)
