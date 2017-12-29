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
	switch event.Ch {
	// q
	case 113:
		viewCurrent = viewTables
	// j
	case 106:
		this.CursorIndex += 1
	// k
	case 107:
		this.CursorIndex -= 1
	}
}

func (this *View) ViewRowsRender() {
	tb.Clear(dc, dc)

	var xOffset int = 1
	var yOffset int = 3
	var columnWidth int = 23

	if this.State == "loading" {
		this.SaveQuery("rows", "select * from "+this.currentTable()+";")
		this.State = "navigation"
	}

	columnNames := this.Data["rows_columns"].([]string)
	for j, columnName := range columnNames {
		tbprint(xOffset+(j*columnWidth), yOffset, columnName, dc, dc)
	}

	rows := this.Data["rows"].([][]string)
	for i, row := range rows {
		if this.CursorIndex < 0 {
			this.CursorIndex = 0
		}
		if this.CursorIndex > len(rows)-1 {
			this.CursorIndex = len(rows) - 1
		}
		for j, value := range row {
			var msg string = value

			if len(value) > columnWidth-3 {
				msg = value[:columnWidth-5] + "..."
			}

			tbprint(xOffset+(j*columnWidth), i+2+yOffset, msg, dc, dc)
		}
	}

	tb.Flush()
}

func maxLength(list []string) int {
	var max int
	for i := 0; i < len(list); i++ {
		if len(list[i]) > max {
			max = len(list[i])
		}
	}
	return max
}
