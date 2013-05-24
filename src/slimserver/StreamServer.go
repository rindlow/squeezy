package slimserver

import (
        "net/http"
	"fmt"
	"os"
	"io"
        "github.com/op/go-logging"
)

var streamLog = logging.MustGetLogger("streamer")

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

	streamLog.Info("Player %s connected to streamer", playerId);

	// Check if there if a stream is configured for this player, if so start streaming

	fmt.Fprintf(w, "Here we should stream you some mp3 data...")

 	fd, _ := os.Open("/data/test.mp3")
	defer fd.Close()
	io.Copy(w, fd)

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
 
