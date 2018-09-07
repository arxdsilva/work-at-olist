package api

import (
	"net/http"

	"github.com/arxdsilva/olist/record"
	"github.com/labstack/echo"
)

func (s *Server) SaveRecord(c echo.Context) (err error) {
	r := new(record.Record)
	err = c.Bind(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = r.DataChecks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if r.Type == "end" {
		id, err := s.Storage.UUIDFromStart(*r)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		r.ID = id
	}
	err = s.Storage.SaveRecord(*r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, nil)
}
