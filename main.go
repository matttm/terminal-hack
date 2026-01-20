// Package main implements a terminal-based hacking mini-game inspired by Fallout's terminal hacking.
// The game presents the player with a grid of words and symbols where they must guess the correct password
// within a limited number of attempts.
package main

import (
	"fmt"
	"os"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/logger"
	"terminal_hack/internal/utilities"
	"terminal_hack/internal/validator"

	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

// main initializes and runs the terminal hacking game.
// It sets up the screen, creates the game grid with words and symbols,
// initializes the cursor and game state, and enters the main game loop
// to handle player input.
func main() {
	// Initialize logger
	if err := logger.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("Application starting")

	if err := run(); err != nil {
		logger.Error("Application error", "error", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	logger.Info("Application exiting normally")
}

func run() error {
	logger.Info("Initializing game")
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		logger.Error("Failed to create screen", "error", e)
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		logger.Error("Failed to initialize screen", "error", e)
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	logger.Info("Screen initialized successfully")
	pad := constants.CONTAINER_PADDING
	hexCWidth := constants.HEX_COLUMN_WIDTH
	s.SetStyle(constants.GetEmptyStyle())
	s.Clear()

	w, h := s.Size()
	dy, dx := h-constants.SCREEN_HEIGHT_REDUCTION, w/constants.SCREEN_WIDTH_DIVISOR
	shift := (w / 2) - (2*dx+pad+hexCWidth+pad)/2
	x1, y1 := shift, constants.OFFSET+constants.GAME_AREA_VERTICAL_OFFSET

	logger.Info("Screen dimensions", "width", w, "height", h)

	wordCount := constants.DEFAULT_WORD_COUNT
	wordLength := constants.DEFAULT_WORD_LENGTH
	logger.Info("Loading word list", "count", wordCount, "length", wordLength)
	words, err := utilities.GetWordList(wordCount, wordLength)
	if err != nil {
		logger.Error("Failed to load word list", "error", err)
		return fmt.Errorf("failed to get word list: %w", err)
	}
	logger.Debug("Words loaded", "count", len(words))
	// for i, word := range words {
	// 	words[i] = strings.ToLower(word)
	// }
	totalChCount := dx * dy
	currentChCount := wordCount * wordLength
	neededChCnt := totalChCount - currentChCount
	words = append(words, utilities.GenerateRandomStrings(neededChCnt)...)
	hexOffsets := utilities.GenerateHexOffsets(dy, constants.HEX_COLUMN_PADDING)

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	c := container.NewContainer(s, x1, y1, dy, dx)
	hexc := container.NewContainer(s, x1+dx+pad, y1, dy, constants.HEX_COLUMN_WIDTH)
	out := container.CreateMessageContainer(s, x1+dx+pad+constants.HEX_COLUMN_WIDTH+pad, y1, dy, dx)
	livesc := container.NewContainer(s, x1, y1-5, constants.LIVES_CONTAINER_HEIGHT, 2*dx+pad+hexCWidth+pad)
	escc := container.NewContainer(s, x1, y1+dy+constants.ESC_CONTAINER_VERTICAL_OFFSET, constants.ESC_CONTAINER_HEIGHT, 2*dx+pad+hexCWidth+pad)

	if err := c.InsertWords(words); err != nil {
		return fmt.Errorf("failed to insert words: %w", err)
	}
	if err := hexc.InsertWords(hexOffsets); err != nil {
		return fmt.Errorf("failed to insert hex offsets: %w", err)
	}
	if err := livesc.InsertWords([]string{}); err != nil {
		return fmt.Errorf("failed to initialize lives container: %w", err)
	}

	validator := validator.NewValidator(c.GetSymbols())

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
	ticker := time.NewTicker(constants.CURSOR_BLINK_INTERVAL_MS * time.Millisecond)
	go func() {
		for range ticker.C {
			cursor.Blink()
			s.Show()
		}
	}()
	defer ticker.Stop()
	lives := constants.LIVES

	logger.Info("Game initialized", "lives", lives, "totalWords", wordCount)

mainloop:

	for {
		livesc.ClearContainer()
		livesc.WriteLineAtPosition(constants.LIVES_TITLE_ROW, constants.UI_TEXT_COLUMN, "Robco Industries (TM) Termlink Protocol")
		livesc.WriteLineAtPosition(constants.LIVES_SUBTITLE_ROW, constants.UI_TEXT_COLUMN, "Enter Password Now")
		livesc.WriteLineAtPosition(constants.LIVES_COUNT_ROW, constants.UI_TEXT_COLUMN, fmt.Sprintf("%d ATTEMPT(S) REMAINING", lives))
		escc.WriteLineAtPosition(0, constants.UI_TEXT_COLUMN, "Press ESC to exit")
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			cursor.ResetSymbol()
			switch ev.Key() {
			case tcell.KeyEscape:
				logger.Info("User pressed ESC, exiting game")
				break mainloop
			case tcell.KeyBS:
				break
			case tcell.KeyUp:
				logger.Debug("Cursor moved up")
				cursor.Displace(0, -1)
				break
			case tcell.KeyDown:
				logger.Debug("Cursor moved down")
				cursor.Displace(0, 1)
				break
			case tcell.KeyLeft:
				logger.Debug("Cursor moved left")
				cursor.Displace(-1, 0)
				break
			case tcell.KeyRight:
				logger.Debug("Cursor moved right")
				cursor.Displace(1, 0)
				break
			case tcell.KeyEnter:
				selectedSym := cursor.GetSelectedSymbol()
				logger.Info("User selected symbol", "symbol", selectedSym.Str, "symbolId", selectedSym.Id)
				isDone, msg := validator.IsEnd(selectedSym)
				if isDone {
					logger.Info("Game ended", "message", msg)
					s.Clear()
					s.Sync()
					s.Fini()
					fmt.Println(msg)
					return nil
				}
				logger.Info("Selection result", "message", msg, "livesRemaining", lives-1)
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
