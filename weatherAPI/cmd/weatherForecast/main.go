package main

import (
	"os"
	"os/signal"

	"github.com/jazaltron10/goAPI/weatherAPI/api"
	"github.com/jazaltron10/goAPI/weatherAPI/internal/server"
	"github.com/labstack/echo/v4"
)

func init(){
	// Logging settings
	// logrus
}

func main() {
	// create channel 
	// to gracefully shutdown you application

	gcQuit := make(chan os.Signal, 1)
	signal.Notify(gcQuit, os.Interrupt)

	s :=server.SetupServer()
	s.BeginServer(gcQuit)

	// In the internal that has a file called server
	// server will Begin(gcQuit)
	// go routine -> start the server
	// <-gcQuit
	// gracefulShutdown()... context timeout ... 
	// shutdown your server
	/*
	e := echo.New()

	// Initialize routes
	api.InitializeRoutes(e)

	// Start the server on port 1323.
	e.Logger.Fatal(e.Start(":1323"))
	*/ 

}
