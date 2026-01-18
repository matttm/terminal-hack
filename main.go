// Package main implements a terminal-based hacking mini-game inspired by Fallout's terminal hacking.
// The game presents the player with a grid of words and symbols where they must guess the correct password
// within a limited number of attempts.
package main

import (
	"fmt"
	"os"
	"strings"
	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/utilities"

	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

// main initializes and runs the terminal hacking game.
// It sets up the screen, creates the game grid with words and symbols,
// initializes the cursor and game state, and enters the main game loop
// to handle player input.
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	pad := 2
	hexCWidth := 8
	s.SetStyle(constants.GetEmptyStyle())
	s.Clear()

	w, h := s.Size()
	dy, dx := h-12, w/6
	shift := (w / 2) - (2*dx+pad+hexCWidth+pad)/2
	x1, y1 := shift, constants.OFFSET+4

	wordCount := 25
	wordLength := 4
	words, err := utilities.GetWordList(wordCount, wordLength)
	if err != nil {
		return fmt.Errorf("failed to get word list: %w", err)
	}

	for i, word := range words {
		words[i] = strings.ToLower(word)

	}
	totalChCount := dx * dy
	currentChCount := wordCount * wordLength
	neededChCnt := totalChCount - currentChCount
	words = append(words, utilities.GenerateRandomStrings(neededChCnt)...)
	hexOffsets := utilities.GenerateHexOffsets(dy, 2)

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	c := container.NewContainer(s, x1, y1, dy, dx)
	hexc := container.NewContainer(s, x1+dx+2, y1, dy, 8)
	out := container.CreateMessageContainer(s, x1+dx+2+8+2, y1, dy, dx)
	livesc := container.NewContainer(s, x1, y1-5, 4, 2*dx+pad+hexCWidth+pad)
	escc := container.NewContainer(s, x1, y1+dy+2, 1, 2*dx+pad+hexCWidth+pad)

	if err := c.InsertWords(words); err != nil {
		return fmt.Errorf("failed to insert words: %w", err)
	}
	if err := hexc.InsertWords(hexOffsets); err != nil {
		return fmt.Errorf("failed to insert hex offsets: %w", err)
	}
	if err := livesc.InsertWords([]string{}); err != nil {
		return fmt.Errorf("failed to initialize lives container: %w", err)
	}

	carnie := carnie.NewCarnie(c.GetSymbols())

	// c.RenderContainer()
	// offsetColumns.RenderContainer()
	// out.RenderContainer()
	// livesc.RenderContainer()
	// escc.RenderContainer()

	c.RenderSymbols()
	hexc.RenderSymbols()
	livesc.RenderSymbols()
	//
	sym, err := c.GetSymbolAt(0, 0)
	if err != nil {
		return fmt.Errorf("failed to get symbol: %w", err)
	}
	cursor := cursor.InitializeCursor(s, c, 0, 0, sym)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				// fmt.Println("Blink")
				cursor.Blink()
				s.Show()
			}
		}
	}()
	defer ticker.Stop()
	lives := constants.LIVES

mainloop:

	for {
		livesc.ClearContainer()
		livesc.WriteLineAtPosition(0, 1, "Robco Industries (TM) Termlink Protocol")
		livesc.WriteLineAtPosition(1, 1, "Enter Password Now")
		livesc.WriteLineAtPosition(3, 1, fmt.Sprintf("%d ATTEMPT(S) REMAINING", lives))
		escc.WriteLineAtPosition(0, 1, "Press ESC to exit")
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			cursor.ResetSymbol()
			switch ev.Key() {
			case tcell.KeyEscape:
				break mainloop
			case tcell.KeyBS:
				break
			case tcell.KeyUp:
				cursor.Displace(0, -1)
				break
			case tcell.KeyDown:
				cursor.Displace(0, 1)
				break
			case tcell.KeyLeft:
				cursor.Displace(-1, 0)
				break
			case tcell.KeyRight:
				cursor.Displace(1, 0)
				break
			case tcell.KeyEnter:
				isDone, msg := carnie.IsEnd(cursor.GetSelectedSymbol())
				if isDone {
					s.Clear()
					s.Sync()
					s.Fini()
					fmt.Println(msg)
					return nil
				}
				out.AddNewMessage(msg)
				lives -= 1
				break
			}
		}

	}
	s.Clear()
	s.Sync()
	s.Fini()
	return nil
}
