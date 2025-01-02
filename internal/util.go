package internal

import (
	"flag"
	"log"
)

func GetCmdPipe() CmdPipe {
	version := flag.Bool("version", false, "Application version")
	fileName := flag.String("f", "input.csv", "give a file name")
	outFilename := flag.String("o", "output.csv", "output file name")
	noWorkers := flag.Int("w", 4, "number of workers")
	timeout := flag.Int("t", 5, "ping timeout [secs]")
	count := flag.Int("c", 4, "packet count")
	jsontype := flag.Bool("json", false, "file type (default csv)")

	flag.Parse()

	log.Println("File accepted:", *fileName, "| output file:", *outFilename)
	log.Println("timeout:", *timeout, "| workers:", *noWorkers, "| packet:", *count, "| json:", *jsontype)

	return CmdPipe{
		Ifile: *fileName, Ofile: *outFilename, Workers: *noWorkers, Count: *count,
		Timeout: *timeout, JsonType: *jsontype, Version: *version,
	}
}
