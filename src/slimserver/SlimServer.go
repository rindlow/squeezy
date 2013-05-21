package slimserver

import (
	//"bytes"
	"encoding/binary"
	"fmt"
	"net"
	//"time"
	"github.com/op/go-logging"
)

var slimLog = logging.MustGetLogger("slimproto")

type SlimCommand struct {
	Command byte
	Player  [6]byte
}

type Message interface {
	Command() string
}

type MessageHeader struct {
	Command [4]byte
	MsgLen  int32
}

type MessageHELO struct {
	DeviceID    byte
	Revision    byte
	Mac         [6]byte
	UUID        [16]byte
	ChannelList uint16
	Received    uint64
	Language    [2]byte
}

type MessageSTAT struct {
	Event                [4]byte
	CRLF                 byte
	MASInit              byte
	MASMode              byte
	BufferSize           uint32
	BufferFullness       uint32
	Received             uint64
	Wireless             uint16
	Jiffies              uint32
	OutputBufferSize     uint32
	OutputBufferFullness uint32
	ElapsedSeconds       uint32
	Voltage              uint16
	ElapsedMilliSeconds  uint32
	TimeStamp            uint32
	ErrorCode            uint16
}

type MessageAude struct {
	SPDIFEnable byte
	DACEnable   byte
}

type MessageAudg struct {
	OldLeft  uint32
	OldRight uint32
	DVC      byte
	Preamp   byte
	NewLeft  uint32
	NewRight uint32
}

type MessageLedc struct {
	zero       byte
	Red        byte
	Green      byte
	Blue       byte
	OnTime     uint16
	OffTime    uint16
	Times      byte
	Transition byte
}

type MessageStrm struct {
	Command         byte
	Autostart       byte
	Format          byte
	PCMSampleSize   byte
	PCMSampleRate   byte
	PCMChannels     byte
	PCMEndian       byte
	Threshold       byte
	SPDIFEnable     byte
	TransPeriod     byte
	TransType       byte
	Flags           byte
	OutputThreshold byte
	reserved        byte
	ReplayGain      [4]byte
	ServerPort      uint16
	ServerIP        [4]byte
}

type SlimServer struct {
	Clients map[string]net.TCPConn
}

func (m MessageHELO) Command() string {
	return "HELO"
}
func (m MessageSTAT) Command() string {
	return "STAT"
}


func (*SlimServer) Serve(commands chan SlimRegChan) {

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

fmt.Println(conn)
		// TBD: Create instance specific chans and pass to EventHandler on the SlimRegChan
		// Kick off a clientHandler for conn and the two chans
		//go clientHandler(conn, commands)
	}
}

// TBD: ClientHandler should do both R/W with client
func clientHandler(conn net.Conn, commands chan SlimCommand) {
	slimLog.Info("Got incomming connection from %s", conn.RemoteAddr())
	for {
		select {
		//case command := <- commands:
		//	slimLog.Info("command")
		//case message := < message
		}
	}

}

// TBD: Only per-message parsing should be made here, all state handling is done in EventHandler per-player FSM
func messageChannel(conn net.Conn, m chan Message) {
	var header MessageHeader

	err := binary.Read(conn, binary.BigEndian, &header)
	if err != nil {
		slimLog.Info("%s", err)
		return
	}
	cmd := string(header.Command[:4])
	fmt.Printf("command = %v, msgLen = %v\n", cmd, header.MsgLen)
	switch cmd {
	case "HELO":
		if header.MsgLen != 36 {
			slimLog.Info("Expecting 36 bytes HELO, got %d\n",
				header.MsgLen)
			return
		}
		var msg MessageHELO
		err = binary.Read(conn, binary.BigEndian, &msg)
		if err != nil {
			slimLog.Info("%s", err)
			return
		}
		m <- msg
	case "STAT":
		if header.MsgLen != 53 {
			slimLog.Info("Expecting 53 bytes STAT, got %d\n",
				header.MsgLen)
			return
		}
		var msg MessageSTAT
		err = binary.Read(conn, binary.BigEndian, &msg)
		if err != nil {
			slimLog.Info("%s", err)
			return
		}
		m <- msg
	}
}
