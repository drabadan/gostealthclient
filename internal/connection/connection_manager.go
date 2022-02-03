package connection

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/drabadan/gostealthclient/config"
)

// Connection Manager struct
type ConnectionManager struct {
	logLevel byte
}

// Constructor for connection manager
func NewConnectionManager(cfg config.Config) *ConnectionManager {
	return &ConnectionManager{logLevel: cfg.LogLevel}
}

func (cm *ConnectionManager) getPort() (scriptPort uint16) {
	if cm.logLevel <= config.LOG_LEVEL_INFO {
		log.Println("Fetching port from stealth")
	}

	servAddr := "localhost:47602"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to connect to stealth.\n Error: %v", err)
		os.Exit(5)
	}

	// get port packet
	bytes := []byte{0x4, 0x0, 0xef, 0xbe, 0xad, 0xde}

	for try := 0; try <= config.SOCKET_MAX_RETRIES; try++ {
		log.Printf("Try #%v", try)

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			log.Println("Dial failed:", err.Error())
			continue
		}

		_, err = conn.Write(bytes)
		if err != nil {
			log.Println("Write to server failed:", err.Error())
			conn.Close()
			continue
		}

		err = conn.SetReadDeadline(time.Now().Add(config.SOCKET_TIMEOUT * time.Second))
		if err != nil {
			log.Println("Connection timed out:", err)
			conn.Close()
			continue
		}

		for i := 0; i < 2; i++ {
			reply := make([]byte, 32)
			_, err = conn.Read(reply)
			if err != nil {
				log.Println("Read from server failed:", err.Error())
				conn.Close()
				continue
			}

			scriptPort = binary.LittleEndian.Uint16(reply)
		}

		conn.Close()

		log.Printf("Converted response: %v", scriptPort)

		if scriptPort > 2 {
			break
		}
	}

	if cm.logLevel <= config.LOG_LEVEL_INFO {
		log.Printf("ScriptPort retrieved: %v\n", scriptPort)
	}
	return
}

func (cm *ConnectionManager) sendSCLangPacket(conn *net.TCPConn) {
	bytes := make([]byte, 13)
	binary.LittleEndian.PutUint16(bytes[0:], 9)
	binary.LittleEndian.PutUint16(bytes[4:], 5)
	binary.LittleEndian.PutUint16(bytes[8:], 255)
	conn.Write(bytes)
}

// Connect to running stealth client application
func (cm *ConnectionManager) Connect() (*net.TCPConn, error) {
	scriptPort := cm.getPort()

	servAddr := fmt.Sprintf(":%v", scriptPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatalf("ResolveTCPAddr failed!\nError: %v", err.Error())
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("Dial failed: %v", err.Error())
		return nil, err
	}

	cm.sendSCLangPacket(conn)
	return conn, nil
}
