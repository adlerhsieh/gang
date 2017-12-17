package main

import (
	// "fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	tbprint(0, 0, "localhost")
	tbprint(0, 1, "staging")
	tbprint(0, 2, "production")

	termbox.Flush()

	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if isExit(event) {
				os.Exit(0)
			}
		case termbox.EventError:
			panic(event.Err)
		}
	}
}

func isExit(event termbox.Event) bool {
	if event.Key == termbox.KeyEsc ||
		// q & Q
		event.Ch == 113 || event.Ch == 81 {
		return true
	}
	return false
}

func tbprint(x int, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x += 1
	}
}
