package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInputChannelsCombineIntoOne(t *testing.T) {

	// Function that combines data from multiple channels into single one,
	// Fan out pattern, it runes goRoutinesNum of goroutines and write after random delay
	// random amount of messages, after all it closes its channel
	// This is implementation of Fan Out pattern

	const goRoutinesNum = 3
	var channels []chan int
	combineChannel := make(chan int)

	for i := 0; i < goRoutinesNum; i++ {
		channels = append(channels, make(chan int))
	}

	go func() {
		combine(channels[0], channels[1], channels[2], combineChannel)
	}()

	for i := 0; i < goRoutinesNum; i++ {
		val := i
		go func(n int, ch chan int) {
			messageNumber := rand.Intn(10) + 2
			for c := 0; c < messageNumber; c++ {
				delay := rand.Intn(10) + 1
				time.Sleep(time.Duration(delay) * time.Millisecond)
				ch <- n + 1
			}
			close(ch)
		}(val, channels[i])
	}

	for {
		out, ok := <-combineChannel
		if !ok {
			fmt.Println("All input channels were closed")
			break
		}
		fmt.Printf("Received data from channel: %d\n", out)
	}
}

func combine(input0, input1, input2 chan int, output chan int) {
	done := 0
	for {
		select {
		case value0, ok := <-input0:
			if ok {
				output <- value0
			} else {
				done += 1
			}
		case value1, ok := <-input1:
			if ok {
				output <- value1
			} else {
				done += 1
			}
		case value2, ok := <-input2:
			if ok {
				output <- value2
			} else {
				done += 1
			}
		}
		if done == 3 {
			close(output)
			return
		}
	}

}
