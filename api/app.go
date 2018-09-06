package api

import (
	"net/http"

	"github.com/arxdsilva/olist/record"
	"github.com/labstack/echo"
)

func call(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func saveRecord(c echo.Context) (err error) {
	r := new(record.Record)
	err = c.Bind(r)
	if err != nil {
		return
	}
	err = r.DataChecks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}
