package gostealthclient

import (
	"log"
	"net"
)

func sender(conn *net.TCPConn) func(spd *scPacketData) {
	return func(spd *scPacketData) {
		packetLog = append(packetLog, spd)
		_, err := conn.Write(spd.bytesToSend)
		if err != nil {
			log.Fatal(err)
		}
	}
}
