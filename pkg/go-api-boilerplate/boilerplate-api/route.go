package boilerplate_api

import (
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.GET("/", HelloRoot)
}
