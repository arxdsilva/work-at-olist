package api

import (
	"fmt"
	"os"

	"github.com/arxdsilva/olist/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Storage storage.Storage
	Port    string
}

func New() *Server {
	return &Server{Port: port()}
}

func (s *Server) Listen() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/records", s.SaveRecord)
	e.GET("/bills/:subscriber", s.Bill)
	e.Logger.Fatal(e.Start(s.Port))
}

func port() (p string) {
	if p = os.Getenv("ADDRS_PORT"); p != "" {
		return fmt.Sprintf(":%s", p)
	}
	return ":8080"
}
