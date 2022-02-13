package gostealthclient

import (
	"os"

	"github.com/drabadan/gostealthclient/config"
	"github.com/drabadan/gostealthclient/internal/connection"
	"github.com/drabadan/gostealthclient/internal/network"
)

var packetLog = make([]network.ScPacketData, 0)

type Middleware = func(readBuff []byte) []byte

func Bootstrap(script func() interface{}) interface{} {
	cfg := config.NewConfig(0)
	cm := connection.NewConnectionManager(*cfg)
	conn, err := cm.Connect()

	if err != nil {
		os.Exit(5)
	}

	defer conn.Close()

	network.NewSender(conn, &packetLog)
	mws := make([]Middleware, 0)
	network.NewReciever(conn, &mws, config.DEBUG)
	return script()
}
