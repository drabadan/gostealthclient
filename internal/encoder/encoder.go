package encoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
	"unicode/utf16"
)

type Encoder struct {
	endian binary.ByteOrder
}

func NewEncoder(endian binary.ByteOrder) *Encoder {
	return &Encoder{endian: endian}
}

func (e *Encoder) encodeString(data string, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	sizeBuf := new(bytes.Buffer)
	encoded := utf16.Encode([]rune(data))
	binary.Write(buf, e.endian, encoded)
	binary.Write(sizeBuf, e.endian, uint32(buf.Len()))
	*dataBytes = append(*dataBytes, sizeBuf.Bytes()...)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeDWord(data uint32, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]uint32, 0)
	r = append(r, uint32(data))
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeWord(data uint16, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]uint16, 0)
	r = append(r, uint16(data))
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeByte(data byte, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]byte, 0)
	r = append(r, byte(data))
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeInt(data int32, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]int32, 0)
	r = append(r, int32(data))
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeBool(data bool, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]byte, 0)
	if data {
		r = append(r, 1)
	} else {
		r = append(r, 0)
	}
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeTime(data time.Time, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	t := time.Date(1899, 12, 30, 00, 00, 00, 00, time.Local)
	delta := data.Sub(t)
	r := float64(delta.Microseconds()) / 1000000 / 60 / 60 / 24
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeIterable(dataBytes *[]byte, data reflect.Value) {
	size := make([]byte, 4)
	binary.LittleEndian.PutUint32(size, uint32(data.Len()))
	*dataBytes = append(*dataBytes, size...)

	for i := 0; i < data.Len(); i++ {
		val := data.Index(i)
		e.TransformType(dataBytes, val.Interface())
	}
}

func (e *Encoder) encodeInt8(data int8, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]int8, 0)
	r = append(r, int8(data))
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) encodeInt16(data int16, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]int16, 0)
	r = append(r, int16(data))
	binary.Write(buf, e.endian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func (e *Encoder) TransformType(dataBytes *[]byte, v interface{}) {
	if v != nil {
		if str, ok := v.(string); ok {
			e.encodeString(str, dataBytes)
		} else if dword, ok := v.(uint32); ok {
			e.encodeDWord(dword, dataBytes)
		} else if byte, ok := v.(byte); ok {
			e.encodeByte(byte, dataBytes)
		} else if word, ok := v.(uint16); ok {
			e.encodeWord(word, dataBytes)
		} else if int, ok := v.(int32); ok {
			e.encodeInt(int, dataBytes)
		} else if boolean, ok := v.(bool); ok {
			e.encodeBool(boolean, dataBytes)
		} else if t, ok := v.(time.Time); ok {
			e.encodeTime(t, dataBytes)
		} else if i8, ok := v.(int8); ok {
			e.encodeInt8(i8, dataBytes)
		} else if i16, ok := v.(int16); ok {
			e.encodeInt16(i16, dataBytes)
		} else {
			log.Fatalf("Failed to parse argument of type %v", fmt.Sprintf("%T", v))
			os.Exit(500)
		}
	}
}

func (e *Encoder) TransformData(dataBytes *[]byte, data []interface{}) {
	for _, v := range data {
		rt := reflect.ValueOf(v).Kind()
		switch rt {
		case reflect.Slice, reflect.Array:
			varr := reflect.ValueOf(v)
			e.encodeIterable(dataBytes, varr)
		default:
			e.TransformType(dataBytes, v)
		}
	}
}
