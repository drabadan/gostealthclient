package gostealthclient

import (
	"fmt"
	"log"
	"math"
	"net"
	"time"
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

func decodeDelphiTime(double float64) time.Time {
	f := double
	t := time.Date(1899, 12, 30, 00, 00, 00, 00, time.Local)

	d, _ := math.Modf(f)

	ds := fmt.Sprintf("%vh", (int32(d) * 24))
	dd, _ := time.ParseDuration(ds)
	hs := fmt.Sprintf("%vh", (24 * (f - d)))
	hh, _ := time.ParseDuration(hs)

	t = t.Add(dd).Add(hh)

	return t
}
