package main

import (
	"fmt"
	"os"
	"terminal_hack/internal/carnie"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/container"
	"terminal_hack/internal/cursor"
	"terminal_hack/internal/utilities"

	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
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

	s.SetStyle(constants.GetEmptyStyle())
	s.Clear()

	// quit := make(chan struct{})
	w, h := s.Size()
	shift := (w / 2) - (2*w/6+2+8+2)/2
	x1, y1, dy, dx := shift, constants.OFFSET, h-2*constants.OFFSET, w/6
	symbolCount := 25
	symbolLength := 4
	words, _ := utilities.GetWordList(symbolCount, symbolLength)

	totalChCount := dx * dy
	currentChCount := symbolCount * symbolLength
	neededChCnt := totalChCount - currentChCount
	words = append(words, utilities.GenerateRandomStrings(neededChCnt)...)
	hexOffsets := utilities.GenerateHexOffsets(dy, 2)

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	// TODO: put container and offset into a "hex-panel"

	c := container.NewContainer(s, x1, y1, dy, dx)
	hexc := container.NewContainer(s, x1+dx+2, y1, dy, 8)
	livesc := container.NewContainer(s, x1, y1-2, 2, dx)
	out := container.CreateMessageContainer(s, x1+dx+2+8+2, y1, dy, dx)

	c.InsertWords(words)
	hexc.InsertWords(hexOffsets)
	livesc.InsertWords([]string{})

	carnie := carnie.NewCarnie(c.GetSymbols())

	// c.RenderContainer()
	// offsetColumns.RenderContainer()
	out.RenderContainer()
	c.RenderSymbols()
	hexc.RenderSymbols()
	livesc.RenderSymbols()
	//
	sym, err := c.GetSymbolAt(0, 0)
	if err != nil {
		panic(err)
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
		livesc.WriteLineAtPosition(0, 1, fmt.Sprintf("%d ATTEMPT(S) REMAINING", lives))
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
				_, winStr := carnie.IsWinner(cursor.GetSelectedSymbol())
				out.AddNewMessage(winStr)
				lives -= 1
				break
			}
		}

	}
	s.Fini()
}
