package slimserver

import (
        "net/http"
	"fmt"
)

// A handler for generic web requests
func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is the generic web handler")
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
    	fmt.Fprintf(w, "This is the stream handler for player %s", r.FormValue("player"))
}

func StreamServer() {
	http.HandleFunc("/stream.mp3", streamHandler)
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(":9001", nil)
}
 
