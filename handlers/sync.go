package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/lwwset-crdt/lwwset"
)

// Sync merges multiple LWWSet present in a network to get them in sync
// It does so by obtaining the LWWSet from each node in the cluster
// and performs a merge operation with the local LWWSet
func Sync(LWWSet lwwset.LWWSet) (lwwset.LWWSet, error) {
	// Obtain addresses of peer nodes in the cluster
	peers := GetPeerList()

	// Return the local LWWSet back if no peers
	// are present along with an error
	if len(peers) == 0 {
		return LWWSet, errors.New("nil peers present")
	}

	// Iterate over the peer list and send a /lwwset/values GET request
	// to each peer to obtain its LWWSet
	for _, peer := range peers {
		peerLWWSet, err := SendListRequest(peer)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending lwwset values request")
			continue
		}

		// Merge the peer's LWWSet with our local LWWSet
		LWWSet = lwwset.Merge(LWWSet, peerLWWSet)
	}

	// DEBUG log in the case of success
	// indicating the new LWWSet
	log.WithFields(log.Fields{
		"set": LWWSet,
	}).Debug("successful lwwset sync")

	// Return the synced new LWWSet
	return LWWSet, nil
}

// SendListRequest is used to send a GET /lwwset/values
// to peer nodes in the cluster
func SendListRequest(peer string) (lwwset.LWWSet, error) {
	var _lwwset lwwset.LWWSet

	// Return an empty LWWSet followed by an error if the peer is nil
	if peer == "" {
		return _lwwset, errors.New("empty peer provided")
	}

	// Resolve the Peer ID and network to generate the request URL
	url := fmt.Sprintf("http://%s.%s/lwwset/values", peer, GetNetwork())
	response, err := SendRequest(url)
	if err != nil {
		return _lwwset, err
	}

	// Return an empty LWWSet followed by an error
	// if the peer's response is not HTTP 200 OK
	if response.StatusCode != http.StatusOK {
		return _lwwset, errors.New("received invalid http response status:" + fmt.Sprint(response.StatusCode))
	}

	// Decode the peer's LWWSet to be usable by our local LWWSet
	var lwwset lwwset.LWWSet
	err = json.NewDecoder(response.Body).Decode(&lwwset)
	if err != nil {
		return _lwwset, err
	}

	// Return the decoded peer's LWWSet
	_lwwset = lwwset
	return _lwwset, nil
}
