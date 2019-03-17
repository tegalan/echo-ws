package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App ...
type App struct {
	Echo *echo.Echo
}

// Initialize ...
func (a *App) Initialize() {
	e := echo.New()
	a.Echo = e

	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	a.InitRouter()

}

// InitRouter ...
func (a *App) InitRouter() {
	e := a.Echo

	e.GET("/ping", a.Ping)
}

// Run the app
func (a *App) Run() {
	a.Echo.Logger.Fatal(a.Echo.Start(":8000"))
}
