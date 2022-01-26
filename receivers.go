package gostealthclient

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const (
	SIZE_READ_TYPE = 0
	BODY_READ_TYPE = 1
)

func defaultReceiver(conn *net.TCPConn, mws *[]func(readBuff []byte) []byte) func(s uint16, rtype byte) []byte {
	paused := false
	var lpb []byte
	return func(s uint16, rtype byte) []byte {
		readBuff := make([]byte, s)
		_, err := conn.Read(readBuff)

		if err != nil {
			log.Fatal("Received err from tcp connection. Shutting down...")
			os.Exit(500)
		}

		if rtype == BODY_READ_TYPE {
			responsesLog = append(responsesLog, readBuff)
			t := binary.LittleEndian.Uint16(readBuff[0:2])
			switch t {
			case 1:
				return readBuff
			case 2:
				log.Fatal("Received terminate script command. Exiting...")
				os.Exit(0)
			case 4:
				if !paused {
					log.Println("[WARNING] Received pause script packet. Waiting to proceed...")
					lpb = receiverFunc(readReplySize(), BODY_READ_TYPE)
					log.Printf("[INFO] Last packet: % x", lpb)

					paused = true

					// wait for unlock
					time.Sleep(time.Millisecond * 100)
					return receiverFunc(readReplySize(), BODY_READ_TYPE)
				} else {
					paused = false
					return lpb
				}
			case 3:
				log.Fatal("Events not implemented! Exiting...")
				os.Exit(1)
			}
		}

		return readBuff
	}
}

const READ_SIZE_RETRY = 3

func readReplySizeAndType(r uint) uint16 {
	readBuff := receiverFunc(4, SIZE_READ_TYPE)
	s := binary.LittleEndian.Uint16(readBuff[0:4])
	if s > 0 {
		return s
	} else if r < READ_SIZE_RETRY {
		r++
		return readReplySizeAndType(r)
	} else {
		log.Fatal("Read packet size failed. Exiting...")
		os.Exit(500)
	}

	return 0
}

func readReplySize() uint16 {
	s := readReplySizeAndType(0)
	return s
}

func readUint16Response() uint16 {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return binary.LittleEndian.Uint16(replyBuf[4:])
}

func receiveUint16(r chan (uint16)) {
	defer close(r)

	time.Sleep(time.Millisecond * 50)
	r <- readUint16Response()
}

func readUint32Response() uint32 {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return binary.LittleEndian.Uint32(replyBuf[4:])
}

func receiveUint32(r chan (uint32)) {
	defer close(r)

	time.Sleep(time.Millisecond * 50)
	r <- readUint32Response()
}

func readBool() bool {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return replyBuf[4] == 0x1
}

func receiveBool(r chan bool) {
	defer close(r)

	time.Sleep(READ_DELAY)
	r <- readBool()
}

func readInt() int32 {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	i := binary.LittleEndian.Uint32(replyBuf[4:])
	return int32(i)
}

func receiveInt(r chan int32) {
	defer close(r)

	time.Sleep(READ_DELAY)
	r <- readInt()
}

func decodeUtf16(inputBytes []byte) string {
	win16be := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(inputBytes[6:]), utf16bom)

	// decode and print:
	decoded, err := ioutil.ReadAll(unicodeReader)
	if err != nil {
		log.Fatal("Failed to parse string from packet...")
		os.Exit(500)
	}

	return strings.Replace(string(decoded), "\r\n", "\n", -1)
}

func readStringResponse() string {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return decodeUtf16(replyBuf)
}

func receiveString(r chan string) {
	defer close(r)
	// time.Sleep(READ_DELAY)
	r <- readStringResponse()
}

func readByte() byte {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return replyBuf[4]
}

func receiveByte(r chan byte) {
	defer close(r)
	time.Sleep(READ_DELAY)
	r <- readByte()
}

func readByteArrayResponse() []byte {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return replyBuf
}

func receiveByteArray(r chan []byte) {
	defer close(r)
	time.Sleep(READ_DELAY)
	r <- readByteArrayResponse()
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func readFloatResponse() float64 {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return Float64frombytes(replyBuf[4:])
}

func receiveFloat(r chan float64) {
	defer close(r)

	time.Sleep(READ_DELAY)
	r <- readFloatResponse()
}

func readTimeResponse() time.Time {
	f := readFloatResponse()
	t := time.Date(1899, 12, 30, 00, 00, 00, 00, time.Local)

	d, _ := math.Modf(f)

	ds := fmt.Sprintf("%vh", (int32(d) * 24))
	dd, _ := time.ParseDuration(ds)
	hs := fmt.Sprintf("%vh", (24 * (f - d)))
	hh, _ := time.ParseDuration(hs)

	t = t.Add(dd).Add(hh)

	return t
}

func receiveTime(r chan time.Time) {
	defer close(r)
	time.Sleep(READ_DELAY)
	r <- readTimeResponse()
}
