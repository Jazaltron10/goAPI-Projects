// server.go
package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/PunitNaran/weather_app/internal/cache"
	"github.com/PunitNaran/weather_app/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logging "github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus" // Import Logrus for structured logging
)

type Server struct {
	e *echo.Echo
	h *handler.Handler
	l *logrus.Logger
}

func NewServer(store cache.Cache, l *logrus.Logger) *Server {
	eRouter := echo.New()

	eRouter.Use(middleware.Logger())
	eRouter.Use(middleware.Recover())

	handle := &handler.Handler{}
	handle.CreateClient(store, l)

	eRouter.GET("/weather", handle.Weather)

	return &Server{
		e: eRouter,
		h: handle,
	}
}

func (s *Server) BeginServer(quit <-chan os.Signal) {
	s.e.Logger.SetLevel(logging.INFO)

	go func() {
		if err := s.e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			s.l.Fatal("shutting down the server")
		}
	}()

	<-quit
	s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		s.l.Fatal(err)
	}
}
