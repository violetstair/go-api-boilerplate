package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var LOGFORMAT = `{` +
	`"time":"${time_rfc3339_nano}",` +
	`"id":"${id}",` +
	`"remote_ip":"${remote_ip}",` +
	`"host":"${host}",` +
	`"method":"${method}",` +
	`"uri":"${uri}",` +
	`"status":${status},` +
	`"latency":${latency},` +
	`"latency_human":"${latency_human}",` +
	`"bytes_in":${bytes_in},` +
	`"bytes_out":${bytes_out}` +
	`}` +
	"\n"

var ALLOWORIGIN = "*"

func Application(e *echo.Echo) {
	e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: LOGFORMAT,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{ALLOWORIGIN},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.POST, echo.OPTIONS, echo.DELETE, echo.PUT},
	}))
}
