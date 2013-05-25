package slimserver

type ServerMessage interface {
	ServerCommandName() string
}

type ClientMessage interface {
	ClientCommandName() string
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

func (m MessageHELO) ClientCommandName() string {
	return "HELO"
}
func (m MessageSTAT) ClientCommandName() string {
	return "STAT"
}
func (m MessageStrm) ServerCommandName() string {
	return "strm"
}

