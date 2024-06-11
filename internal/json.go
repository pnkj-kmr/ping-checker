package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type _json struct {
	ifile   string
	ofile   string
	count   int
	timeout int
}

func NewJSON(ifile, ofile string, count, timeout int) PingChecker {
	return &_json{ifile, ofile, count, timeout}
}

func (c *_json) GetInput() (out []Input) {
	fileData, err := os.ReadFile(c.ifile)
	if err != nil {
		log.Fatal("Error while read the json file", err)
	}

	err = json.Unmarshal(fileData, &out)
	if err != nil {
		log.Fatal("Error while converting to json", err)
	}

	return out
}

func (c *_json) ProduceOutput(ch <-chan Output, exitCh chan<- struct{}) {
	var out []Output
	for r := range ch {
		out = append(out, r)
	}

	outJson, _ := json.Marshal(out)
	err := ioutil.WriteFile(fmt.Sprintf("%s", c.ofile), outJson, 0644)
	if err != nil {
		log.Fatal("Error while writing into file", err)
	}

	exitCh <- struct{}{}
}
