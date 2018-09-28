package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/joho/godotenv"
	"github.com/violetstair/go-api-boilerplate/pkg/go-api-boilerplate"
)

func init() {
	if err := godotenv.Load("configs/env"); err != nil {
		os.Exit(1)
	}
}

func main() {
	e := echo.New()
	boilerplate.Application(e)
	boilerplate.Route(e)
	e.Logger.Fatal(e.Start(boilerplate.GetEnv("ADDRESS", ":8000")))
}