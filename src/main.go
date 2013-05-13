package main

import (
	"fmt"
	//"log"
	//"slimserver"
	"musiclibrary"
	"time"
)

func main() {
	//slimcommands := make(chan slimserver.SlimCommand)
	//slimsrv := new(slimserver.SlimServer)

	// Create the library
	//library := new(musiclibrary.MusicLibrary)

	// Initiate a library update
        fmt.Println("Updating media library")
	c := make(chan int)
	go func() {
		musiclibrary.UpdateLibrary("/data/music")
		c <- 1
	}()

	//go slimserver.DiscoveryServer()
	//go slimsrv.Serve(slimcommands)

	// Wait for the update to complete before continuing
	<-c

        fmt.Println("Done with update")

	// Print library statistics
        musiclibrary.PrintLibrary()

	// Do stuff...
	time.Sleep(10 * time.Second)
}
