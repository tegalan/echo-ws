package app

import (
	"echo-ws/ws"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

// WSHandler ...
func (a *App) WSHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	client := &ws.Client{
		Hub:  a.hub,
		Conn: conn,
		Send: make(chan ws.Message, 256),
	}

	a.hub.Register <- client

	client.Listen()

	return nil
}
