package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func call(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func saveRecord(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
