package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Remove is the HTTP handler used to remove
// values to the LWWSet node in the server
func Remove(w http.ResponseWriter, r *http.Request) {
	var err error

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	// Remove the given value to our stored LWWSet
	LWWSet, err = LWWSet.Removal(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to remove value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new LWWSet and the value removed
	log.WithFields(log.Fields{
		"set":   LWWSet,
		"value": value,
	}).Debug("successful lwwset removal")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
