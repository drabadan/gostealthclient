package network

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/drabadan/gostealthclient/pkg/constants"
	"github.com/drabadan/gostealthclient/pkg/model"
	"github.com/ghostiam/binstruct"
)

type ScPacketData struct {
	bytesToSend []byte
}

func (p *ScPacketData) setSendBytes(pNum uint16, args ...interface{}) {
	encoder := NewEncoder(binary.LittleEndian)
	composer := NewComposer(encoder)
	composer.SetDatabytes(args)
	composer.SetHeader(pNum)
	composer.SetPacketId()
	p.bytesToSend = composer.GetBytesToSend()
	sender := GetInstance()
	sender.Send(p)
}

func (p *ScPacketData) Send() {

}

type scCompositePacketData struct {
	ScPacketData
	rb chan []byte
}

type voidPacket struct {
	ScPacketData
}

func NewVoidPacket(packetNum uint16, args ...interface{}) *voidPacket {
	p := &voidPacket{}
	p.setSendBytes(packetNum, args...)
	return p
}

type uint16Packet struct {
	ScPacketData
	Out chan uint16
}

func NewUint16Packet(packetNum uint16, args ...interface{}) *uint16Packet {
	p := &uint16Packet{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan uint16)
	go receiveUint16(p.Out)
	return p
}

type uint32Packet struct {
	ScPacketData
	Out chan uint32
}

func NewUint32Packet(packetNum uint16, args ...interface{}) *uint32Packet {
	p := &uint32Packet{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan uint32)
	go receiveUint32(p.Out)
	return p
}

type stringPacket struct {
	ScPacketData
	Out chan string
}

func NewStringPacket(packetNum uint16, args ...interface{}) *stringPacket {
	p := &stringPacket{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan string)
	go receiveString(p.Out)
	return p
}

type bytePacket struct {
	ScPacketData
	Out chan byte
}

func NewBytePacket(packetNum uint16, args ...interface{}) *bytePacket {
	p := &bytePacket{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan byte)
	go receiveByte(p.Out)
	return p
}

type byteArrayPacket struct {
	ScPacketData
	Out chan []byte
}

func NewByteArrayPacket(packetNum uint16, args ...interface{}) *byteArrayPacket {
	p := &byteArrayPacket{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan []byte)
	go receiveByteArray(p.Out)
	return p
}

type boolPacket struct {
	ScPacketData
	Out chan bool
}

func NewBoolPacket(packetNum uint16, args ...interface{}) *boolPacket {
	p := &boolPacket{}
	p.setSendBytes(packetNum, args...)

	p.Out = make(chan bool)
	go receiveBool(p.Out)
	return p
}

type int8Packet struct {
	ScPacketData
	Out chan int8
}

func NewInt8Packet(packetNum uint16, args ...interface{}) *int8Packet {
	p := &int8Packet{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan int8)
	go receiveInt8(p.Out)
	return p
}

type intPacket struct {
	ScPacketData
	Out chan int32
}

func NewIntPacket(packetNum uint16, args ...interface{}) *intPacket {
	p := &intPacket{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan int32)
	go receiveInt(p.Out)
	return p
}

type floatPacket struct {
	ScPacketData
	Out chan float64
}

func NewFloatPacket(packetNum uint16, args ...interface{}) *floatPacket {
	p := &floatPacket{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan float64)
	go receiveFloat(p.Out)
	return p
}

type timePacket struct {
	ScPacketData
	Out chan time.Time
}

func NewTimePacket(packetNum uint16, args ...interface{}) *timePacket {
	p := &timePacket{}
	p.setSendBytes(packetNum, args...)
	p.Out = make(chan time.Time)
	go receiveTime(p.Out)
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
	defer close(p.Out)
	b := <-p.rb
	p.Out <- stealthClientInfo{
		stealthVersion: b[4:10],
		build:          binary.LittleEndian.Uint16(b[10:12]),
		buildDate:      time.Now(),
		gitRevNumber:   binary.LittleEndian.Uint16(b[20:22]),
		gitRevision:    DecodeUtf16(b[22:]),
	}
}

type stealthClientInfoPacket struct {
	scCompositePacketData
	Out chan stealthClientInfo
}

func NewStealthClientInfoPacket() *stealthClientInfoPacket {
	p := &stealthClientInfoPacket{}
	p.setSendBytes(constants.SCGetStealthInfo)
	p.rb = make(chan []byte)
	p.Out = make(chan stealthClientInfo)
	go receiveByteArray(p.rb)
	return p
}

