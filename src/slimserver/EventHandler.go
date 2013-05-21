package slimserver

import (
	"time"
	"fmt"
)

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
type SlimRegChan struct {
	// MAC of the new player
	Mac string	

	// Chan for communicating an event from a player to the FSM
	EventChan chan SlimPlayerEvent	

	// Chan for communicating an action from the FSM to the player
	ActionChan chan SlimPlayerAction	
}

// The core FSM engine
func EventHandler(streamChans StreamServerFSMChans, slimChan chan SlimRegChan) {

        go func() {
                for {
			fmt.Println("FSM loop")
			time.Sleep(time.Second)

			// Check for events from either stream server or one of the player chans,
			// run FSM for these events. Emit actions if necessary.

			// Check if there is a new player registered on the meta-chan, if so set it up

			// Iterate
		}
	}()


}



