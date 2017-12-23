package main

import (
	tb "github.com/nsf/termbox-go"

	"database/sql"
)

func updateViewTables(db *sql.DB) {
	data := make(map[string]interface{})
	data["db"] = db

	viewTables.Data = data
	viewTables.HandleEvent = viewTables.ViewTablesHandleEvent
	viewTables.Render = viewTables.ViewTablesRender
	viewTables.State = "loading"
	viewDatabases.CursorIndex = 0
}

func (this *View) ViewTablesHandleEvent(event tb.Event) {
	switch event.Ch {
	// j
	case 106:
		this.CursorIndex += 1
	// k
	case 107:
		this.CursorIndex -= 1
	}
}

func (this *View) ViewTablesRender() {
	tb.Clear(dc, dc)

	var xOffset int = 1
	var yOffset int = 3

	if this.State == "loading" {
		this.SaveQuery("tables", "show tables;")
		this.State = "navigation"
	}

	tbprint(xOffset, 1, "Tables", dc, dc)

	tableNames := this.Data["tables"].([]string)
	for i, tableName := range tableNames {
		if this.CursorIndex < 0 {
			this.CursorIndex = 0
		}
		if this.CursorIndex > len(tableNames)-1 {
			this.CursorIndex = len(tableNames) - 1
		}
		if i == this.CursorIndex {
			tbprint(xOffset, i+yOffset, "➜ "+tableName, dc, 7)
		} else {
			tbprint(xOffset, i+yOffset, "➜ "+tableName, dc, dc)
		}
	}

	tb.Flush()
}
