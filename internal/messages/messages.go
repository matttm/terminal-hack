package messages

import (
	"terminal_hack/internal/player"
)

type AddPlayer struct {
	SrcId  uint32
	DstId  uint32
	Player player.Player
}
type GameBoardResponse struct {
	// Starting with one container for now
	Containers int
	Words      []string
}
type GameBoardRequest struct {
	// NOTE: only hosts will send a response
	SrcId int
	DstId int
}

type PlayerMove struct {
	SrcId uint32
	DstId uint32
	X     int
	Y     int
}
type GameMessage struct {
	MessageType string
	Data        interface{}
}

const (
	GameMessageTopic      = "GAME_MESSAGE"
	AddPlayerType         = "ADD_PLAYER"
	PlayerMoveType        = "PLAYEER_MOVE"
	GameBoardRequestType  = "GAMEBOARD_STATE_REQUEST"
	GameBoardResponseType = "GAMEBOARD_STATE_RSPONSE"
)
