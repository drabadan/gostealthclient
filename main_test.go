package gostealthclient_test

import (
	"log"
	"testing"
	"time"
)

func testScript() {
	// log.Printf("% x", <-stealth.Self())
	for i := 0; i < 1; i++ {
		time.Sleep(time.Second)
		log.Print(<-stealth.ProfileName())
	}
}

func TestMain(t *testing.T) {
	stealth.Bootstrap(testScript)

}
