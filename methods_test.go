package gostealthclient_test

import (
	"log"
	"strings"
	"testing"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func TestInjournalBetweenTimes(t *testing.T) {
	t.Run("Should return -1 if string is not found", func(t *testing.T) {
		s := func() interface{} {
			sc.ClearJournal()
			tb, _ := time.Parse(time.RFC3339, "2020-12-30T00:00:00Z")
			return <-sc.InJournalBetweenTimes("test", tb, time.Now())
		}
		ans := sc.Bootstrap(s)
		res, ok := ans.(int32)
		if !ok || res != -1 {
			t.Fatalf("Returned %v, want -1", res)
		}
	})

	const str = "test"
	t.Run("Should return >= 0 if string is found", func(t *testing.T) {
		s := func() interface{} {
			tb := time.Now()
			sc.UOSay(str)
			time.Sleep(time.Second * 2)
			return <-sc.InJournalBetweenTimes(str, tb, time.Now())
		}

		ans := sc.Bootstrap(s)
		res, ok := ans.(int32)
		if !ok || res == -1 {
			t.Fatalf("Returned %v, want >0", res)
		}
	})
}

func TestJournalLine(t *testing.T) {
	t.Run("If line not found should return -1", func(t *testing.T) {
		s := func() interface{} {
			return <-sc.InJournal("{{)#)#)")
		}
		ans := sc.Bootstrap(s)
		res, ok := ans.(int32)
		if !ok || res != -1 {
			t.Errorf("InJournal('{{)#)#)') = %v; want -1", res)
		}
	})

	const str = "Hello from GO 123123qwe"
	t.Run("If line found should return >= 0", func(t *testing.T) {
		s := func() interface{} {
			sc.UOSay(str)
			time.Sleep(time.Second * 2)
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

func TestFindTypesArrayEx(t *testing.T) {
	t.Run("Should find item", func(t *testing.T) {
		s := func() interface{} {
			return <-sc.FindTypesArrayEx(
				[]uint16{0x190, 0x25d, 0x191, 0x25e},
				[]uint16{0xffff, 0x2222},
				[]uint32{0x0},
				true,
			)
		}
		ans := sc.Bootstrap(s)
		res, ok := ans.(uint32)
		if !ok || res == 0 {
			t.Errorf("Failed to find any items")
		}
	})
}

func TestGetTooltip(t *testing.T) {
	s := func() interface{} {
		return <-sc.GetTooltip(<-sc.Backpack())
	}
	ans := sc.Bootstrap(s)
	res, ok := ans.(string)
	if !ok || !strings.Contains(res, "backpack|Weight") {
		t.Errorf("Failed to get tooltip, received: %v, want %v", res, "backpack|Weight")
	}
}

func TestReadStaticXY(t *testing.T) {
	s := func() interface{} {
		return <-sc.ReadStaticsXY(<-sc.GetX(<-sc.Self()), <-sc.GetY(<-sc.Self()), <-sc.WorldNum())
	}
	ans := sc.Bootstrap(s)
	res, ok := ans.([]m.StaticsXY)
	if !ok || len(res) == 0 {
		t.Errorf("Failed to read statics")
	}
}

func TestIsWorldCellPassablePacket(t *testing.T) {
	s := func() interface{} {
		x, y := <-sc.GetX(<-sc.Self()), <-sc.GetY(<-sc.Self())
		z := <-sc.GetZ(<-sc.Self())
		return <-sc.IsWorldCellPassable(
			x,
			y,
			z,
			2547,
			535,
			<-sc.WorldNum(),
		)
	}

	ans := sc.Bootstrap(s)
	res, ok := ans.(m.WorldCellPassable)

	if !ok || res.Passable || res.Z != 0 {
		t.Errorf("Failed to resolve TestIsWorldCellPassablePacket %v res.", res)
	}
}

func TestGetZ(t *testing.T) {
	s := func() interface{} {
		return <-sc.GetZ(<-sc.Self())
	}

	ans := sc.Bootstrap(s)
	res, ok := ans.(int8)

	if !ok || res != 0 {
		t.Error(res, ok)
	}
}

func TestBuffBarInfo(t *testing.T) {
	s := func() interface{} {
		sc.CastToObj("Agility", <-sc.Self())
		time.Sleep(time.Second * 3)
		return <-sc.GetBuffBarInfo()
	}

	ans := sc.Bootstrap(s)
	res, ok := ans.(m.BuffBarInfo)
	if !ok || res.Count < 1 {
		t.Errorf("Failed to get buffbar, or char didn't cast Agility. Res: %v", res)
	}
}

func TestExtInfo(t *testing.T) {
	s := func() interface{} {
		return <-sc.GetExtInfo()
	}
	ans := sc.Bootstrap(s)
	res, ok := ans.(m.ExtendedInfo)

	if !ok || res.MaxWeight == 0 {
		t.Errorf("Failed to resolve Ext info. %v", res)
	}

}

func TestMethod_AddToSystemJournal(t *testing.T) {
	s := func() interface{} {
		sc.AddToSystemJournal("Hello World")
		return nil
	}
	sc.Bootstrap(s)
}

func TestUOSayColor(t *testing.T) {
	s := func() interface{} {
		sc.UOSayColor("test", 0x190)
		return nil
	}

	sc.Bootstrap(s)
}

func TestConnection(t *testing.T) {
	s := func() interface{} {
		sc.AddToSystemJournal("Hello World")
		return nil
	}

	for i := 0; i < 5; i++ {
		sc.Bootstrap(s)
		log.Println("=================")
		time.Sleep(time.Second * 2)
	}
}

// Failing to read properly value has x00/ in the beginning
func TestSetGetGlobalVar(t *testing.T) {
	s := func() interface{} {
		sc.SetGlobal(0, "testVar", "testValue")
		return <-sc.GetGlobal(0, "testVar")
	}

	ans := sc.Bootstrap(s)
	res, ok := ans.(string)

	if !ok || res != "testValue" {
		t.Errorf("Failed to set or get global var. Received: %v", res)
	}

}

func TestGetMapCell(t *testing.T) {
	type response struct {
		zOrig int8
		mCell m.MapCell
	}

	s := func() interface{} {
		return response{
			zOrig: <-sc.GetZ(<-sc.Self()),
			mCell: <-sc.GetCell(<-sc.GetX(<-sc.Self()), <-sc.GetY(<-sc.Self()), <-sc.WorldNum()),
		}
	}

	ans := sc.Bootstrap(s)
	res, ok := ans.(response)

	if !ok || res.zOrig != res.mCell.Z {
		t.Errorf("Failed to set or get global var. Received: %v", res)
	}
}

func TestClientTargetResponse(t *testing.T) {
	s := func() interface{} {
		sc.ClientRequestObjectTarget()
		// sc.ClientRequestTileTarget()
		<-sc.WaitForClientTargetResponse(time.Second * 10)
		return <-sc.ClientTargetResponse()
	}

	ans := sc.Bootstrap(s)
	res, ok := ans.(m.TargetInfo)

	if !ok {
		t.Errorf("Failed to set or get global var. Received: %v", res)
	}
}

func Test_StaticTilesArray(t *testing.T) {
	s := func() interface{} {
		s := <-sc.Self()
		x := <-sc.GetX(s)
		y := <-sc.GetY(s)

		tt := []uint16{0xFFFF}
		return <-sc.GetStaticTilesArray(x-2, y-2, x+2, y+2, <-sc.WorldNum(), tt)
	}

	ans := sc.Bootstrap(s)
	res, ok := ans.([]m.FoundTile)

	if !ok {
		t.Errorf("Failed to set or get global var. Received: %v", res)
	}
}
