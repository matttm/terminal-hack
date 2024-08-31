package coordinator

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"

	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/messages"
	"terminal_hack/internal/operator"
	"terminal_hack/internal/player"
	"terminal_hack/internal/symbol"
	"terminal_hack/internal/utilities"
)

// TODO:

type Coordinator struct {
	localPlayerUuid uint32
	width           int
	height          int
	players         map[uint32]*player.Player
	containersCount int
	carnie          *carnie.Carnie
	containers      []*container.Container
	doneChan        chan bool
	op              *operator.Operator
	SelfPlayerState chan interface{}
	logger          *slog.Logger
}

func Initialize(logger *slog.Logger, _containers int, _player *player.Player, done chan bool) *Coordinator {
	c := new(Coordinator)
	c.logger = logger
	c.localPlayerUuid = _player.Id.ID()
	c.doneChan = done
	c.players = make(map[uint32]*player.Player)
	c.players[c.localPlayerUuid] = _player
	c.SelfPlayerState = make(chan interface{})

	c.op = operator.New(c.logger, c.doneChan)
	c.op.InitializePubsub(_player)
	c.ConstructBoard(_containers)
	return c
}
func (c *Coordinator) ConstructBoard(_containers int) {
	c.containersCount = _containers
	c.containers = make([]*container.Container, _containers)
	w, h := termbox.Size()
	c.width = w
	c.height = h

	// subtract 1 so rooom is left for output
	for i := 0; i < c.containersCount-1; i++ {
		_container := container.NewContainer(constants.OFFSET, constants.OFFSET, h-2*constants.OFFSET, w/3)
		_container.RenderContainer()
		words, _ := utilities.GetWordList(125)
		words = append(words, utilities.GenerateRandomStrings(109)...)

		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		_container.InsertWords(words)
		_container.RenderSymbols()

		c.containers[i] = _container
	}
	out := container.NewContainer(2*constants.OFFSET+w/3, constants.OFFSET, h-2*constants.OFFSET, w/3)

	out.RenderContainer()
	c.carnie = carnie.NewCarnie(c.containers[0].GetSymbols())
	c.containers[_containers-1] = out

	c.initializeCursor(c.localPlayerUuid)
}

func (c *Coordinator) initializeCursor(id uint32) {
	c.logger.Info("Initialiing Cursor")
	playerId := id
	c.players[playerId].Cursor = cursor.InitializeCursor(c.containers[0])
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		slog.Info("Dispatch Blink thread")
		for {
			select {
			case <-ticker.C:
				slog.Debug("Blink")
				c.players[playerId].Cursor.Blink()
			case <-c.doneChan:
				ticker.Stop()
				return
			}
		}
	}()
	// defer ticker.Stop()
	termbox.Flush()
}
func (c *Coordinator) listenToPeers() {
	select {
	case msg := <-c.op.Messages:
		switch msg.GetTopic() {
		case "MESSAGE":
			bytes := msg.GetData()
			payload := new(messages.GameMessage)
			err := json.Unmarshal(bytes, payload)
			if err != nil {
				panic(err)
			}
			switch payload.MessageType {
			case messages.PlayerMoveType: // player position update
				var playerMove messages.PlayerMove = payload.Data.(messages.PlayerMove)
				player := playerMove.Player

				c.UpdatePlayer(player.Id.ID(), &player)
				break
			case messages.AddPlayerType:
				break
			case messages.GameBoardType:
				break
			}
			break
		}
	case <-c.doneChan:
		break

	}
}
func (c *Coordinator) DisplaceLocal(x, y int) {
	c.Displace(c.localPlayerUuid, x, y)
	c.SelfPlayerState <- messages.GameMessage{
		MessageType: "MESSAGE",
		Data: messages.PlayerMove{
			SrcId:  c.localPlayerUuid,
			DstId:  0,
			Player: *c.GetLocalPlayer().Clone(),
		},
	}
}

func (c *Coordinator) Displace(playerUuid uint32, x, y int) {
	cursor := c.players[playerUuid].Cursor
	cursor.ResetSymbol()
	cursor.Displace(x, y)
}
func (c *Coordinator) EvaluatePlayer() {
	_, winStr := c.carnie.IsWinner(c.players[c.localPlayerUuid].Cursor.GetSelectedSymbol())
	c.GetConsole().WriteLine(winStr)
}

func (c *Coordinator) GetConsole() *container.Container {
	return c.containers[c.containersCount-1] // console should always be last terminal
}
func (c *Coordinator) UpdatePlayer(id uint32, player *player.Player) {
	p := c.players[id]
	p.Cursor.ResetSymbol()
	p.Cursor.X = player.Cursor.X
	p.Cursor.Y = player.Cursor.Y
	p.Cursor.Selection = player.Cursor.Selection
}
func (c *Coordinator) getGameboard() [][][]symbol.Symbol {
	gameboard := make([][][]symbol.Symbol, c.containersCount)
	return gameboard
}
func (c *Coordinator) GetLocalPlayer() *player.Player {
	return c.players[c.localPlayerUuid]
}
