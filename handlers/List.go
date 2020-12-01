package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// List is the HTTP handler used to return
// all the values present in the LWWSet node in the server
func List(w http.ResponseWriter, r *http.Request) {
	// Sync the LWWSets if multiple nodes
	// are present in a cluster
	if len(GetPeerList()) != 0 {
		LWWSet, _ = Sync(LWWSet)
	}

	// Get the values from the LWWSet
	set := LWWSet.List()

	// DEBUG log in the case of success
	// indicating the new LWWSet
	log.WithFields(log.Fields{
		"set": set,
	}).Debug("successful lwwset list")

	// JSON encode response value
	json.NewEncoder(w).Encode(set)
}
