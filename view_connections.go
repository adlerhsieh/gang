package main

import (
	tb "github.com/nsf/termbox-go"
)

var connectionIndex int
var connections []string = []string{"localhost", "staging", "production"}

func initViewConnections() View {
	return View{
		Data:        connections,
		Render:      ViewConnectionsRender,
		HandleEvent: ViewConnectionsHandleEvent,
	}
}

func ViewConnectionsHandleEvent(event tb.Event) {
	switch event.Ch {
	// j
	case 106:
		connectionIndex += 1
	// k
	case 107:
		connectionIndex -= 1
	}
}

func ViewConnectionsRender() {
	var xOffset int = 0
	var yOffset int = 3

	tbprint(xOffset, 0, "-----------", dc, dc)
	tbprint(xOffset, 1, "Connections", dc, dc)
	tbprint(xOffset, 2, "-----------", dc, dc)

	for i := 0; i < len(connections); i++ {
		if connectionIndex < 0 {
			connectionIndex = 0
		}
		if connectionIndex > len(connections)-1 {
			connectionIndex = len(connections) - 1
		}
		if i == connectionIndex {
			tbprint(xOffset, i+yOffset, "➜ "+connections[i], dc, tb.ColorGreen)
		} else {
			tbprint(xOffset, i+yOffset, "➜ "+connections[i], dc, dc)
		}
	}
	tb.Flush()
}
