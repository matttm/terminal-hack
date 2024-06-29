package main

import (
	"terminal_hack/internal/carnie"
	"terminal_hack/internal/coordinator"
	cooordinator "terminal_hack/internal/cursor"

	"github.com/nsf/termbox-go"
)

func main() {
	coordinator := coordinator.Initialize(2)
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
				coordinator.Displace(0, -1)
				break
			case termbox.KeyArrowDown:
				coordinator.Displace(0, 1)
				break
			case termbox.KeyArrowLeft:
				coordinator.Displace(-1, 0)
				break
			case termbox.KeyArrowRight:
				coordinator.Displace(1, 0)
				break
			case termbox.KeyEnter:
				_, winStr := carnie.IsWinner(coordinator.GetSelectedSymbol())
				out.WriteLine(winStr)
				break
			}
		}

	}
}
