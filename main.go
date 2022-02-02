package gostealthclient

import (
	"os"

	"github.com/drabadan/gostealthclient/config"
	"github.com/drabadan/gostealthclient/internal/connection"
)

var senderFunc func(spd *scPacketData)
var receiverFunc func(s uint16, rtype byte) []byte
var packetLog = make([]*scPacketData, 0)
var responsesLog = make([][]byte, 0)

func Bootstrap(script func() interface{}) interface{} {
	cfg := config.NewConfig(0)
	cm := connection.NewConnectionManager(*cfg)
	conn, err := cm.Connect()
	defer conn.Close()

	if err != nil {
		os.Exit(5)
	}

	senderFunc = sender(conn)
	mws := make([]func(readBuff []byte) []byte, 0)
	receiverFunc = defaultReceiver(conn, &mws)
	return script()
}
