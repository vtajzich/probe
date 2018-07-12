package main

import (
	"flag"
	"log"
)

func parseFlags() (Parameters) {

	socketServerUrlPtr := flag.String("socket-host", "localhost", "socket server hostname")
	socketServerPortPtr := flag.Int("socket-port", 4444, "socket server port")

	sourcePtr := flag.String("source", "", "Source url in format 'http://host:port/any'")
	statusCheckIntervalPtr := flag.Int("check-interval", 2, "status check interval in seconds")

	flag.Parse()

	socketServerUrl := *socketServerUrlPtr
	socketServerPort := *socketServerPortPtr

	source := *sourcePtr

	statusCheckInterval := *statusCheckIntervalPtr

	log.Println("Socket server ", socketServerUrl, ":", socketServerPort)
	log.Println("Source: ", source)
	log.Println("Check interval: ", statusCheckInterval)

	if source == "" {
		log.Fatal("Source must be set!")
	}

	return Parameters{socketServerUrl, socketServerPort, source, statusCheckInterval}
}
