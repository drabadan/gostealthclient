package gostealthclient

import (
	"os"
)

var senderFunc func(spd *scPacketData)
var receiverFunc func(s uint16, rtype byte) []byte
var packetLog = make([]*scPacketData, 0)
var responsesLog = make([][]byte, 0)

func Bootstrap(script func()) {
	conn, err := connectScript()
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	senderFunc = sender(conn)
	mws := make([]func(readBuff []byte) []byte, 0)
	receiverFunc = defaultReceiver(conn, &mws)
	script()
}
