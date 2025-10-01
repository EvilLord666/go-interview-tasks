package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCombineAllChannel(t *testing.T) {

	// Function that combines

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
			delay := rand.Intn(30)
			time.Sleep(time.Duration(delay) * time.Millisecond)
			ch <- n + 1
		}(val, channels[i])
	}

	for i := 0; i < goRoutinesNum; i++ {
		out := <-combineChannel
		fmt.Println(out)
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
