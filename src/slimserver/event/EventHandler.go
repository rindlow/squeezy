package event

// Manage player states, will most probably need two FSMs per player:
// * SlimProto state (e.g. play/pause/etc)
// * Menu state (e.g. IR messages)

import (
	"time"
        "github.com/op/go-logging"
	"slimserver/slimproto"
	"cgl.tideland.biz/state"
)

var eventLog = logging.MustGetLogger("event")

type SlimPlayerFSM struct {
	Fsm *state.FSM // The SlimProto FSM
	EventChan chan slimproto.ClientMessage	
	ActionChan chan slimproto.ServerMessage	
}

// The core FSM engine
func EventHandler(slimReg chan slimproto.SlimReg) {

        go func() {
                players := make([]SlimPlayerFSM, 1)


                for {
			eventLog.Info("FSM loop")
			time.Sleep(time.Second)

			// Iterate all known players, checking for incoming events
			for _, p := range players {
                        	select {
                                	case evt := <- p.EventChan:
						p.Fsm.Handle("ClientMessage", evt)
                                	default:
                        	}
			}

			// Check if there is a new player registered on the meta-chan, if so set it up
			select {
				case reg := <- slimReg:
				eventLog.Info("Received notification of new player, setting up chans and FSM")

				// Register the new player in the slice
				var p SlimPlayerFSM
				p.Fsm = state.New(NewSlimHandler(reg), 60*time.Second)
				p.EventChan=reg.EventChan
				p.ActionChan=reg.ActionChan
				players = append(players, p)

				p.Fsm.Handle("connect", nil)
  				default:
  			}


		}
	}()
}

