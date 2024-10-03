package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	START_PORT = 1
	END_PORT   = 65535
	RESET      = "\033[0m"
	RED        = "\033[31m"
)

type PortsOpened struct {
	Type string
	Port int
}

func main() {

	var host string
	var protocol string
	var maxWorkers int
	var requestTimeout int

	flag.StringVar(&host, "host", "google.com", "Destination host")
	flag.StringVar(&protocol, "protocol", "tcp", "Protocol to use")
	flag.IntVar(&requestTimeout, "requestTimeout", 2, "Duration in seconds for response")
	flag.IntVar(&maxWorkers, "maxWorkers", 100, "Number of workers to run in given period")

	flag.Parse()
	fmt.Println("Host: ", host)
	fmt.Println("Protocol: ", protocol)
	fmt.Println("Max Workers: ", maxWorkers)
	fmt.Println("Timeout: ", requestTimeout)

	portsChan := make(chan int)
	resultsChan := make(chan PortsOpened)

	var wg sync.WaitGroup

	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go func() {
			Scan(protocol, host, portsChan, resultsChan, requestTimeout)
			defer wg.Done()
		}()
	}

	go func() {
		for i := START_PORT; i <= END_PORT; i++ {
			portsChan <- i
		}
		close(portsChan)
	}()

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		fmt.Println(RED, "Port is open: ", result.Port, "Protocol: ", result.Type, RESET)
	}

}

func Scan(protocol string, host string, portsChan chan int, results chan PortsOpened, requestTimeout int) {
	for port := range portsChan {
		conn, err := net.DialTimeout(protocol, host+":"+strconv.Itoa(port), time.Second*time.Duration(requestTimeout))
		if err == nil {
			results <- PortsOpened{Type: protocol, Port: port}
			conn.Close()
		}
	}
}
