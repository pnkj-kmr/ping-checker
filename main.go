package main

import (
	"flag"
	"fmt"
	"log"
	"ping-checker/internal"
	"sync"
)

func main() {
	fileName := flag.String("f", "input.csv", "give a file name")
	outFilename := flag.String("o", "output.csv", "output file name")
	noWorkers := flag.Int("w", 4, "number of worker")
	timeout := flag.Int("t", 5, "ping timeout - secs")
	flag.Parse()
	log.Println("File accepted:", *fileName, "| output file:", *outFilename)
	log.Println("Ping timeout:", *timeout, "| Worker processes:", *noWorkers)

	ips := internal.GetIPList(*fileName)
	log.Println("Total IPs -- ", len(ips))

	exitCh := make(chan struct{})
	ch := make(chan internal.Result)
	go internal.PutOutput(*outFilename, ch, exitCh)

	var wg sync.WaitGroup
	c := make(chan int, *noWorkers)
	for i := 0; i < len(ips); i++ {
		wg.Add(1)
		c <- 1
		go func(ip internal.Csv, i, t int) {
			defer func() { wg.Done(); <-c }()
			ok, loss, err := internal.Ping(ip.IP, t)
			if err == nil {
				err = fmt.Errorf("")
			}
			ch <- internal.Result{IP: ip.IP, Tag: ip.Tag, Ok: ok, PktLass: loss, Err: err}
		}(ips[i], i, *timeout)
		log.Println("IP sent for ping -- ", ips[i], i+1)
	}
	wg.Wait()
	close(ch)
	<-exitCh
	log.Println("Execution completed.")
}
