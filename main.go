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
	noWorkers := flag.Int("w", 4, "number of workers")
	timeout := flag.Int("t", 5, "ping timeout [secs]")
	count := flag.Int("c", 4, "packet count")
	jsontype := flag.Bool("json", false, "file type - default[csv]")
	flag.Parse()
	log.Println("File accepted:", *fileName, "| output file:", *outFilename)
	log.Println("Ping timeout:", *timeout, "| Worker processes:", *noWorkers)

	var pinger internal.PingChecker
	if *jsontype {
		pinger = internal.NewJSON(*fileName, *outFilename, *count, *timeout)
	} else {
		pinger = internal.NewCSV(*fileName, *outFilename, *count, *timeout)
	}
	ips := pinger.GetInput()
	log.Println("Total IPs -- ", len(ips))

	exitCh := make(chan struct{})
	ch := make(chan internal.Output, *noWorkers)
	go pinger.ProduceOutput(ch, exitCh)

	var wg sync.WaitGroup
	c := make(chan int, *noWorkers)
	for i := 0; i < len(ips); i++ {
		wg.Add(1)
		c <- 1
		go func(ip internal.Input, count, timeout int) {
			defer func() { wg.Done(); <-c }()
			out, err := internal.Ping(ip, count, timeout)
			if err == nil {
				err = fmt.Errorf("")
				out.Err = err
			}
			ch <- out
		}(ips[i], *count, *timeout)
		log.Println("IP sent for ping -- ", ips[i], i+1)
	}
	wg.Wait()
	close(ch)
	<-exitCh
	log.Println("Execution completed.")
}
