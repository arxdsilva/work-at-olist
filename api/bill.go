package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) bill(c echo.Context) (err error) {
	number := c.Param("number")
	month := c.QueryParam("month")
	return c.JSON(http.StatusOK, number+month)
}
