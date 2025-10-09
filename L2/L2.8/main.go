package main

import (
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	const ntpServer = "ntp.ultracoder.org" // Russian NTP server

	currentTime, err := ntp.Time(ntpServer)
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Printf("NTP query failed: %v", err)
		os.Exit(1)
	}

	log.Printf("Current time: %v", currentTime)
}
