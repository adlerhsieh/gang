package main

import (
	tb "github.com/nsf/termbox-go"
)

func initViewConnections() View {
	connections := []map[string]string{
		map[string]string{
			"name":     "localhost",
			"host":     "localhost",
			"port":     "3306",
			"username": "root",
			"password": "",
		},
		map[string]string{
			"name":     "production",
			"host":     "localhost",
			"port":     "3306",
			"username": "root",
			"password": "",
		},
		map[string]string{
			"name":     "staging",
			"host":     "localhost",
			"port":     "3306",
			"username": "root",
			"password": "",
		},
	}

	data := make(map[string]interface{})
	data["connections"] = connections

	view := View{
		Data: data,
	}
	view.Render = view.ViewConnectionsRender
	view.HandleEvent = view.ViewConnectionsHandleEvent
	return view
}

func (this *View) ViewConnectionsHandleEvent(event tb.Event) {
	switch event.Ch {
	// j
	case 106:
		this.CursorIndex += 1
	// k
	case 107:
		this.CursorIndex -= 1
	}
}

func (this *View) ViewConnectionsRender() {
	var xOffset int = 0
	var yOffset int = 3

	tbprint(xOffset, 0, "-----------", dc, dc)
	tbprint(xOffset, 1, "Connections", dc, dc)
	tbprint(xOffset, 2, "-----------", dc, dc)

	connections := this.Data.(map[string]interface{})["connections"].([]map[string]string)

	for i := 0; i < len(connections); i++ {
		if this.CursorIndex < 0 {
			this.CursorIndex = 0
		}
		if this.CursorIndex > len(connections)-1 {
			this.CursorIndex = len(connections) - 1
		}
		if i == this.CursorIndex {
			tbprint(xOffset, i+yOffset, "➜ "+connections[i]["name"], dc, tb.ColorGreen)
		} else {
			tbprint(xOffset, i+yOffset, "➜ "+connections[i]["name"], dc, dc)
		}
	}
	tb.Flush()
}
