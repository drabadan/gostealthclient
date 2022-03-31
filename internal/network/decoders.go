package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"

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

func DecodeDelphiTime(double float64) time.Time {
	f := double
	t := time.Date(1899, 12, 30, 00, 00, 00, 00, time.UTC)

	d, _ := math.Modf(f)

	ds := fmt.Sprintf("%vh", (int32(d) * 24))
	dd, _ := time.ParseDuration(ds)
	hs := fmt.Sprintf("%vh", (24 * (f - d)))
	hh, _ := time.ParseDuration(hs)

	t = t.Add(dd).Add(hh)

	return t
}
