package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/seguins/georges-manager/api"
)

func initLogger() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	initLogger()

	server, err := api.NewServer()
	if err != nil {
		log.Errorf("Error during the server creation %s", err.Error())
		return
	}
	server.InitRouter()
	if err = server.ServeAPI(); err != nil {
		log.Errorf("Error during the server launching %s", err.Error())
	}
}
