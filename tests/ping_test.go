package pinger_test

import (
	"fmt"
	"testing"

	"github.com/go-ping/ping"
)

func TestPing(t *testing.T) {
	pinger, err := ping.NewPinger("www.google.com")
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(stats)
}
