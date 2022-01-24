package gostealthclient_test

import (
	"log"
	"testing"
	"time"

	sc "github.com/drabadan/gostealthclient"
)

func testScript() interface{} {
	// log.Printf("% x", <-stealth.Self())
	for i := 0; i < 1; i++ {
		time.Sleep(time.Second)
		log.Print(sc.ProfileName())
	}

	return 0
}

func TestMain(t *testing.T) {
	sc.Bootstrap(testScript)
}
