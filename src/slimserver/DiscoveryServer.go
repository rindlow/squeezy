package slimserver

import (
	"bytes"
	"encoding/binary"
	"net"
	//"fmt"
	//"time"
        "github.com/op/go-logging"
)

var discoLog = logging.MustGetLogger("disco")

type DiscoveryResponseD struct {
	code byte
	res1 byte
	ip   [4]byte
	port [2]byte
	res2 [10]byte
}

func DiscoveryServer() {
	var mac net.HardwareAddr
	addr, err := net.ResolveUDPAddr("udp", ":3483")
	if err != nil {
		discoLog.Panic("%s", err)
	}
	server, err := net.ListenUDP("udp", addr)
	if err != nil {
		discoLog.Panic("%s", err)
	}

	go func() {
		for {
			payload := make([]byte, 256)
			n, raddr, err := server.ReadFrom(payload)
			if err != nil {
				discoLog.Info("%s", err)
				continue
			}

			discoLog.Info("Received %d bytes from %v\n", n, raddr)
			if n < 1 {
				continue
			}
			switch payload[0] {
			case 'd':
				if n < 18 {
					discoLog.Info("%d bytes is too short\n",
						n)
					continue
				}
				mac = payload[12:18]
				discoLog.Info("Discovery (d) from %v %v at %v\n",
					deviceName(payload[2]), mac, raddr)
				uraddr, err := net.ResolveUDPAddr("udp4",
					raddr.String())
				if err != nil {
					discoLog.Panic("%s", err)
				}

				response := new(DiscoveryResponseD)
				response.code = 'D'
				response.port[0] = 155
				response.port[1] = 13
				buf := new(bytes.Buffer)
				err = binary.Write(buf,
					binary.LittleEndian, response)
				if err != nil {
					discoLog.Panic("%s", err)
					continue
				}
				l, err := server.WriteToUDP(buf.Bytes(), uraddr)
				if err != nil {
					discoLog.Panic("%s", err)
					continue
				}
				discoLog.Info("sent %d bytes to %v\n", l, uraddr)
			}
		}
	}()
}
