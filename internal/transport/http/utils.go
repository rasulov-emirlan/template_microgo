package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func respondErr(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"error": err.Error(),
	})
}
