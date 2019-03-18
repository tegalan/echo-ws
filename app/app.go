package app

import (
	"echo-ws/ws"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App ...
type App struct {
	Echo *echo.Echo
	hub  *ws.Hub
}

// Initialize ...
func (a *App) Initialize() {
	e := echo.New()
	a.Echo = e

	e.HideBanner = true

	a.hub = ws.NewHub()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	a.InitRouter()

}

// InitRouter ...
func (a *App) InitRouter() {
	e := a.Echo

	e.File("/ws-client", "public/ws.html")
	e.GET("/ping", a.Ping)
	e.GET("/ws", a.WSHandler)
}

// Run the app
func (a *App) Run() {
	go a.hub.Run()
	a.Echo.Logger.Fatal(a.Echo.Start(":8000"))
}
