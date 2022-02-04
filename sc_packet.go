package gostealthclient

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/drabadan/gostealthclient/internal/composer"
	"github.com/drabadan/gostealthclient/internal/encoder"
	"github.com/ghostiam/binstruct"
)

type scPacketData struct {
	bytesToSend []byte
}

func (p *scPacketData) setSendBytes(pNum uint16, args ...interface{}) {
	encoder := encoder.NewEncoder(binary.LittleEndian)
	composer := composer.NewComposer(encoder)
	composer.SetDatabytes(args)
	composer.SetHeader(pNum)
	composer.SetPacketId()
	p.bytesToSend = composer.GetBytesToSend()
}

func (p *scPacketData) send(sender func(packet *scPacketData)) {
	if debug {
		log.Printf("Packet sent: % x", p.bytesToSend)
	}
	sender(p)
}

type scCompositePacketData struct {
	scPacketData
	rb chan []byte
}

type voidPacket struct {
	scPacketData
}

func NewVoidPacket(packetNum uint16, args ...interface{}) *voidPacket {
	p := &voidPacket{}
	p.setSendBytes(packetNum, args...)
	return p
}

type uint16Packet struct {
	scPacketData
	out chan uint16
}

func NewUint16Packet(packetNum uint16, args ...interface{}) *uint16Packet {
	p := &uint16Packet{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan uint16)
	go receiveUint16(p.out)
	return p
}

type uint32Packet struct {
	scPacketData
	out chan uint32
}

func NewUint32Packet(packetNum uint16, args ...interface{}) *uint32Packet {
	p := &uint32Packet{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan uint32)
	go receiveUint32(p.out)
	return p
}

type stringPacket struct {
	scPacketData
	out chan string
}

func NewStringPacket(packetNum uint16, args ...interface{}) *stringPacket {
	p := &stringPacket{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan string)
	go receiveString(p.out)
	return p
}

type bytePacket struct {
	scPacketData
	out chan byte
}

func NewBytePacket(packetNum uint16, args ...interface{}) *bytePacket {
	p := &bytePacket{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan byte)
	go receiveByte(p.out)
	return p
}

type boolPacket struct {
	scPacketData
	out chan bool
}

func NewBoolPacket(packetNum uint16, args ...interface{}) *boolPacket {
	p := &boolPacket{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan bool)
	go receiveBool(p.out)
	return p
}

type int8Packet struct {
	scPacketData
	out chan int8
}

func NewInt8Packet(packetNum uint16, args ...interface{}) *int8Packet {
	p := &int8Packet{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan int8)
	go receiveInt8(p.out)
	return p
}

type intPacket struct {
	scPacketData
	out chan int32
}

func NewIntPacket(packetNum uint16, args ...interface{}) *intPacket {
	p := &intPacket{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan int32)
	go receiveInt(p.out)
	return p
}

type floatPacket struct {
	scPacketData
	out chan float64
}

func NewFloatPacket(packetNum uint16, args ...interface{}) *floatPacket {
	p := &floatPacket{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan float64)
	go receiveFloat(p.out)
	return p
}

type timePacket struct {
	scPacketData
	out chan time.Time
}

func NewTimePacket(packetNum uint16, args ...interface{}) *timePacket {
	p := &timePacket{}
	p.setSendBytes(packetNum, args...)
	p.out = make(chan time.Time)
	go receiveTime(p.out)
	return p
}

type stealthClientInfo struct {
	// result['StealthVersion'] = _struct.unpack('<3H', data[4:10])
	// result['Build'] = _struct.unpack('<H', data[10:12])[0]
	// result['BuildDate'] = _ddt2pdt(_struct.unpack('<d', data[12:20])[0])
	// result['GITRevNumber'] = _struct.unpack('<H', data[20:22])[0]
	// result['GITRevision'] = _str.from_buffer(data[22:]).value
	stealthVersion []byte
	build          uint16
	buildDate      time.Time
	gitRevNumber   uint16
	gitRevision    string
}

