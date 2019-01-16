package main

import (
	"github.com/labstack/echo"
	"github.com/violetstair/go-api-boilerplate/pkg/go-api-boilerplate/boilerplate-api"
)

func main() {
	e := echo.New()
	boilerplate_api.Application(e)
	boilerplate_api.Route(e)
	e.Logger.Fatal(e.Start(":8000"))
}
