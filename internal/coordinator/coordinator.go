package coordinator

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/nsf/termbox-go"

	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/player"
	"terminal_hack/internal/utilities"
)

type Coordinator struct {
	localPlayerUuid uuid.UUID
	width           int
	height          int
	players         map[uuid.UUID]*player.Player
	containersCount int
	carnie          *carnie.Carnie
	containers      []*container.Container
}

func Initialize(_containers int, player *player.Player) *Coordinator {
	c := new(Coordinator)
	c.localPlayerUuid = player.Id
	c.players = make(map[uuid.UUID]*player.Player)
	c.players[c.localPlayerUuid] = player
	c.ConstructBoard(_containers)
	return c
}
func (c *Coordinator) ConstructBoard(_containers int) {
	err := termbox.Init()
	c.containersCount = _containers
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	w, h := termbox.Size()
	c.width = w
	c.height = h

	words, _ := utilities.GetWordList(125)
	words = append(words, utilities.GenerateRandomStrings(500)...)

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	// subtract 1 so rooom is left for output
	for i := 0; i < c.containersCount-1; i++ {
	}
	i := 0
	_container := container.NewContainer(constants.OFFSET, constants.OFFSET, h-2*constants.OFFSET, w/3)
	out := container.NewContainer(2*constants.OFFSET+w/3, constants.OFFSET, h-2*constants.OFFSET, w/3)
	_container.InsertWords(words)
	c.carnie = carnie.NewCarnie(_container.GetSymbols())

	_container.RenderContainer()
	out.RenderContainer()
	_container.RenderSymbols()
	c.containers[i] = _container

	c.initializeCursor()
}

func (c *Coordinator) initializeCursor() {
	i := 0
	playerIdx := 0
	sym, err := c.containers[i].GetSymbolAt(0, 0)
	if err != nil {
		panic(err)
	}
	c.players[playerIdx].Cursor = cursor.InitializeCursor(c.containers[i], 0, 0, sym)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				// fmt.Println("Blink")
				c.players[playerIdx].Cursor.Blink()
			}
		}
	}()
	defer ticker.Stop()

	termbox.Flush()
}
func (c *Coordinator) DisplaceLocal(x, y int) {
	c.Displace(c.localPlayerUuid, x, y)
}

func (c *Coordinator) Displace(playerUuid uuid.UUID, x, y int) {
	c.players[playerUuid].Cursor.Displace(x, y)

}
