package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInterruptLongRunningGoroutine(t *testing.T) {
	ctx := context.Background()
	// this test task is to start long task and interrupt it if it exceeded maximum
	// allowed timeout, there are 2 cases:
	// 1. In delay smaller than timeout inner task will be finished before <-context.Done signal
	// 2. If delay > timeout, task will be interrupted by signal from context channel
	longRunningFunc(ctx, 2000)
}

func longRunningFunc(ctx context.Context, timeout int) {
	// this function start function that could finished fast or be working very slow
	// for the cases we interrupt a long-running task by context limitation
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
	defer cancel()
	taskCompletedCh := make(chan int)
	go func(c context.Context) {
		fmt.Println("Long Running Func Started")
		delay := rand.Intn(4000)
		ch := make(chan int)
		go func() {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			ch <- 1
			fmt.Println("Long Running Func Finished")
		}()
		select {
		case _, ok := <-ch:
			if ok {
				fmt.Println("Long Running Func Completed before Timeout")
			}
		case <-ctxWithTimeout.Done():
			fmt.Println("Long Running Func Timed Out, interrupted!!!!!")
		}

		taskCompletedCh <- 1
	}(ctxWithTimeout)

	<-taskCompletedCh
}
