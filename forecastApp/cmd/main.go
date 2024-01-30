package main

import (
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

 
func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) 





	


}
