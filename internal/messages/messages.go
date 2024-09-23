package messages

import (
	"terminal_hack/internal/player"
)

type AddPlayer struct {
	Player player.Player
}
type GameBoardResponse struct {
	// Starting with one container for now
	Containers int
	Words      []string
	Players    []player.Player
}
type GameBoardRequest struct {
	// NOTE: only hosts will send a response
}

type PlayerMove struct {
	X int
	Y int
}
type GameMessage struct {
	MessageType string
	SrcId       uint32
	DstId       uint32
	Data        interface{}
}

const (
	GameMessageTopic      = "GAME_MESSAGE"
	AddPlayerType         = "ADD_PLAYER"
	PlayerMoveType        = "PLAYER_MOVE"
	GameBoardRequestType  = "GAMEBOARD_STATE_REQUEST"
	GameBoardResponseType = "GAMEBOARD_STATE_RSPONSE"
)
