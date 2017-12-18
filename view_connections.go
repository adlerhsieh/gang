package main

import (
	tb "github.com/nsf/termbox-go"
)

var listIndex int
var list []string = []string{"localhost", "staging", "production"}

func initViewConnections() View {
	return View{
		Data:        list,
		Render:      ViewConnectionsRender,
		HandleEvent: ViewConnectionsHandleEvent,
	}
}

func ViewConnectionsHandleEvent(event tb.Event) {
	switch event.Ch {
	// j
	case 106:
		listIndex += 1
	// k
	case 107:
		listIndex -= 1
	}
}

func ViewConnectionsRender() {
	var listXOffset int = 0
	var listYOffset int = 2

	tbprint(listXOffset, 0, "Databases", dc, dc)
	tbprint(listXOffset, 1, "---------", dc, dc)

	for i := 0; i < len(list); i++ {
		if listIndex < 0 {
			listIndex = 0
		}
		if listIndex > len(list)-1 {
			listIndex = len(list) - 1
		}
		if i == listIndex {
			tbprint(listXOffset, i+listYOffset, list[i], dc, tb.ColorGreen)
		} else {
			tbprint(listXOffset, i+listYOffset, list[i], dc, dc)
		}
	}
	tb.Flush()
}
