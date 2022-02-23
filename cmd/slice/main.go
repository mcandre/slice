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

const DefaultRate = float64(0.1)

var flagRate = flag.Float64("rate", DefaultRate, fmt.Sprintf("Probability of preserving each line, in (0.0, 1.0]. Default: %v", DefaultRate))
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
	defer func() { chDone<-struct{}{} }()

	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for {
			fmt.Println(<-chOut)
		}
	}()

	for scanner.Scan() {
		chIn<-scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
