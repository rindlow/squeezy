package main

import (
	//"fmt"
	//"log"
	"slimserver"
	"time"
)

func main() {
	slimcommands := make(chan slimserver.SlimCommand)
	slimsrv := new(slimserver.SlimServer)
	go slimserver.DiscoveryServer()
	go slimsrv.Serve(slimcommands)
	time.Sleep(10 * time.Second)
}
