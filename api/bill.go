package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// Bill calculates a specific bill of a subscriber
// If a month is not specified, It'll be the last
// closed month
func (s *Server) Bill(c echo.Context) (err error) {
	number := c.Param("number")
	month := c.QueryParam("month")
	return c.JSON(http.StatusOK, number+month)
}
