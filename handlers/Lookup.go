package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// IsPresent is the JSON struct
// encapsulating the Lookup Response
type IsPresent struct {
	Present bool `json:"present"`
}

// Lookup is the HTTP handler used to return
// if a given value is present in the LWWSet node in the server
func Lookup(w http.ResponseWriter, r *http.Request) {
	var err error
	var present bool

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	// Sync the LWWSets if multiple nodes
	// are present in a cluster
	if len(GetPeerList()) != 0 {
		LWWSet, _ = Sync(LWWSet)
	}

	// Lookup given value in the LWWSet
	present, err = LWWSet.Lookup(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to lookup lwwset value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new LWWSet, the lookup value and if its present
	log.WithFields(log.Fields{
		"set":     LWWSet,
		"value":   value,
		"present": present,
	}).Debug("successful lwwset lookup")

	isPresent := IsPresent{present}

	JSONResponse, err := json.Marshal(isPresent)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to json marshall lookup lwwset value")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(JSONResponse)
	w.WriteHeader(http.StatusOK)
}
