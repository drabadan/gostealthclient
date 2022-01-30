package composer

import (
	"encoding/binary"

	"github.com/drabadan/gostealthclient/internal/encoder"
)

var increment uint16 = 0

type Composer struct {
	encoder *encoder.Encoder
	db      []byte
	bs      []byte
}

func NewComposer(e *encoder.Encoder) *Composer {
	return &Composer{
		encoder: e,
		db:      make([]byte, 0),
		bs:      make([]byte, 8),
	}
}

func (c *Composer) SetDatabytes(d []interface{}) {
	if d != nil {
		c.encoder.TransformData(&c.db, d)
	}
}

func (c *Composer) SetHeader(pNum uint16) {
	binary.LittleEndian.PutUint32(c.bs[0:], uint32(len(c.db)+4))
	binary.LittleEndian.PutUint16(c.bs[4:], pNum)
}

func (c *Composer) SetPacketId() {
	if increment >= 65535 {
		increment = 0
	}

	increment++

	binary.LittleEndian.PutUint16(c.bs[6:], increment)
}

func (c *Composer) GetBytesToSend() []byte {
	return append(c.bs, c.db...)
}
