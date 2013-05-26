package main

import (
	"slimserver"
	"musiclibrary"
	"time"
	"flag"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func main() {

	var startSlim bool
	var updateLibrary bool
	var finalSleep int
	var libraryBase string

	// Setup logging levels
	logging.SetLevel(logging.WARNING, "disco")
	logging.SetLevel(logging.INFO, "event")
	logging.SetLevel(logging.INFO, "main")

	// Parse command line flags
	flag.BoolVar(&startSlim, "slim", true, "Start the slimserver")
	flag.BoolVar(&updateLibrary, "update", false, "Initiate a library update")
	flag.IntVar(&finalSleep, "sleep", 600, "Number of seconds to sleep before exit")
	flag.StringVar(&libraryBase, "base", "/data/music", "Basedir for mp3 files")
	flag.Parse()

	// Init bootstrap
	log.Info("Starting up...")

	// Should we start the server processes
	if(startSlim) {
		go slimserver.SlimServer()

	}

	// Should library be updated
	if(updateLibrary) {
		log.Info("Updating media library from " + libraryBase) 
		c := make(chan int)
		go func() {
			musiclibrary.UpdateLibrary(libraryBase)
			c <- 1
		}()

		// Wait for the update to complete before continuing
		<-c
		log.Info("Done with update")
	}


	// Final sleep, should be "forever"
	if(finalSleep > 0) {
		log.Info("Sleeping for %d seconds...\n", finalSleep)
		for i := 0; i < finalSleep; i++ {
			time.Sleep(time.Second)
		}
	}

	// Die hard
	log.Info("Exiting...");
}
