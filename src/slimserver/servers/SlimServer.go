package servers

import (
	"encoding/binary"
	"net"
	"fmt"
	"slimserver/slimproto"
	"github.com/op/go-logging"
)

type SlimServer struct {
        Clients map[string]net.TCPConn
}

var slimLog = logging.MustGetLogger("slimproto")

func (*SlimServer) Serve(slimRegChan chan slimproto.SlimReg) {

	//var mac net.HardwareAddr
	slimLog.Info("Starting up listener for tcp 3483")
	listener, err := net.Listen("tcp", ":3483")
	if err != nil {
		slimLog.Panic("%s", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			slimLog.Panic("%s", err)
		}

		slimLog.Info("Got incomming connection from %s", conn.RemoteAddr())

		// A new player has connected, start by creating FSM-chans for it
                reg := new(slimproto.SlimReg)
                reg.EventChan = make(chan slimproto.ClientMessage, 100)
                reg.ActionChan = make(chan slimproto.ServerMessage, 100)

		// Send the reg object to the Eventhandler on the meta-chan
		slimRegChan <- *reg

		// Kick off a clientReader which reads from the player and creates FSM event
		go clientEventReader(conn, reg.EventChan)

		// Kick off a clientReader which reads FSM actions and writes to the player
		go clientActionSender(conn, reg.ActionChan)
	}

}

func clientActionSender(conn net.Conn, actions <-chan slimproto.ServerMessage) {
	slimLog.Info("Starting to listen for events for %s", conn.RemoteAddr())
	for {
		// Wait for an action
		action := <- actions

		// Must make a type assertion
                switch t := action.(type) {
                case slimproto.MessageStrm :
	                slimLog.Info("Got a MessageStrm of type %s", string(t.Command))

// TBD: This is just for testing... Need to wrap these properly
streamUrl :=  "GET /stream.mp3?track=1 HTTP/1.0\015\012\015\012"
binary.Write(conn, binary.BigEndian, int8(0))
binary.Write(conn, binary.BigEndian, int8(28+len(streamUrl)))
fmt.Fprintf(conn, "strm")

// Make the above way more generic, utilize ServerCommandName() method

			err := binary.Write(conn, binary.BigEndian, &t)
			if err != nil {
				slimLog.Error("FAILED to write message to player: %s", err)
				return
			}

			fmt.Fprintf(conn, streamUrl)

		default:
			slimLog.Warning("Got unknown action %s", t)
		}
	}

}

func clientEventReader(conn net.Conn, events chan<- slimproto.ClientMessage) {
	slimLog.Info("Starting to listen for actions for %s", conn.RemoteAddr())

	for {
	var header slimproto.MessageHeader

	err := binary.Read(conn, binary.BigEndian, &header)
	if err != nil {
		slimLog.Info("FAILED to read header from player: %s", err)
		return
	}
	cmd := string(header.Command[:4])
	switch cmd {
	case "HELO":
		if header.MsgLen != 36 {
			slimLog.Info("Expecting 36 bytes HELO, got %d\n",
				header.MsgLen)
			return
		}

		var msg slimproto.MessageHELO
 
		err = binary.Read(conn, binary.BigEndian, &msg)
		if err != nil {
			slimLog.Info("FAILED to read HELO: %s", err)
			return
		}


		// Send the event to the FSM	
slimLog.Debug("Sending HELO to event processor")
		events <- msg
		
	case "STAT":
		if header.MsgLen != 53 {
			slimLog.Info("Expecting 53 bytes STAT, got %d\n",
				header.MsgLen)
			return
		}
		var msg slimproto.MessageSTAT

		err = binary.Read(conn, binary.BigEndian, &msg)
		if err != nil {
			slimLog.Info("FAILED to read STAT: %s", err)
			return
		}
slimLog.Debug("Sending STAT to event processor")
		events <- msg
	}
	}
}