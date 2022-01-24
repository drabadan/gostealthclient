package gostealthclient

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func getPort() (scriptPort uint16) {
	servAddr := "localhost:47602"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	// get port packet
	bytes := []byte{0x4, 0x0, 0xef, 0xbe, 0xad, 0xde}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(bytes)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	for i := 0; i < 2; i++ {
		reply := make([]byte, 32)
		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		scriptPort = binary.LittleEndian.Uint16(reply)
	}
	conn.Close()
	return
}

func sendSCLangPacket(conn *net.TCPConn) {
	// bytes := []byte{0x9, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x0, 0x0}
	bytes := make([]byte, 13)
	binary.LittleEndian.PutUint16(bytes[0:], 9)
	binary.LittleEndian.PutUint16(bytes[4:], 5)
	binary.LittleEndian.PutUint16(bytes[8:], 255)
	conn.Write(bytes)
}

func connectScript() (*net.TCPConn, error) {
	scriptPort := getPort()
	if debug {
		fmt.Printf("Stealth script port: %v", scriptPort)
	}

	servAddr := fmt.Sprintf("localhost:%v", scriptPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		return nil, err
	}

	sendSCLangPacket(conn)
	return conn, nil
}
