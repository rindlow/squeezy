package main

import (
	"fmt"
	"slimserver"
	"musiclibrary"
	"time"
	"flag"
)

func main() {

	var startSlim bool
	var startDisc bool
	var updateLibrary bool
	var printLibrary bool

	// Parse command line flags
	flag.BoolVar(&startSlim, "slim", false, "Start the slimserver")
	flag.BoolVar(&startDisc, "disco", false, "Start the discovery server")
	flag.BoolVar(&updateLibrary, "update", true, "Initiate a library update")
	flag.BoolVar(&printLibrary, "print", true, "Print library content on startup")
	flag.Parse()

	if(startSlim) {
		fmt.Println("Starting slimserver");
		slimcommands := make(chan slimserver.SlimCommand)
		slimsrv := new(slimserver.SlimServer)
		go slimsrv.Serve(slimcommands)
	}

	if(startDisc) {
		go slimserver.DiscoveryServer()
	}

	if(updateLibrary) {
		fmt.Println("Updating media library")
		c := make(chan int)
		go func() {
			musiclibrary.UpdateLibrary("/data/music/Various")
			c <- 1
		}()

		// Wait for the update to complete before continuing
		<-c
		fmt.Println("Done with update")
	}


	if(printLibrary) {
		musiclibrary.PrintLibrary()
	}

	// Do stuff...
	fmt.Println("Sleeping for a while...")
	for i := 0; i < 10; i++ {
		fmt.Printf(".");
		time.Sleep(time.Second)
	}
	fmt.Printf("\n");

	fmt.Println("Exiting...");
}
