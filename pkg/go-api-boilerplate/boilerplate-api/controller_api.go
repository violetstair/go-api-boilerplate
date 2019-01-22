package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	boilerplate_structure "github.com/violetstair/go-api-boilerplate/pkg/go-api-boilerplate/boilerplate-structure"
	boilerplate_utiles "github.com/violetstair/go-api-boilerplate/pkg/go-api-boilerplate/boilerplate-utiles"
)

func HelloRoot(c echo.Context) error {
	var response boilerplate_structure.RESPONSE

	svr, err := boilerplate_utiles.GetServerVersion()
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = fmt.Sprintln(err)
		response.Items = svr
		return c.JSON(response.Status, response)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Items = svr
	return c.JSON(response.Status, response)
}
