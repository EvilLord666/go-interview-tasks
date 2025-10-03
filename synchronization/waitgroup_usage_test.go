package synchronization

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroupGoroutinesCancelIfOneFails(t *testing.T) {
	// Task : Run a number of goroutines and interrupt all if one fails
	wg := sync.WaitGroup{}
	// interruptChan := make(chan int)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	for i := 0; i < 5; i++ {
		val := i
		wg.Add(1)
		go func(c int) {
			//var err error
			completedCh := make(chan bool)
			go func() {
				fmt.Printf("Started task : %d \n", c)
				time.Sleep(time.Duration(100) * time.Millisecond)
				if val == 3 {
					time.Sleep(time.Duration(500) * time.Millisecond)
					fmt.Printf("Interruptted task : %d \n", c)
					cancelFunc()
					return
				}
				time.Sleep(time.Duration(c*400) * time.Millisecond)
				completedCh <- true
				fmt.Printf("Finished task : %d \n", c)
			}()

			select {
			case <-cancelCtx.Done():
				wg.Done()
				break
			case <-completedCh:
				wg.Done()
				break
			}

		}(val)
	}
	wg.Wait()
}
