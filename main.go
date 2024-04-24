package main

import (
	"flag"
	"fmt"
	"log"
	"ping-poller/internal"
	"sync"
	"time"
)

func main() {
	fileName := flag.String("f", "", "give a file name")
	noWorkers := flag.Int("w", 4, "number of worker")
	timeout := flag.Int("t", 5, "ping timeout - secs")
	flag.Parse()
	if *fileName == "" {
		*fileName = "./input.csv"
	}
	log.Println("File accepted:", *fileName)
	log.Println("Ping timeout:", *timeout, "| Worker processes:", *noWorkers)

	// var ips []string
	ips := internal.GetIPList(*fileName)
	log.Println("Total IPs -- ", len(ips))

	ch := make(chan internal.Result)
	go internal.PutOutput(ch)

	var wg sync.WaitGroup
	c := make(chan int, *noWorkers)
	for i := 0; i < len(ips); i++ {
		wg.Add(1)
		c <- 1
		go func(ip internal.Csv, i, t int) {
			defer func() { wg.Done(); <-c }()
			// fmt.Println("ip--- ", ip)
			ok, err := internal.Ping(ip.IP, t)
			if err == nil {
				err = fmt.Errorf("")
			}
			// fmt.Println("ip--- ", ip, ok, err)
			ch <- internal.Result{IP: ip.IP, Tag: ip.Tag, Ok: ok, Err: err}
		}(ips[i], i, *timeout)
		log.Println("IP sent for ping -- ", ips[i], i+1)
	}
	wg.Wait()
	time.Sleep(time.Second * 5)
	close(ch)
	log.Println("Execution completed.")
}
