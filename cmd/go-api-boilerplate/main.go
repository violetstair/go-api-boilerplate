package main

import (
	"github.com/labstack/echo"
	api "github.com/violetstair/go-api-boilerplate/pkg/go-api-boilerplate/boilerplate-api"
)

func main() {
	e := echo.New()
	api.Application(e)
	api.Route(e)
	e.Logger.Fatal(e.Start(":8000"))
}
