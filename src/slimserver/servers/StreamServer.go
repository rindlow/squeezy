package servers

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

// A handler for generic web requests
func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is the generic web handler")
}

func streamHandler(w http.ResponseWriter, r *http.Request) {

	trackId:=r.FormValue("track")

	if(trackId == "") {
		streamLog.Warning("trackId nil")
		return
	}

	streamLog.Info("Incoming stream request for track %s", trackId);

	id, err := strconv.Atoi(trackId)
	if(err != nil) {
		streamLog.Warning("trackID parse error: %s", err)
		return
	}

	// TBD: The GetTrackById should return a (track, error)
	fname:=musiclibrary.GetTrackById(id).FName
	if(fname == "") {
		streamLog.Warning("trackID not found")
		return
	}

	streamLog.Info("Starting to stream %s", fname)

 	fd, _ := os.Open(fname)
	defer fd.Close()
	io.Copy(w, fd)

}

func StreamServer() {
	http.HandleFunc("/stream.mp3", streamHandler)
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(":9000", nil)
}
 
