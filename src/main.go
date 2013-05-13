package main

import (
	//"fmt"
	//"log"
	//"slimserver"
	"musiclibrary"
	"time"
)

func main() {
	//slimcommands := make(chan slimserver.SlimCommand)
	//slimsrv := new(slimserver.SlimServer)

	// Create the library
	library := new(musiclibrary.MusicLibrary)

	// Initiate a library update
	c := make(chan int)
	go func() {
		musiclibrary.UpdateLibrary(library, "/data/music")
		c <- 1
	}()

	//go slimserver.DiscoveryServer()
	//go slimsrv.Serve(slimcommands)

	// Wait for the update to complete before continuing
	<-c

	// Print library statistics
        musiclibrary.PrintLibrary()

	// Do stuff...
	time.Sleep(10 * time.Second)
}
