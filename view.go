package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	tb "github.com/nsf/termbox-go"
)

type View struct {
	Data        map[string]interface{}
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

func (this *View) DB() *sql.DB {
	return this.Data["db"].(*sql.DB)
}
