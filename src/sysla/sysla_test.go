package main

import (
	"testing"
)

func getParser(buffer string) *Sysla {
	s := &Sysla{Buffer: string(buffer)}
	s.Init()
	return s
}

func testParse(buffer string, t *testing.T, negative bool) {
	p := getParser(buffer)
	err := p.Parse()
	if !negative {
		if err != nil {
			t.Error(err)
		}
	} else {
		if err == nil {
			t.Error(err)
		}
	}
}

func TestInit(t *testing.T) {
	buffer := ""
	s := &Sysla{Buffer: string(buffer)}
	s.Init()
}

func TestParseForm1(t *testing.T) {
	testParse("The interface \"eth0\" has \"netmask\" matching /\\d+/.", t, false)
	testParse("The interface \"eth1\" has \"mac_address\" equals \"12:34:56:78:9A:BC\".", t, false)
	testParse("The device \"/dev/sd1\" have \"free_space_bytes\" more than 30000000.", t, false)
	testParse("All interfaces have \"address\" matching  /.*/.", t, false)
}
