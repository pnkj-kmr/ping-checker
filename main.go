package main

import (
	_ "embed"
	"fmt"
	"log"
	"ping-checker/internal"
	"sync"
	"time"
)

//go:embed version.txt
var appVersion string

func main() {
	st := time.Now()

	cmdPipe := internal.GetCmdPipe()
	if cmdPipe.Version {
		log.Println("Version: ", appVersion)
		return
	}

	var pinger internal.PingChecker
	if cmdPipe.JsonType {
		pinger = internal.NewJSON(cmdPipe.Ifile, cmdPipe.Ofile, cmdPipe.Count, cmdPipe.Timeout)
	} else {
		pinger = internal.NewCSV(cmdPipe.Ifile, cmdPipe.Ofile, cmdPipe.Count, cmdPipe.Timeout)
	}
	ips := pinger.GetInput()
	log.Println("Total IPs -- ", len(ips))

	exitCh := make(chan struct{})
	ch := make(chan internal.Output, cmdPipe.Workers)
	go pinger.ProduceOutput(ch, exitCh)

	var wg sync.WaitGroup
	c := make(chan int, cmdPipe.Workers)
	for i := 0; i < len(ips); i++ {
		wg.Add(1)
		c <- 1
		go func(ip internal.Input, count, timeout int) {
			defer func() { wg.Done(); <-c }()
			out, err := internal.Ping(ip, count, timeout)
			if err == nil {
				err = fmt.Errorf("")
				out.Err = err.Error()
			}
			ch <- out
		}(ips[i], cmdPipe.Count, cmdPipe.Timeout)
		log.Println("IP sent for ping -- ", ips[i], i+1)
	}
	wg.Wait()
	close(ch)
	<-exitCh
	log.Println("Execution completed. time taken", time.Since(st))
}
