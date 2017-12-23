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
	viewTables.State = "connecting"
}

func (this *View) ViewTablesHandleEvent(event tb.Event) {
}

func (this *View) ViewTablesRender() {
	tb.Clear(dc, dc)

	rows, err := this.DB().Query("show databases;")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	vs := []string{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vs = append(vs, value)
		}
	}

	for i, v := range vs {
		tbprint(0, i, v, dc, dc)
	}

	tb.Flush()
}