type readStaticsXYPacket struct {
	scCompositePacketData
	Out chan []model.StaticsXY
}

func (p *readStaticsXYPacket) transform() {
	defer close(p.Out)
	b := <-p.rb

	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := make([]model.StaticsXY, 0)

	size := 9

	wb := b[8:]

	for i := 0; i < count; i++ {
		offset := i * size
		a := model.StaticsXY{
			Tile:  binary.LittleEndian.Uint16(wb[offset : offset+2]),
			X:     binary.LittleEndian.Uint16(wb[offset+2 : offset+4]),
			Y:     binary.LittleEndian.Uint16(wb[offset+4 : offset+6]),
			Z:     wb[offset+6],
			Color: binary.LittleEndian.Uint16(wb[offset+7 : offset+9]),
		}
		r = append(r, a)
	}

	p.Out <- r
}

func NewReadStaticsXYPacket(args ...interface{}) *readStaticsXYPacket {
	p := &readStaticsXYPacket{}
	p.setSendBytes(constants.SCReadStaticsXY, args...)
	p.rb = make(chan []byte)
	p.Out = make(chan []model.StaticsXY)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type getMultisPacket struct {
	scCompositePacketData
	Out chan []model.Multi
}

func NewGetMultisPacket(args ...interface{}) *getMultisPacket {
	p := &getMultisPacket{}
	p.setSendBytes(constants.SCGetMultis)
	p.rb = make(chan []byte)
	p.Out = make(chan []model.Multi)
	go receiveByteArray(p.rb)
	go p.transform()

	return p
}

func (p *getMultisPacket) transform() {
	defer close(p.Out)
	b := <-p.rb

	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := make([]model.Multi, 0)

	size := 9

	wb := b[8:]

	for i := 0; i < count; i++ {
		offset := i * size
		a := model.Multi{
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

	p.Out <- r
}

type uint32ArrayPacket struct {
	scCompositePacketData
	Out chan []uint32
}

func (p *uint32ArrayPacket) transform() {
	defer close(p.Out)
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

	p.Out <- r
}

func NewUint32ArrayPacket(packetNum uint16, args ...interface{}) *uint32ArrayPacket {
	p := &uint32ArrayPacket{}
	p.setSendBytes(packetNum, args)
	p.rb = make(chan []byte)
	p.Out = make(chan []uint32)
	go receiveByteArray(p.rb)
	go p.transform()

	return p
}

type isWorldCellPassablePacket struct {
	scCompositePacketData
	Out chan model.WorldCellPassable
}

func (p *isWorldCellPassablePacket) transform() {
	defer close(p.Out)
	b := <-p.rb
	r := model.WorldCellPassable{
		Passable: b[4] == 0x1,
		Z:        int8(b[5]),
	}

	p.Out <- r
}

func NewIsWorldCellPassablePacket(args ...interface{}) *isWorldCellPassablePacket {
	p := &isWorldCellPassablePacket{}
	p.setSendBytes(constants.SCIsWorldCellPassable, args...)
	p.rb = make(chan []byte)
	p.Out = make(chan model.WorldCellPassable)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scPoint2DPacket struct {
	scCompositePacketData
	Out chan model.Point2D
}

func (p *scPoint2DPacket) transform() {
	defer close(p.Out)
	b := <-p.rb
	r := model.Point2D{
		X: binary.LittleEndian.Uint16(b[4:6]),
		Y: binary.LittleEndian.Uint16(b[6:8]),
	}
	p.Out <- r
}

func NewPoint2DPacket(packetNum uint16, args ...interface{}) *scPoint2DPacket {
	p := &scPoint2DPacket{}
	p.setSendBytes(packetNum, args...)
	p.rb = make(chan []byte)
	p.Out = make(chan model.Point2D)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetBuffBarInfoPacket struct {
	scCompositePacketData
	Out chan model.BuffBarInfo
}

func (p *scGetBuffBarInfoPacket) transform() {
	defer close(p.Out)
	b := <-p.rb
	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := model.BuffBarInfo{
		Count: byte(count),
	}
	r.Buffs = make([]model.BuffIcon, 0)

	size := 20

	wb := b[8:]

	for i := 0; i < count; i++ {
		offset := i * size
		a := model.BuffIcon{
			Attribute_ID: binary.LittleEndian.Uint16(wb[offset : offset+2]),
			TimeStart:    DecodeDelphiTime(Float64frombytes(wb[offset+2 : offset+10])),
			Seconds:      binary.LittleEndian.Uint16(wb[offset+10 : offset+12]),
			ClilocID1:    binary.LittleEndian.Uint32(wb[offset+12 : offset+16]),
			ClilocID2:    binary.LittleEndian.Uint32(wb[offset+16 : offset+20]),
		}
		r.Buffs = append(r.Buffs, a)
	}

	p.Out <- r
}

func NewBuffBarInfo() *scGetBuffBarInfoPacket {
	p := &scGetBuffBarInfoPacket{}
	p.setSendBytes(constants.SCGetBuffBarInfo)
	p.rb = make(chan []byte)
	p.Out = make(chan model.BuffBarInfo)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetExtInfoPacket struct {
	scCompositePacketData
	Out chan model.ExtendedInfo
}

func (p *scGetExtInfoPacket) transform() {
	defer close(p.Out)
	b := <-p.rb

	var ei model.ExtendedInfo
	err := binstruct.UnmarshalLE(b[4:], &ei)

	if err != nil {
		log.Fatalf("Failed to parse Ext info! Exiting...")
	}
	p.Out <- ei
}

func NewGetExtInfoPacket() *scGetExtInfoPacket {
	p := &scGetExtInfoPacket{}
	p.setSendBytes(constants.SCGetExtInfo)
	p.rb = make(chan []byte)
	p.Out = make(chan model.ExtendedInfo)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetShopListPacket struct {
	scCompositePacketData
	Out chan []string
}

func (p *scGetShopListPacket) transform() {
	defer close(p.Out)
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

	p.Out <- r
}

// Needs testing
func NewGetShopListPacket() *scGetShopListPacket {
	p := &scGetShopListPacket{}
	p.setSendBytes(constants.SCGetShopList)
	p.rb = make(chan []byte)
	p.Out = make(chan []string)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetMapCellPacket struct {
	scCompositePacketData
	Out chan model.MapCell
}

func (p *scGetMapCellPacket) transform() {
	defer close(p.Out)
	b := <-p.rb
	var c model.MapCell
	err := binstruct.UnmarshalLE(b[4:], &c)
	if err != nil {
		log.Fatalf("Failed to parse MapCell info! Exiting...")
	}
	p.Out <- c
}

// Needs testing
func NewGetMapCellPacket(args ...interface{}) *scGetMapCellPacket {
	p := &scGetMapCellPacket{}
	p.setSendBytes(constants.SCGetCell, args...)
	p.rb = make(chan []byte)
	p.Out = make(chan model.MapCell)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scGetStaticTilesArrayPacket struct {
	scCompositePacketData
	Out chan []model.FoundTile
}

func (p *scGetStaticTilesArrayPacket) transform() {
	defer close(p.Out)
	b := <-p.rb
	count := int(binary.LittleEndian.Uint16(b[4:8]))
	r := make([]model.FoundTile, 0)

	size := 14

	wb := b[8:]

	for i := 0; i < count; i++ {
		var f model.FoundTile
		offset := i * size
		err := binstruct.UnmarshalLE(wb[offset:offset+size], &f)
		if err != nil {
			log.Fatalf("Failed to parse FoundTile info! Exiting...")
		}

		r = append(r, f)
	}

	p.Out <- r
}

// Needs testing
func NewGetStaticTilesArrayPacket(args ...interface{}) *scGetStaticTilesArrayPacket {
	p := &scGetStaticTilesArrayPacket{}
	p.setSendBytes(constants.SCGetStaticTilesArray, args...)
	p.rb = make(chan []byte)
	p.Out = make(chan []model.FoundTile)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}

type scClientTargetInfoPacket struct {
	scCompositePacketData
	Out chan model.TargetInfo
}

func (p *scClientTargetInfoPacket) transform() {
	defer close(p.Out)
	b := <-p.rb

	var r model.TargetInfo

	err := binstruct.UnmarshalLE(b, &r)
	if err != nil {
		log.Fatalf("Failed to parse TargetInfo info! Exiting...")
	}

	p.Out <- r
}

// Needs testing
func NewClientTargetInfoPacket() *scClientTargetInfoPacket {
	p := &scClientTargetInfoPacket{}
	p.setSendBytes(constants.SCClientTargetResponse)
	p.rb = make(chan []byte)
	p.Out = make(chan model.TargetInfo)
	go receiveByteArray(p.rb)
	go p.transform()
	return p
}
