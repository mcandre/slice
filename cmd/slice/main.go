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

var flagRate = flag.Float64("rate", 0.1, "Probability of preserving each line, in (0.0, 1.0]")
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

	rate := *flagRate

	chIn, chOut, chDone := slice.Slice(rate)
	defer func() { chDone <- struct{}{} }()

	reader := os.Stdin

	args := flag.Args()

	if len(args) > 0 {
		r, err := os.Open(args[0])

		if err != nil {
			log.Fatal(err)
		}

		defer r.Close()
		reader = r
	}

	go func() {
		for {
			fmt.Println(<-chOut)
		}
	}()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		chIn <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
