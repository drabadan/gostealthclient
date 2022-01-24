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

func TestJournalLine(t *testing.T) {
	t.Run("If line not found should return -1", func(t *testing.T) {
		s := func() interface{} {
			return <-sc.InJournal("{{)#)#)")
		}
		ans := sc.Bootstrap(s)
		res, ok := ans.(int32)
		if !ok || res != -1 {
			t.Errorf("InJournal('hello') = %v; want -1", res)
		}
	})

	const str = "Hello from GO 123123qwe"
	t.Run("If line found should return >= 0", func(t *testing.T) {
		s := func() interface{} {
			sc.UOSay(str)
			return <-sc.InJournal(str)
		}
		ans := sc.Bootstrap(s)
		res, ok := ans.(int32)
		if !ok || res == -1 {
			t.Errorf("InJournal('%v') = %v; want >= 0", str, res)
		}
	})
}

func TestChangeProfile(t *testing.T) {
	t.Run("Change profile should return -4", func(t *testing.T) {
		s := func() interface{} {
			return <-sc.ChangeProfile("UNKNOWN")
		}
		ans := sc.Bootstrap(s)
		res, ok := ans.(int32)
		if !ok || res == -1 {
			t.Errorf("ChangeProfile returned = %v; want -4", ans)
		}
	})
}
