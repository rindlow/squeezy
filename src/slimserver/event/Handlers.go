package event

import (
  "github.com/op/go-logging"
  "slimserver/slimproto"
  "cgl.tideland.biz/state"
)

// Setup logger
var handlerLog = logging.MustGetLogger("handlers")

// A new FSM, initialize and handle the "connect" action (i.e. TCP connect successful)
func (sh *SlimHandler) HandleNew(t *state.Transition) string {
  switch t.Command {
  case "connect":
    handlerLog.Info("A player Foo was connected")
    return "connected"
  }
  handlerLog.Info("Illegal command %q during state 'new'!", t.Command)
  return "new"
}

// A connected socket, waiting for a HELO
func (sh *SlimHandler) HandleConnected(t *state.Transition) string {
  cm, _ := t.Payload.(slimproto.ClientMessage)

  switch t.Command {
  case "ClientMessage":
     switch t := cm.(type) {
     case slimproto.MessageHELO :
       handlerLog.Info("Got a MessageHELO for %d", t.DeviceID)

       // Upon HELO we currently send a strms request for trackId 1
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
       sh.action <- msg

       return "idle"
    }
  }
  handlerLog.Info("Illegal command %q during state 'connected'!", t.Command)
  return "connected" // TBD: Should probably terminate here
}

// An idling client, accept messages
func (sh *SlimHandler) HandleIdle(t *state.Transition) string {
  cm, _ := t.Payload.(slimproto.ClientMessage)

  switch t.Command {
  case "ClientMessage":
    switch t := cm.(type) {
    case slimproto.MessageSTAT :
      handlerLog.Info("Got a MessageSTAT: %s (%d)", string(t.Event[:4]), t.ErrorCode)
      return "idle"
    }
  }

  handlerLog.Info("Illegal command %q during state 'idle'!", t.Command)
  return "idle"
}


