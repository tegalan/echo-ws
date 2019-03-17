package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Ping ...
func (a *App) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "Pong!"})
}
