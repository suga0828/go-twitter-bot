package server

import (
	"github.com/labstack/echo"
)

// Create a new echo server
func Create() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}
