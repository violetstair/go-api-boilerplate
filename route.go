package main

import (
	"github.com/labstack/echo"

)

func route(e *echo.Echo) {
	e.GET("/", HelloRoot)
}
