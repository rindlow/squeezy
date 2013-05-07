package main

import (
	"fmt"
	//"log"
	"slimserver"
	"time"
)

func main() {
	//slimcommands := make(chan slimserver.SlimCommand)
	//slimsrv := new(slimserver.SlimServer)

	// Create the library
	library := new(slimserver.MusicLibrary)

	// Initiate a library update
	c := make(chan int)
	go func() {
		slimserver.UpdateLibrary(library, "/data/music")
		c <- 1
	}()

	//go slimserver.DiscoveryServer()
	//go slimsrv.Serve(slimcommands)

	// Wait for the update to complete before continuing
	<-c

	// Print library statistics
	fmt.Println("Found", library.NumDirs, "directories.")

	// For now print the library content
	for key, value := range library.Albums {
		fmt.Println(key, value)
	}

	// Do stuff...
	time.Sleep(10 * time.Second)
}
