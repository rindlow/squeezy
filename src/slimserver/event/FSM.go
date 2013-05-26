package event

import (
  "github.com/op/go-logging"
  "slimserver/slimproto"
  "cgl.tideland.biz/state"
)

// Setup logger
var fsmLog = logging.MustGetLogger("fsm")

// TBD: Add state data here, such as the menu FSM, current playlist etc
type SlimHandler struct {
  action chan slimproto.ServerMessage
}

// Constructor for the state handler type
func NewSlimHandler(reg slimproto.SlimReg) *SlimHandler {
  var sh SlimHandler
  sh.action = reg.ActionChan
  return &sh
}

// Setup the handler for a new player
func (sh *SlimHandler) Init() (*state.HandlerMap, string) {
  hm := state.NewHandlerMap(sh)
  hm.Assign("new", "HandleNew")
  hm.Assign("connected", "HandleConnected")
  hm.Assign("idle", "HandleIdle")
  return hm, "new"
}

func (sh *SlimHandler) Error(t *state.Transition, err error) string {
  fsmLog.Info("Handle error: %v", err)
  sh.Init()
  return "terminate"
}

func (sh *SlimHandler) Terminate() {
  fsmLog.Info("Terminating.")
}

