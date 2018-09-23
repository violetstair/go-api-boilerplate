package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		os.Exit(1)
	}
}

func main() {
	e := echo.New()
	application(e)
	route(e)
	e.Logger.Fatal(e.Start(getEnv("ADDRESS", ":8000")))
}