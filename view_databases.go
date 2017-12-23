package main

import (
	tb "github.com/nsf/termbox-go"

	"database/sql"
)

func updateViewDatabases(db *sql.DB) {
	data := make(map[string]interface{})
	data["db"] = db

	viewDatabases.Data = data
	viewDatabases.HandleEvent = viewDatabases.ViewDatabasesHandleEvent
	viewDatabases.Render = viewDatabases.ViewDatabasesRender
	viewDatabases.State = "connecting"
	viewDatabases.CursorIndex = 0
}

func (this *View) ViewDatabasesHandleEvent(event tb.Event) {
	switch event.Ch {
	// j
	case 106:
		this.CursorIndex += 1
	// k
	case 107:
		this.CursorIndex -= 1
	}
}

func (this *View) getDatabases() {
	this.SaveQuery("databases", "show databases;")
}

func (this *View) ViewDatabasesRender() {
	tb.Clear(dc, dc)

	var xOffset int = 1
	var yOffset int = 3

	if this.State == "connecting" {
		this.getDatabases()
		this.State = "navigation"
	}

	tbprint(xOffset, 1, "Databases", dc, dc)

	databaseNames := this.Data["databases"].([]string)
	for i, databaseName := range databaseNames {
		if this.CursorIndex < 0 {
			this.CursorIndex = 0
		}
		if this.CursorIndex > len(databaseNames)-1 {
			this.CursorIndex = len(databaseNames) - 1
		}
		if i == this.CursorIndex {
			tbprint(xOffset, i+yOffset, "➜ "+databaseName, dc, 7)
		} else {
			tbprint(xOffset, i+yOffset, "➜ "+databaseName, dc, dc)
		}
	}

	tb.Flush()
}
