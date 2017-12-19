package main

import (
	"os"

	tb "github.com/nsf/termbox-go"
)

var dc tb.Attribute = tb.ColorDefault

var (
	viewCurrent     View
	viewConnections View
)

func init() {
	viewConnections = initViewConnections()

	viewCurrent = viewConnections
}

func main() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	viewCurrent.Render()

	tb.SetInputMode(tb.InputEsc)
	tb.Clear(dc, dc)

	for {
		switch event := tb.PollEvent(); event.Type {
		case tb.EventError:
			panic(event.Err)
		case tb.EventKey:
			if isExit(event) {
				os.Exit(0)
			}
			viewCurrent.HandleEvent(event)
			viewCurrent.Render()
		}
	}
}

func isExit(event tb.Event) bool {
	if event.Key == tb.KeyEsc ||
		// q & Q
		event.Ch == 113 || event.Ch == 81 {
		return true
	}
	return false
}
