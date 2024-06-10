package internal

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type _csv struct {
	ifile   string
	ofile   string
	count   int
	timeout int
}

func NewCSV(ifile, ofile string, count, timeout int) PingChecker {
	return &_csv{ifile, ofile, count, timeout}
}

func (c *_csv) GetInput() (out []Input) {
	file, err := os.Open(c.ifile)

	if err != nil {
		log.Fatal("Error while opening the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error while reading records", err)
	}

	for _, r := range records {
		if len(r) > 1 {
			out = append(out, Input{r[0], r[1], c.count, c.timeout})
		} else if len(r) > 0 {
			out = append(out, Input{r[0], "", c.count, c.timeout})
		}
	}
	return out
}

func (c *_csv) ProduceOutput(ch <-chan Output, exitCh chan<- struct{}) {
	file, err := os.Create(fmt.Sprintf("./%s", c.ofile))
	if err != nil {
		log.Fatal("Error while writing into file", err)
	}
	defer file.Close()

	file.Write([]byte("ip,tag,result,packetloss,rtt,stddevrrt,error_if_any\n"))
	for r := range ch {
		file.Write([]byte(fmt.Sprintf("%s,%s,%t,%f,%d,%d,%s\n",
			r.I.IP, r.I.Tag, r.Ok, r.PacketLoss, r.AvgRtt, r.StdDevRtt,
			r.Err)))
	}

	exitCh <- struct{}{}
}
