package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Values is the HTTP handler to return the local LWWSet's values
// without syncing it with other nodes in a cluster
func Values(w http.ResponseWriter, r *http.Request) {
	// Get the local LWWSet values
	set := LWWSet

	// DEBUG log in the case of successful
	// list indicating the set
	log.WithFields(log.Fields{
		"set": set,
	}).Debug("successful lwwset values")

	// json encode response value
	json.NewEncoder(w).Encode(set)
}
