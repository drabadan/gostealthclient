package network

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net"
	"time"
)

const (
	SIZE_READ_TYPE = 0
	BODY_READ_TYPE = 1
	READ_DELAY     = 0 // time.Microsecond * 50
)

var receiver *Receiver
var receiverFunc func(s uint16, rtype byte) []byte

type Receiver struct {
	conn  *net.TCPConn
	mws   *[]func(readBuff []byte) []byte
	debug bool
}

func NewReciever(conn *net.TCPConn, mws *[]func(readBuff []byte) []byte, debug bool) {
	receiver = &Receiver{
		conn,
		mws,
		debug,
	}

	receiverFunc = receiver.receive()
}

func (r *Receiver) receive() func(s uint16, rtype byte) []byte {
	paused := false
	var lpb []byte

	return func(s uint16, rtype byte) []byte {
		readBuff := make([]byte, s)
		_, err := r.conn.Read(readBuff)

		if err != nil {
			log.Fatal("Received err from tcp connection. Shutting down...")
		}

		if rtype == BODY_READ_TYPE {
			if r.debug {
				log.Printf("Received body: % x", readBuff)
			}

			// responsesLog = append(responsesLog, readBuff)
			t := binary.LittleEndian.Uint16(readBuff[0:2])
			switch t {
			case 1:
				return readBuff
			case 2:
				log.Fatal("Received terminate script command. Exiting...")
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

	time.Sleep(READ_DELAY)
	r <- readUint16Response()
}

func readUint32Response() uint32 {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return binary.LittleEndian.Uint32(replyBuf[4:])
}

func receiveUint32(r chan (uint32)) {
	defer close(r)

	time.Sleep(READ_DELAY)
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

func readInt8() int8 {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return int8(replyBuf[4])
}

func receiveInt8(r chan int8) {
	defer close(r)
	time.Sleep(READ_DELAY)
	r <- readInt8()
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

func readStringResponse() string {
	replyBuf := receiverFunc(readReplySize(), BODY_READ_TYPE)
	return DecodeUtf16(replyBuf)
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
