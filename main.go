package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	"os"
	"ping_destination/pingerfolder"
	"strconv"
	"time"
)

func main() {
	var d = 3000 * time.Microsecond
	var t = time.Now().Add(d)
	i := 1
	for i <= 5 {
		if time.Now().Before(t) {
			continue
		}
		raw_result := pingerfolder.PingerThing()
		sent := strconv.Itoa(raw_result.PacketsSent)
		recv := strconv.Itoa(raw_result.PacketsSent)
		addr := raw_result.Addr

		fmt.Println("Sent = " + sent + " Received = " + recv + " Address = " + addr)
		i++
		f, err := os.OpenFile("results.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString("Sent = " + sent + " Received = " + recv + " Address = " + addr + "\n"); err != nil {
			log.Println(err)
		}
	}
}