//TODO: Broken - mapping not working correct
func (p *stealthClientInfoPacket) transform() {
	defer close(p.out)
	b := <-p.rb
	p.out <- stealthClientInfo{
		stealthVersion: b[4:10],
		build:          binary.LittleEndian.Uint16(b[10:12]),
		buildDate:      time.Now(),
		gitRevNumber:   binary.LittleEndian.Uint16(b[20:22]),
		gitRevision:    DecodeUtf16(b[22:]),
	}
}

type stealthClientInfoPacket struct {
	scCompositePacketData
	out chan stealthClientInfo
}

func NewStealthClientInfoPacket() *stealthClientInfoPacket {
	p := &stealthClientInfoPacket{}
	p.setSendBytes(SCGetStealthInfo)
	p.rb = make(chan []byte)
	p.out = make(chan stealthClientInfo)
	go receiveByteArray(p.rb)
	return p
}

type readStaticsXYPacket struct {
	scCompositePacketData
	out chan []StaticsXY
}

func (p *readStaticsXYPacket) transform() {
	defer close(p.out)
	b := <-p.rb

	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := make([]StaticsXY, 0)

	size := 9

	wb := b[8:]

	for i := 0; i < count; i++ {
		offset := i * size
		a := StaticsXY{
			Tile:  binary.LittleEndian.Uint16(wb[offset : offset+2]),
			X:     binary.LittleEndian.Uint16(wb[offset+2 : offset+4]),
			Y:     binary.LittleEndian.Uint16(wb[offset+4 : offset+6]),
			Z:     wb[offset+6],
			Color: binary.LittleEndian.Uint16(wb[offset+7 : offset+9]),
		}
		r = append(r, a)
	}

	p.out <- r
}

