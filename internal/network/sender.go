package network

import (
	"log"
	"net"

	"github.com/drabadan/gostealthclient/config"
)

var instance *Sender

type Sender struct {
	conn *net.TCPConn
	pl   *[]ScPacketData
}

func NewSender(conn *net.TCPConn, pl *[]ScPacketData) {
	instance = &Sender{
		conn: conn,
		pl:   pl,
	}
}

func GetInstance() *Sender {
	if instance == nil {
		log.Fatal("Sender not initialized!")
	}

	return instance
}

// Default sender func that writes to connection
func (s *Sender) Send(spd *ScPacketData) {
	*s.pl = append(*s.pl, *spd)
	if config.DEBUG {
		log.Printf("Send packet: % x", spd.bytesToSend)
	}
	_, err := s.conn.Write(spd.bytesToSend)
	if err != nil {
		log.Fatal(err)
	}
}
