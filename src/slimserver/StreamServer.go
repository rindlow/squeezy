package slimserver

import (
        "net/http"
	"fmt"
	"log"
)

// Communication channels with the FSM

// A channel for passing events to the FSM
var streamEvent chan<- StreamEvent

// A channel for receiving actions from the FSM
var streamAction <-chan StreamAction

// A handler for generic web requests
func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is the generic web handler")
}

func streamHandler(w http.ResponseWriter, r *http.Request) {

	playerId:=r.FormValue("player")

	log.Printf("Player %s connected", playerId);

	// Check if there if a stream is configured for this player, if so start streaming

	fmt.Fprintf(w, "Here we should stream you some mp3 data...")
}

func StreamServer(chans StreamServerFSMChans) {

	// Fetch chan references from the param
	streamEvent=chans.StreamEvent
	streamAction=chans.StreamAction	

	// Setup the stream server
	http.HandleFunc("/stream.mp3", streamHandler)
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(":9000", nil)

	// TBD: Listen for actions on streamAction, push events to streamEvent

	// If a stream action is received, update some internal data structure so that
	// the streamHandler will know what to stream
}
 
