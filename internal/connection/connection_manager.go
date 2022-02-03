package connection

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/drabadan/gostealthclient/config"
)

// ConnectionManager basic struct
type ConnectionManager struct {
	logLevel byte
}

// NewConnectionManager is a constructor for connection manager
func NewConnectionManager(cfg config.Config) *ConnectionManager {
	return &ConnectionManager{logLevel: cfg.LogLevel}
}

// getPort returns Stealth port used for running external scripts
// to get rid of i\o timeouts retries are used
func (cm *ConnectionManager) getPort() (scriptPort uint16) {
	if cm.logLevel <= config.LOG_LEVEL_INFO {
		log.Println("Fetching port from stealth")
	}

	// get port packet
	bytes := []byte{0x4, 0x0, 0xef, 0xbe, 0xad, 0xde}

	for try := 0; try <= config.SOCKET_MAX_RETRIES; try++ {
		log.Printf("Try #%v", try)

		conn, err := cm.getConnection(config.SOCKET_HOST, config.SOCKET_PORT)
		if err != nil {
			log.Println("Failed to get connection, original error:", err.Error())
			continue
		}

		err = conn.SetReadDeadline(time.Now().Add(config.SOCKET_TIMEOUT * time.Second))
		if err != nil {
			log.Println("Connection timed out:", err)
			conn.Close()
			continue
		}

		_, err = conn.Write(bytes)
		if err != nil {
			log.Println("Write to server failed:", err.Error())
			conn.Close()
			continue
		}

		// Will get 2 replies - length followed by actual port
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

		// If got port, not a length as a response
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

// getConnection returns *net.TCPConn if address resolving and dial succeeded
func (cm *ConnectionManager) getConnection(host string, port uint16) (*net.TCPConn, error) {
	if host == "" {
		host = config.SOCKET_HOST
	}
	if port == 0 {
		port = config.SOCKET_PORT
	}

	servAddr := fmt.Sprintf("%v:%v", host, port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		return nil, fmt.Errorf("ResolveTCPAddr failed!\nError: %w", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("Dial failed: %w", err)
	}

	return conn, err
}

// Connect connects to running stealth client application
func (cm *ConnectionManager) Connect() (*net.TCPConn, error) {
	scriptPort := cm.getPort()

	conn, err := cm.getConnection(config.SOCKET_HOST, scriptPort)

	if err != nil {
		log.Fatalf("Unable to get connection. Original error: %v", err.Error())
	}

	cm.sendSCLangPacket(conn)
	return conn, nil
}
