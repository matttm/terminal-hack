package main

import (
	"os"

	"log/slog"
	"terminal_hack/internal/coordinator"
	"terminal_hack/internal/player"

	"github.com/nsf/termbox-go"
)

func main() {

	// Create a file handler
	f, err := os.Create("app.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a text handler that writes to the file
	handler := slog.NewTextHandler(f, nil)

	// Create a logger with the file handler
	logger := slog.New(handler)

	doneChan := make(chan bool)
	logger.Info("Initializing termbox")
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	p := player.CreatePlayer(1)
	logger.Info("Constructing player")
	coordinator := coordinator.Initialize(logger, 2, p, doneChan)
	termbox.Flush()
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// x, y := cursor.X, cursor.Y
			// coordinator.ResetSymbol()
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeySpace:
				// cursor.Blink()
				break
			case termbox.KeyArrowUp:
				coordinator.DisplaceLocal(0, -1)
				break
			case termbox.KeyArrowDown:
				coordinator.DisplaceLocal(0, 1)
				break
			case termbox.KeyArrowLeft:
				coordinator.DisplaceLocal(-1, 0)
				break
			case termbox.KeyArrowRight:
				coordinator.DisplaceLocal(1, 0)
				break
			case termbox.KeyEnter:
				coordinator.EvaluatePlayer()
				break
			}
		}

	}
	doneChan <- true
}
