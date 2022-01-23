package servers

import (
	"back/methods"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartMainServer() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/create_workspace", methods.ServerPostWorkspace())
	e.POST("/login", methods.ServerPostLogin())

	e.Start(":8080")
}
