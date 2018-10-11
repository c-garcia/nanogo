package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const InitialMsg = "Attempting connection to %s:%d. Max: %d attempts. Waiting; %d\n"

func main() {
	var host *string = flag.String("host", "localhost", "host to connect to")
	var port *int = flag.Int("port", 5432, "port to connect to")
	var retries *int = flag.Int("retries", 10, "number of retries")
	var delay *int = flag.Int("delay", 1, "seconds to wait between retries")
	flag.Parse()

	log.Printf(InitialMsg, *host, *port, *retries, *delay)

	for attempt := 0; attempt < *retries; attempt++ {
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
		if err == nil {
			log.Printf("Connection failed, attempt: %d\n", attempt+1)
			<-time.After(time.Duration(*delay) * time.Second)
		} else {
			log.Printf("Connection made")
			os.Exit(0)
		}
	}
	log.Fatalf("Giving up")
	os.Exit(1)
}
