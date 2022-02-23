package slice

import (
	"math/rand"
)

// Version is semver.
const Version = "0.0.2"

// DefaultRate controls the normal probability preservation rate of each line.
const DefaultRate = float64(0.1)

// Slice samples text.
//
// rate specifies the probability of preserving each line.
//
// Sampling is tuneable via the Seed function from math/rand.
//
// Returns an input channel for submitting population lines;
// an output channel for receiving sample lines;
// and a done channel for concluding the sampling operation.
func Slice(rate *float64, skip *int64) (chan<- string, <-chan string, chan<- struct{}) {
	chIn := make(chan string)
	chOut := make(chan string)
	chDone := make(chan struct{})

	go func() {
		i := int64(1)
		var line string

		for {
			select {
			case <-chDone:
				break
			case line = <-chIn:
				if skip == nil {
					if rand.Float64() < *rate {
						chOut <- line
					}
				} else {
					if i == *skip {
						i = 1
					} else {
						chOut <- line
						i++
					}
				}
			}
		}
	}()

	return chIn, chOut, chDone
}
