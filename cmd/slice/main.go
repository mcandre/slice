package main

import (
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

func main() {
	flag.Parse()

	switch {
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	rate := *flagRate
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if rand.Float64() < rate {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
