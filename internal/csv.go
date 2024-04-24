package internal

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Csv represent the input
type Csv struct {
	IP  string
	Tag string
}

// Result represent the output
type Result struct {
	IP  string
	Tag string
	Ok  bool
	Err error
}

// GetIPList helps get IP list from csv
func GetIPList(f string) []Csv {
	file, err := os.Open(f)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error while Error reading records - ", err)
	}

	var out []Csv
	for _, r := range records {
		if len(r) > 1 {
			out = append(out, Csv{r[0], r[1]})
		} else if len(r) > 0 {
			out = append(out, Csv{r[0], ""})
		}
		// fmt.Println(out)
	}
	return out
}

// PutOutput helps to write result into file
func PutOutput(ch <-chan Result) {
	file, err := os.Create("./output.csv")
	if err != nil {
		log.Fatal("Unable to write into file -", err)
	}
	defer file.Close()
	file.Write([]byte("ip,tag,ping_result,error_if_any\n"))
	for r := range ch {
		// log.Println("-------", r)
		file.Write([]byte(fmt.Sprintf("%s,%s,%t,%s\n", r.IP, r.Tag, r.Ok, r.Err.Error())))
	}
}
