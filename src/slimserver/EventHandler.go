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

// TBD: Fill with data fields
type SlimPlayerAction struct {
}

// TBD: Fill with data fields
type SlimPlayerEvent struct {
}

// The meta-channel tieing the EventHandler to the SlimServer (e.g. informing EventHandler about new players)
type SlimReg struct {
	// MAC of the new player
	Mac string	

	// Chan for communicating an event from a player to the FSM
	EventChan chan SlimPlayerEvent	

	// Chan for communicating an action from the FSM to the player
	ActionChan chan SlimPlayerAction	
}

// The core FSM engine
func EventHandler(streamChans StreamServerFSMChans, slimReg chan SlimReg) {

        go func() {
                for {
			eventLog.Debug("FSM loop")
			time.Sleep(time.Second)

			// Check for events (using select) from any of the player chans,
			// run FSM for these events. Emit actions if necessary.
			// TBD!!

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
				eventLog.Info("BINGO: %s", reg)
				// TBD: Create a FSM for this Mac, associate with the two chans

  				default:
  			}


		}
	}()


}


// TBD: Implement the FSM helpers here

