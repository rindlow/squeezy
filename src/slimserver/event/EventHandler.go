package event

import (
	"time"
        "github.com/op/go-logging"
	"slimserver/slimproto"
//	"cgl.tideland.biz/state"
)

var eventLog = logging.MustGetLogger("event")

type SlimPlayerFSM struct {
	State string // for now...
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
					switch t := evt.(type) {
					case slimproto.MessageHELO :
						eventLog.Info("Got a MessageHELO with DeviceID %d", t.DeviceID)


// TBD: This is just bogus testing!
// Just for the fun of it... Tell the player to start streaming whenever it connects
var msg slimproto.MessageStrm
msg.Command='s'
msg.Autostart='1'
msg.PCMSampleSize='?'
msg.PCMSampleRate='?'
msg.PCMChannels='?'
msg.PCMEndian='?'
msg.TransType='0'
msg.Format='m'
msg.ServerPort=9000
p.ActionChan <- msg

					case slimproto.MessageSTAT :
						eventLog.Info("Got a MessageSTAT: %s (%d)", string(t.Event[:4]), t.ErrorCode)
					default:
						eventLog.Info("Type is default")
					}

// TBD: Pass the message to the associated FSM and return actions

                                	default:
                        	}
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

