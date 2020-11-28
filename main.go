package main

// The following implements the main Go
// package starting up the LWWSet server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/LWWSet-crdt/handlers"
)

const (
	// PORT is the LWWSet
	// server port
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started LWWSet node server")

	http.ListenAndServe(":"+PORT, r)
}
