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
	var startDisc bool
	var startStream bool
	var updateLibrary bool
	var printLibrary bool
	var finalSleep int
	var libraryBase string

	// Parse command line flags
	flag.BoolVar(&startSlim, "slim", false, "Start the slimserver")
	flag.BoolVar(&startDisc, "disco", false, "Start the discovery server")
	flag.BoolVar(&startStream, "stream", false, "Start the streaming server")
	flag.BoolVar(&updateLibrary, "update", false, "Initiate a library update")
	flag.BoolVar(&printLibrary, "print", false, "Print library content on startup")
	flag.IntVar(&finalSleep, "sleep", 60, "Number of seconds to sleep before exit")
	flag.StringVar(&libraryBase, "base", "/data/music", "Basedir for mp3 files")
	flag.Parse()

	log.Println("Starting up...")

	if(startSlim) {
		log.Println("Starting slimserver");
		slimcommands := make(chan slimserver.SlimCommand)
		slimsrv := new(slimserver.SlimServer)
		go slimsrv.Serve(slimcommands)
	}

	if(startDisc) {
		go slimserver.DiscoveryServer()
	}

	if(startStream) {
		go slimserver.StreamServer()
	}

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


	if(printLibrary) {
		musiclibrary.PrintLibrary()
	}

	// Do stuff...
	if(finalSleep > 0) {
		log.Printf("Sleeping for %d seconds...\n", finalSleep)
		for i := 0; i < finalSleep; i++ {
			time.Sleep(time.Second)
		}
	}

	log.Println("Exiting...");
}
