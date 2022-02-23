package main

import (
	"github.com/mcandre/slice"

	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var flagRate = flag.Float64("rate", slice.DefaultRate, "Probability of preserving each line, in (0.0, 1.0]")
var flagSkip = flag.Int64("skip", 0, "Deterministically skip every nth line. Disables rate.")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Printf("slice %v\n", slice.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(0)
	}

	rand.Seed(time.Now().UnixNano())

	rate := flagRate
	var skip *int64

	if flagSkip != nil && *flagSkip != 0 {
		rate = nil
		skip = flagSkip
	}

	args := flag.Args()
	var scanners []*bufio.Scanner

	if len(args) == 0 {
		scanners = append(scanners, bufio.NewScanner(os.Stdin))
	} else {
		for _, arg := range args {
			reader, err := os.Open(arg)

			if err != nil {
				log.Fatal(err)
			}

			defer reader.Close()
			scanners = append(scanners, bufio.NewScanner(reader))
		}
	}

	chIn, chOut, chDone := slice.Slice(rate, skip)
	defer func() { chDone <- struct{}{} }()

	go func() {
		for {
			fmt.Println(<-chOut)
		}
	}()

	for _, scanner := range scanners {
		for scanner.Scan() {
			chIn <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
