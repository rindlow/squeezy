package slimproto

// The meta-channel tieing the EventHandler to the SlimServer (e.g. informing EventHandler about new players)
type SlimReg struct {
        // Chan for communicating an event from a player to the FSM
        EventChan chan ClientMessage

        // Chan for communicating an action from the FSM to the player
        ActionChan chan ServerMessage
}

