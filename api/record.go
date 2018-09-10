package api

import (
	"net/http"

	"github.com/arxdsilva/olist/record"
	"github.com/labstack/echo"
)

// SaveRecord implements the db insertion of data
// provided by external sources
// Responses:
// 201 Created
// 400 Bad Request
// 500 Internal Server Error
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
