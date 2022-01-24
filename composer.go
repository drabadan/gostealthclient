package gostealthclient

import "encoding/binary"

type composer struct {
	db []byte
	bs []byte
}

func newComposer() *composer {
	return &composer{
		db: make([]byte, 0),
		bs: make([]byte, 8),
	}
}

func (c *composer) setDatabytes(d []interface{}) {
	if d != nil {
		transformData(&c.db, d)
	}
}

func (c *composer) setHeader(pNum uint16) {
	binary.LittleEndian.PutUint32(c.bs[0:], uint32(len(c.db)+4))
	binary.LittleEndian.PutUint16(c.bs[4:], pNum)
}

func (c *composer) setPacketId() {
	if increment >= 65535 {
		increment = 0
	}

	increment++

	binary.LittleEndian.PutUint16(c.bs[6:], increment)
}

func (c *composer) getBytesToSend() []byte {
	return append(c.bs, c.db...)
}
