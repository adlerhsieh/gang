package main

import (
	tb "github.com/nsf/termbox-go"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func initViewConnections() View {
	connections := []map[string]string{
		map[string]string{
			"name":     "localhost",
			"host":     "localhost",
			"port":     "3306",
			"username": "root",
			"password": "etufo32",
		},
		map[string]string{
			"name":     "production",
			"host":     "localhost",
			"port":     "3307",
			"username": "root",
			"password": "asdf1234",
		},
		map[string]string{
			"name":     "staging",
			"host":     "localhost",
			"port":     "3308",
			"username": "root",
			"password": "99889900",
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
	if event.Key == tb.KeyEnter {
		connection := this.CurrentConnection()

		connectionString := connection["username"] +
			":" +
			connection["password"] +
			"@/" +
			"gorm"

		db, err := sql.Open("mysql", connectionString)
		if err != nil {
			panic(err)
		}

		updateViewTables(db)
		viewCurrent = viewTables

		// data := this.Data.(map[string]interface{})
		// data["connection"] = conn
		// this.Data = data

		return
	}
	switch event.Ch {
	// j
	case 106:
		this.CursorIndex += 1
	// k
	case 107:
		this.CursorIndex -= 1
	}
}

func (this *View) GetConnections() []map[string]string {
	return this.Data["connections"].([]map[string]string)
}

func (this *View) CurrentConnection() map[string]string {
	connections := this.GetConnections()
	return connections[this.CursorIndex]
}

func (this *View) ViewConnectionsRender() {
	tb.Clear(dc, dc)

	var xOffset int = 1
	var yOffset int = 3

	tbprint(xOffset, 1, "⚡️ Quick Connection", dc, dc)

	connections := this.GetConnections()

	if len(connections) == 0 {
		connections = append(connections, map[string]string{
			"name": "none",
		})
	}

	for i := 0; i < len(connections); i++ {
		if this.CursorIndex < 0 {
			this.CursorIndex = 0
		}
		if this.CursorIndex > len(connections)-1 {
			this.CursorIndex = len(connections) - 1
		}
		if i == this.CursorIndex {
			tbprint(xOffset, i+yOffset, "➜ "+connections[i]["name"], dc, 7)
		} else {
			tbprint(xOffset, i+yOffset, "➜ "+connections[i]["name"], dc, dc)
		}
	}

	tbprint(31, 1, "🔎  Connection Details", dc, dc)
	currentConnection := connections[this.CursorIndex]
	tbprint(31, 3, "Host:     "+currentConnection["host"], dc, dc)
	tbprint(31, 4, "Post:     "+currentConnection["port"], dc, dc)
	tbprint(31, 5, "Username: "+currentConnection["username"], dc, dc)
	tbprint(31, 6, "Password: "+currentConnection["password"], dc, dc)

	tb.Flush()
}
