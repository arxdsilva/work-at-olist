package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// Bill calculates a specific bill of a subscriber
// If a month is not specified, It'll be the last
// closed month
// month and year are strings of integers
func (s *Server) Bill(c echo.Context) (err error) {
	sub := c.Param("subscriber")
	month := c.QueryParam("month")
	year := c.QueryParam("year")
	return c.JSON(http.StatusOK, sub+month+year)
}
