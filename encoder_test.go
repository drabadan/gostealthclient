package gostealthclient_test

import (
	"fmt"
	"log"
	"math"
	"testing"
	"time"

	sc "github.com/drabadan/gostealthclient"
)

func readTimeResponse(f float64) time.Time {
	loc, _ := time.LoadLocation("Europe/Berlin")
	t := time.Date(1899, 12, 30, 00, 00, 00, 00, loc)

	d, _ := math.Modf(f)

	ds := fmt.Sprintf("%vh", (int32(d) * 24))
	dd, _ := time.ParseDuration(ds)
	hs := fmt.Sprintf("%vh", (24 * (f - d)))
	hh, _ := time.ParseDuration(hs)

	t = t.Add(dd).Add(hh)

	return t
}

func TestEncodeTime(t *testing.T) {
	tt := time.Unix(1643107323, 0)
	b := make([]byte, 0)
	sc.EncodeTime(tt, &b)

	tt1 := readTimeResponse(44586.6147187569)
	tt2 := readTimeResponse(44586.61473034962)

	log.Print(tt1, tt2)

	t.Log(tt)
}
