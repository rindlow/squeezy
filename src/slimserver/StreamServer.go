package slimserver

import (
        "net/http"
	"fmt"
	"log"
)

// A handler for generic web requests
func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is the generic web handler")
}

func streamHandler(w http.ResponseWriter, r *http.Request) {

	playerId:=r.FormValue("player")

	log.Printf("Player %s connected", playerId);

	// TBD: Setup a chan to the SlimServer informing it that
	//	a player is ready to receive data.

	fmt.Fprintf(w, "Here we should stream you some mp3 data...")
}

func StreamServer() {
	http.HandleFunc("/stream.mp3", streamHandler)
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(":9001", nil)
}
 
