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
}

func (this *View) ViewDatabasesHandleEvent(event tb.Event) {
}

func (this *View) ViewDatabasesRender() {
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
