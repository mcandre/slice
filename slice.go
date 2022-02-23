package slice

import (
	"math/rand"
)

// Version is semver.
const Version = "0.0.1"

// Slice samples text.
//
// rate specifies the probability of preserving each line.
//
// Returns an input channel for submitting population lines;
// an output channel for receiving sample lines;
// and a done channel for concluding the sampling operation.
func Slice(rate float64) (chan<- string, <-chan string, chan<- struct{}) {
	chIn := make(chan string)
	chOut := make(chan string)
	chDone := make(chan struct{})

	go func() {
		var line string

		for {
			select {
			case <-chDone:
				break
			case line = <-chIn:
				if rand.Float64() < rate {
					chOut<- line
				}
			}
		}
	}()

	return chIn, chOut, chDone
}
