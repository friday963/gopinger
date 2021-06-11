package pingerfolder

import (
	// "fmt"

	"github.com/go-ping/ping"
)
	
func PingerThing() *ping.Statistics {
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
	return stats
}
