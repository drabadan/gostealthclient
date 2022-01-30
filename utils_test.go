package gostealthclient

import (
	"log"
	"reflect"
	"testing"

	"github.com/ghostiam/binstruct"
)

func TestStructFields(t *testing.T) {
	s := reflect.TypeOf(ExtendedInfo{})
	fs := reflect.VisibleFields(s)
	log.Print(fs)
}

func TestDecoder(t *testing.T) {
	b := make([]byte, 0)
	b = append(b, 0x16, 0x05, 0xb9, 0x09, 0x13, 0x02, 0x00, 0x00, 0x00, 0x01)

	var actual StaticsXY
	err := binstruct.UnmarshalLE(b, &actual) // UnmarshalLE() or Unmarshal()
	if err != nil {
		log.Fatal(err)
	}

	if actual.Tile != 1302 || actual.X != 2489 || actual.Y != 531 || actual.Color != 0 || actual.Z != 0 {
		t.Errorf("Failed to deserialize struct.\nReceived: %v", actual)
	}
}
