package slimserver

import (
        "net/http"
	"fmt"
	"os"
	"io"
	"strconv"
        "github.com/op/go-logging"
	"musiclibrary"
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

	trackId:=r.FormValue("track")

	streamLog.Info("Incoming stream request for track %s", trackId);

	// TBD: Verify that trackId is numeic
	id, _ := strconv.Atoi(trackId)
	fname:=musiclibrary.GetTrackById(id).FName

	streamLog.Info("Starting to stream %s", fname)

	// TBD: Error handling and such...

 	fd, _ := os.Open(fname)
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
 
