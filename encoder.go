package gostealthclient

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

func encodeString(data string, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	sizeBuf := new(bytes.Buffer)
	encoded := utf16.Encode([]rune(data))
	binary.Write(buf, binary.LittleEndian, encoded)
	binary.Write(sizeBuf, binary.LittleEndian, uint32(buf.Len()))
	*dataBytes = append(*dataBytes, sizeBuf.Bytes()...)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func encodeDWord(data uint32, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]uint32, 0)
	r = append(r, uint32(data))
	binary.Write(buf, binary.LittleEndian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func encodeWord(data uint16, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]uint16, 0)
	r = append(r, uint16(data))
	binary.Write(buf, binary.LittleEndian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func encodeByte(data byte, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]byte, 0)
	r = append(r, byte(data))
	binary.Write(buf, binary.LittleEndian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func encodeInt(data int32, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]int32, 0)
	r = append(r, int32(data))
	binary.Write(buf, binary.LittleEndian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func encodeBool(data bool, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	r := make([]byte, 0)
	if data {
		r = append(r, 1)
	} else {
		r = append(r, 0)
	}
	binary.Write(buf, binary.LittleEndian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func EncodeTime(data time.Time, dataBytes *[]byte) {
	buf := new(bytes.Buffer)
	loc, _ := time.LoadLocation("Europe/Berlin")
	t := time.Date(1899, 12, 30, 00, 00, 00, 00, loc)
	delta := data.Sub(t)
	r := float64(delta.Microseconds()) / 1000000 / 60 / 60 / 24
	binary.Write(buf, binary.LittleEndian, r)
	*dataBytes = append(*dataBytes, buf.Bytes()...)
}

func encodeIterable(dataBytes *[]byte, data reflect.Value) {
	size := make([]byte, 4)
	binary.LittleEndian.PutUint32(size, uint32(data.Len()))
	*dataBytes = append(*dataBytes, size...)

	for i := 0; i < data.Len(); i++ {
		val := data.Index(i)
		transformType(dataBytes, val.Interface())
	}
}

func transformType(dataBytes *[]byte, v interface{}) {
	if v != nil {
		if str, ok := v.(string); ok {
			encodeString(str, dataBytes)
		} else if dword, ok := v.(uint32); ok {
			encodeDWord(dword, dataBytes)
		} else if byte, ok := v.(byte); ok {
			encodeByte(byte, dataBytes)
		} else if word, ok := v.(uint16); ok {
			encodeWord(word, dataBytes)
		} else if int, ok := v.(int32); ok {
			encodeInt(int, dataBytes)
		} else if boolean, ok := v.(bool); ok {
			encodeBool(boolean, dataBytes)
		} else if t, ok := v.(time.Time); ok {
			EncodeTime(t, dataBytes)
		} else {
			log.Fatalf("Failed to parse argument of type %v", fmt.Sprintf("%T", v))
			os.Exit(500)
		}
	}
}

func transformData(dataBytes *[]byte, data []interface{}) {
	for _, v := range data {
		rt := reflect.ValueOf(v).Kind()
		switch rt {
		case reflect.Slice, reflect.Array:
			varr := reflect.ValueOf(v)
			encodeIterable(dataBytes, varr)
		default:
			transformType(dataBytes, v)
		}
	}
}
