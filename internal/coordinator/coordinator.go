package coordinator

import (
	"log/slog"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"

	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/player"
	"terminal_hack/internal/utilities"
)

type Coordinator struct {
	localPlayerUuid uint32
	width           int
	height          int
	players         map[uint32]*player.Player
	containersCount int
	carnie          *carnie.Carnie
	containers      []*container.Container
	doneChan        chan bool
}

func Initialize(_containers int, _player *player.Player, done chan bool) *Coordinator {
	c := new(Coordinator)
	c.localPlayerUuid = _player.Id.ID()
	c.doneChan = done
	c.players = make(map[uint32]*player.Player)
	c.players[c.localPlayerUuid] = _player
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
		// _container.RenderContainer()
		words, _ := utilities.GetWordList(125)
		words = append(words, utilities.GenerateRandomStrings(109)...)

		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		_container.InsertWords(words)
		// _container.RenderSymbols()

		c.containers[i] = _container
	}
	// out := container.NewContainer(2*constants.OFFSET+w/3, constants.OFFSET, h-2*constants.OFFSET, w/3)

	// out.RenderContainer()
	c.carnie = carnie.NewCarnie(c.containers[0].GetSymbols())
	/// c.containers[_containers-1] = out

	c.initializeCursor(c.localPlayerUuid)
}

func (c *Coordinator) initializeCursor(id uint32) {
	slog.Info("Initialiing Cursor")
	playerId := id
	c.players[playerId].Cursor = cursor.InitializeCursor(c.containers[0])
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		slog.Info("Dispatch Blink thread")
		for {
			slog.Info("Inside of debug")
			select {
			case <-ticker.C:
				slog.Info("Blink")
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
func (c *Coordinator) DisplaceLocal(x, y int) {
	c.Displace(c.localPlayerUuid, x, y)
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
