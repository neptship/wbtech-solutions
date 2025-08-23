package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	mainCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(mainCh)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			case mainCh <- counter:
				counter++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mainCh {
			fmt.Printf("worker received: %d\n", v)
		}
		fmt.Println("worker channel closed, exiting")
	}()

	N := 8 * time.Second
	<-time.After(N)
	cancel()

	wg.Wait()
}
