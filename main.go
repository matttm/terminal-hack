package main

import (
	"terminal_hack/internal/carnie"
	"terminal_hack/internal/coordinator"
	"terminal_hack/internal/player"

	"github.com/nsf/termbox-go"
)

func main() {
	p := player.CreatePlayer(1)
	coordinator := coordinator.Initialize(2, p)
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// x, y := cursor.X, cursor.Y
			coordinator.ResetSymbol()
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
				_, winStr := carnie.IsWinner(coordinator.GetSelectedSymbol())
				out.WriteLine(winStr)
				break
			}
		}

	}
}