func NewReadStaticsXYPacket(args ...interface{}) *readStaticsXYPacket {
	p := &readStaticsXYPacket{}
	p.setSendBytes(SCReadStaticsXY, args...)
	p.rb = make(chan []byte)
	p.out = make(chan []StaticsXY)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type getMultisPacket struct {
	scCompositePacketData
	out chan []Multi
}

func NewGetMultisPacket(args ...interface{}) *getMultisPacket {
	p := &getMultisPacket{}
	p.setSendBytes(SCGetMultis)
	p.rb = make(chan []byte)
	p.out = make(chan []Multi)
	go receiveByteArray(p.rb)
	go p.transform()

	return p
}

func (p *getMultisPacket) transform() {
	defer close(p.out)
	b := <-p.rb

	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := make([]Multi, 0)

	size := 9

	wb := b[8:]

	for i := 0; i < count; i++ {
		offset := i * size
		a := Multi{
			Id:     binary.LittleEndian.Uint32(wb[offset : offset+4]),
			X:      binary.LittleEndian.Uint16(wb[offset+4 : offset+6]),
			Y:      binary.LittleEndian.Uint16(wb[offset+6 : offset+8]),
			Z:      wb[offset+8],
			XMin:   binary.LittleEndian.Uint16(wb[offset+9 : offset+11]),
			XMax:   binary.LittleEndian.Uint16(wb[offset+11 : offset+13]),
			YMin:   binary.LittleEndian.Uint16(wb[offset+13 : offset+15]),
			YMax:   binary.LittleEndian.Uint16(wb[offset+15 : offset+17]),
			Width:  binary.LittleEndian.Uint16(wb[offset+17 : offset+19]),
			Height: binary.LittleEndian.Uint16(wb[offset+19 : offset+21]),
		}
		r = append(r, a)
	}

	p.out <- r
}

type uint32ArrayPacket struct {
	scCompositePacketData
	out chan []uint32
}

func (p *uint32ArrayPacket) transform() {
	defer close(p.out)
	b := <-p.rb

	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := make([]uint32, 0)

	size := 4

	wb := b[8:]

	for i := 0; i < count; i++ {
		offset := i * size
		a := binary.LittleEndian.Uint32(wb[offset : offset+4])
		r = append(r, a)
	}

	p.out <- r
}

func NewUint32ArrayPacket(packetNum uint16, args ...interface{}) *uint32ArrayPacket {
	p := &uint32ArrayPacket{}
	p.setSendBytes(packetNum, args)
	p.rb = make(chan []byte)
	p.out = make(chan []uint32)
	go receiveByteArray(p.rb)
	go p.transform()

	return p
}

type isWorldCellPassablePacket struct {
	scCompositePacketData
	out chan WorldCellPassable
}

func (p *isWorldCellPassablePacket) transform() {
	defer close(p.out)
	b := <-p.rb
	r := WorldCellPassable{
		Passable: b[4] == 0x1,
		Z:        int8(b[5]),
	}

	p.out <- r
}

func NewIsWorldCellPassablePacket(args ...interface{}) *isWorldCellPassablePacket {
	p := &isWorldCellPassablePacket{}
	p.setSendBytes(SCIsWorldCellPassable, args...)
	p.rb = make(chan []byte)
	p.out = make(chan WorldCellPassable)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scPoint2DPacket struct {
	scCompositePacketData
	out chan Point2D
}

func (p *scPoint2DPacket) transform() {
	defer close(p.out)
	b := <-p.rb
	r := Point2D{
		X: binary.LittleEndian.Uint16(b[4:6]),
		Y: binary.LittleEndian.Uint16(b[6:8]),
	}
	p.out <- r
}

func NewPoint2DPacket(packetNum uint16, args ...interface{}) *scPoint2DPacket {
	p := &scPoint2DPacket{}
	p.setSendBytes(packetNum, args...)
	p.rb = make(chan []byte)
	p.out = make(chan Point2D)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetBuffBarInfoPacket struct {
	scCompositePacketData
	out chan BuffBarInfo
}

func (p *scGetBuffBarInfoPacket) transform() {
	defer close(p.out)
	b := <-p.rb
	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := BuffBarInfo{
		Count: byte(count),
	}
	r.Buffs = make([]BuffIcon, 0)

	size := 20

	wb := b[8:]

	for i := 0; i < count; i++ {
		offset := i * size
		a := BuffIcon{
			Attribute_ID: binary.LittleEndian.Uint16(wb[offset : offset+2]),
			TimeStart:    decodeDelphiTime(Float64frombytes(wb[offset+2 : offset+10])),
			Seconds:      binary.LittleEndian.Uint16(wb[offset+10 : offset+12]),
			ClilocID1:    binary.LittleEndian.Uint32(wb[offset+12 : offset+16]),
			ClilocID2:    binary.LittleEndian.Uint32(wb[offset+16 : offset+20]),
		}
		r.Buffs = append(r.Buffs, a)
	}

	p.out <- r
}

func NewBuffBarInfo() *scGetBuffBarInfoPacket {
	p := &scGetBuffBarInfoPacket{}
	p.setSendBytes(SCGetBuffBarInfo)
	p.rb = make(chan []byte)
	p.out = make(chan BuffBarInfo)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetExtInfoPacket struct {
	scCompositePacketData
	out chan ExtendedInfo
}

func (p *scGetExtInfoPacket) transform() {
	defer close(p.out)
	b := <-p.rb

	var ei ExtendedInfo
	err := binstruct.UnmarshalLE(b[4:], &ei)

	if err != nil {
		log.Fatalf("Failed to parse Ext info! Exiting...")
	}
	p.out <- ei
}

func NewGetExtInfoPacket() *scGetExtInfoPacket {
	p := &scGetExtInfoPacket{}
	p.setSendBytes(SCGetExtInfo)
	p.rb = make(chan []byte)
	p.out = make(chan ExtendedInfo)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetShopListPacket struct {
	scCompositePacketData
	out chan []string
}

func (p *scGetShopListPacket) transform() {
	defer close(p.out)
	b := <-p.rb
	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := make([]string, 0)

	wb := b[8:]

	var size int
	for i := 0; i < count; i++ {
		offset := i * size
		size = int(binary.LittleEndian.Uint32(wb[offset : offset+4]))
		r = append(r, DecodeUtf16(wb[offset+4:offset+size]))
	}

	p.out <- r
}

// Needs testing
func NewGetShopListPacket() *scGetShopListPacket {
	p := &scGetShopListPacket{}
	p.setSendBytes(SCGetShopList)
	p.rb = make(chan []byte)
	p.out = make(chan []string)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}
