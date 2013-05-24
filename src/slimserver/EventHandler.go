package slimserver

import (
	"time"
        "github.com/op/go-logging"
)

var eventLog = logging.MustGetLogger("event")

// TBD: Fill with data fields
type StreamAction struct {
}

// TBD: Fill with data fields
type StreamEvent struct {
}

// The pair of channels which tie the EventHandler to the StreamServer
type StreamServerFSMChans struct {
	StreamEvent	chan StreamEvent
	StreamAction	chan StreamAction
}

type SlimPlayerAction struct {
	msg Message // The message which is to be passed to the FSM
}

type SlimPlayerEvent struct {
	msg Message // The message which is to be passed to the FSM
}

// The meta-channel tieing the EventHandler to the SlimServer (e.g. informing EventHandler about new players)
type SlimReg struct {
	// Chan for communicating an event from a player to the FSM
	EventChan chan SlimPlayerEvent	

	// Chan for communicating an action from the FSM to the player
	ActionChan chan SlimPlayerAction	
}

type SlimPlayerFSM struct {
	State string // for now...
	EventChan chan SlimPlayerEvent	
	ActionChan chan SlimPlayerAction	
}

// The core FSM engine
func EventHandler(streamChans StreamServerFSMChans, slimReg chan SlimReg) {

        go func() {
                players := make([]SlimPlayerFSM, 1)


                for {
			eventLog.Info("FSM loop")
			time.Sleep(time.Second)

			// Iterate all known players, checking for incoming events
			for _, p := range players {
                        	select {
                                	case evt := <- p.EventChan:
					switch t := evt.msg.(type) {
					case MessageHELO :
						eventLog.Info("Got a MessageHELO with DeviceID %d", t.DeviceID)


// TBD: This is just bogus testing!
// Just for the fun of it... Tell the player to start streaming
var msg MessageStrm
msg.Command='s'
msg.Autostart='0'
msg.TransType='0'
msg.Format='m'
msg.ServerPort=9000
a := new(SlimPlayerAction)
a.msg=msg
p.ActionChan <- *a

					case MessageSTAT :
						eventLog.Info("Got a MessageSTAT")
					default:
						eventLog.Info("Type is default")
					}

// TBD: Pass the message to the associated FSM and return actions

                                	default:
                        	}
			}

			// Check for events from the stream server
			select {
				case evt := <- streamChans.StreamEvent:
				eventLog.Info("STREAM-EVT:: %s", evt)
				// TBD: Pass this event to the appropriate client FSM

  				default:
  			}

			// Check if there is a new player registered on the meta-chan, if so set it up
			select {
				case reg := <- slimReg:
				eventLog.Info("Received notification of new player, setting up chans and FSM")

				// Register the new player in the slice
				var p SlimPlayerFSM
				p.EventChan=reg.EventChan
				p.ActionChan=reg.ActionChan
				players = append(players, p)

  				default:
  			}


		}
	}()


}


// TBD: Implement the FSM helpers here

