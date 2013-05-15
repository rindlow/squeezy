package main

import (
	"fmt"
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
	var finalSleep bool
	var libraryBase string

	// Parse command line flags
	flag.BoolVar(&startSlim, "slim", false, "Start the slimserver")
	flag.BoolVar(&startDisc, "disco", false, "Start the discovery server")
	flag.BoolVar(&startStream, "stream", false, "Start the streaming server")
	flag.BoolVar(&updateLibrary, "update", false, "Initiate a library update")
	flag.BoolVar(&printLibrary, "print", false, "Print library content on startup")
	flag.BoolVar(&finalSleep, "sleep", false, "Sleep 10 secs before exiting")
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
	if(finalSleep) {
		fmt.Println("Sleeping for a while...")
		for i := 0; i < 10; i++ {
			fmt.Printf(".");
			time.Sleep(time.Second)
		}
	fmt.Printf("\n");
	}

	log.Println("Exiting...");
}
