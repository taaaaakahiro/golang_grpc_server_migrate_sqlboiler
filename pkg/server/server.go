package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	s := &Server{}
	s.registerHandler()

	return s
}

func (s *Server) registerHandler() {
	s.Echo = echo.New()
	s.Echo.GET("/healthz", healthzHandler(s.Echo.AcquireContext()))
}

func healthzHandler(c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
}
