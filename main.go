package main

import (
	"fmt"
	"terminal_hack/internal/coordinator"
	"terminal_hack/internal/player"

	"github.com/nsf/termbox-go"
)

func main() {
	fmt.Println("Initializing")
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	fmt.Println("Initializing termbox")
	p := player.CreatePlayer(1)
	fmt.Println("Constructing player")
	coordinator := coordinator.Initialize(2, p)
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
}
