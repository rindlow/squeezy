package slimserver

func deviceName(id byte) string {
	switch id {
	case 1: return "SLIMP3"
	case 2: return "SqueezeBox"
	case 3: return "SoftSqueeze"
	case 4: return "SqueezeBox2"
	case 5: return "Transporter"
	case 6: return "SoftSqueeze3"
	case 7: return "Receiver" 
	case 8: return "SqueezeSlave"
	case 9: return "Controller"
	case 10: return "Boom"
	case 11: return "SoftBoom"
	case 12: return "SqueezePlay" 
	}
	return "Unknown"
}
