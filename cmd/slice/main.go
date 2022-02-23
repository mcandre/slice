package main

import (
	"github.com/mcandre/slice"

	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
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

	chIn, chOut, chDone := slice.Slice(rate, skip)
	defer func() { chDone <- struct{}{} }()

	go func() {
		for {
			fmt.Println(<-chOut)
		}
	}()

	args := flag.Args()
	var wg sync.WaitGroup
	argLen := len(args)

	if argLen == 0 {
		wg.Add(1)
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			chIn <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		wg.Done()
	} else {
		wg.Add(argLen)

		for _, arg := range args {
			reader, err := os.Open(arg)

			if err != nil {
				log.Fatal(err)
			}

			defer reader.Close()
			scanner := bufio.NewScanner(reader)

			go func(wg *sync.WaitGroup) {
				for scanner.Scan() {
					chIn <- scanner.Text()
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

				wg.Done()
			}(&wg)
		}
	}

	wg.Wait()
}
