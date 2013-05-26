package slimserver

import (
	"slimserver/slimproto"
	"slimserver/servers"
	"slimserver/event"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func SlimServer() {

  log.Info("Setting upp FSM chans...");

  // Create a metachannel for the slimserver and the event handler
  slimChans := make(chan slimproto.SlimReg, 100)

  // Start Discovery server
  log.Info("Starting Discovery server...")
  go servers.DiscoveryServer()

  // Start Stream server
  log.Info("Starting Streaming server...")
  go servers.StreamServer()

  // Start SlimProto server
  log.Info("Starting SlimProto server...")
  slimsrv := new(servers.SlimServer)
  go slimsrv.Serve(slimChans)

  // Start EventHandler
  log.Info("Starting EventHandler...")
  go event.EventHandler(slimChans)

}
 
