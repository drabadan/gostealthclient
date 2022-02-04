package gostealthclient

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func DecodeUtf16(inputBytes []byte) string {
	win16be := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(inputBytes[6:]), utf16bom)

	// decode and print:
	decoded, err := ioutil.ReadAll(unicodeReader)
	if err != nil {
		log.Fatal("Failed to parse string from packet...")
	}

	return strings.Replace(string(decoded), "\r\n", "\n", -1)
}
