package main

import (
	"os"
	"os/signal"

	"github.com/PunitNaran/weather_app/internal/cache"
	"github.com/PunitNaran/weather_app/internal/server"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	cacheStore := cache.NewMemoryCache(log)
	server := server.NewServer(cacheStore, log)

	server.BeginServer(quit)
}
