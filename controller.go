package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloRoot(c echo.Context) error {
	serverver, err := GetServerVersion()
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, serverver)
}


