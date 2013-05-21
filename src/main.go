package main

import (
	"log"
	"slimserver"
	"musiclibrary"
	"time"
	"flag"
)

func main() {

	var startSlim bool
	var updateLibrary bool
	var finalSleep int
	var libraryBase string

	// Parse command line flags
	flag.BoolVar(&startSlim, "slim", true, "Start the slimserver")
	flag.BoolVar(&updateLibrary, "update", false, "Initiate a library update")
	flag.IntVar(&finalSleep, "sleep", 60, "Number of seconds to sleep before exit")
	flag.StringVar(&libraryBase, "base", "/data/music", "Basedir for mp3 files")
	flag.Parse()

	// Init bootstrap
	log.Println("Starting up...")

	// Should we start the server processes
	if(startSlim) {

		// All the mess with creating channels should be encapsulated somewhere

		log.Println("Setting upp FSM chans...");

		// The StreamServer use two static chans, wrap them up
		streamChans := new(slimserver.StreamServerFSMChans)
		streamChans.StreamEvent = make(chan slimserver.StreamEvent)
		streamChans.StreamAction = make(chan slimserver.StreamAction)

		// The SlimServer allocates per-player chans, create meta-chan
		slimChans := make(chan slimserver.SlimRegChan)

		// Start Disco
		log.Println("Starting Discovery server...")
		go slimserver.DiscoveryServer()

		// Start Streamer
		log.Println("Starting Streaming server...")
		go slimserver.StreamServer(*streamChans)

		// Start SlimProto
		log.Println("Starting SlimProto server...")
		slimsrv := new(slimserver.SlimServer)
		go slimsrv.Serve(slimChans)

		// Start EventHandler
		log.Println("Starting EventHandler...")
		go slimserver.EventHandler(*streamChans, slimChans)
	}

	// Should library be updated
	if(updateLibrary) {
		log.Println("Updating media library from " + libraryBase) 
		c := make(chan int)
		go func() {
			musiclibrary.UpdateLibrary(libraryBase)
			c <- 1
		}()

		// Wait for the update to complete before continuing
		<-c
		log.Println("Done with update")
	}


	// Final sleep, should be "forever"
	if(finalSleep > 0) {
		log.Printf("Sleeping for %d seconds...\n", finalSleep)
		for i := 0; i < finalSleep; i++ {
			time.Sleep(time.Second)
		}
	}

	// Die hard
	log.Println("Exiting...");
}
