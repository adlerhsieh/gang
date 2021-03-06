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

func (this *View) SaveQuery(key string, statement string) {
	rows, err := this.DB().Query(statement)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	this.Data[key+"_columns"] = columns

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	square := [][]string{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var row []string
		var value string
		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			row = append(row, value)
		}

		square = append(square, row)
	}

	this.Data[key] = square
}

func flatten(square [][]string) []string {
	list := []string{}
	for _, slice := range square {
		for _, ele := range slice {
			list = append(list, ele)
		}
	}
	return list
}

func connectionString(username string, password string, database string) string {
	return username + ":" + password + "@/" + database
}
