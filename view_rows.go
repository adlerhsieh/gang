package main

import (
	tb "github.com/nsf/termbox-go"

	"database/sql"
)

func updateViewRows(db *sql.DB, table string) {
	data := make(map[string]interface{})
	data["db"] = db
	data["table"] = table

	viewRows.Data = data
	viewRows.HandleEvent = viewRows.ViewRowsHandleEvent
	viewRows.Render = viewRows.ViewRowsRender
	viewRows.State = "loading"
	viewDatabases.CursorIndex = 0
}

func (this *View) currentTable() string {
	return this.Data["table"].(string)
}

func (this *View) ViewRowsHandleEvent(event tb.Event) {
}

func (this *View) ViewRowsRender() {
	tb.Clear(dc, dc)

	var xOffset int = 1
	var yOffset int = 3

	if this.State == "loading" {
		this.SaveQuery("rows", "select * from "+this.currentTable()+";")
		this.State = "navigation"
	}

	list := this.Data["rows"].([]string)
	for i, name := range list {
		if this.CursorIndex < 0 {
			this.CursorIndex = 0
		}
		if this.CursorIndex > len(list)-1 {
			this.CursorIndex = len(list) - 1
		}
		if i == this.CursorIndex {
			tbprint(xOffset, i+yOffset, "➜ "+name, dc, dc)
		} else {
			tbprint(xOffset, i+yOffset, "  "+name, dc, dc)
		}
	}

	tb.Flush()
}
