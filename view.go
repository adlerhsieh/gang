package main

import (
	tb "github.com/nsf/termbox-go"
)

type View struct {
	Data        interface{}
	Render      func()
	HandleEvent func(tb.Event)
	CursorIndex int
	State       string
}

func tbprint(x int, y int, msg string, fg tb.Attribute, bg tb.Attribute) {
	for _, c := range msg {
		tb.SetCell(x, y, c, fg, bg)
		x += 1
	}
}
