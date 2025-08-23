package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func workers(ctx context.Context, id int, mainCh <-chan int, doneCh chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
		case value, ok := <-mainCh:
			if !ok {
				fmt.Printf("Worker %d channel closed, exiting\n", id)
				doneCh <- struct{}{}
				return
			}
			fmt.Printf("Worker %d received: %d\n", id, value)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("use: go run main.go (workers)")
		return
	}
	workerCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("err")
		return
	}
	if workerCount <= 0 {
		fmt.Println("err")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	mainCh := make(chan int)
	doneCh := make(chan struct{})

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < workerCount; i++ {
		go workers(ctx, i, mainCh, doneCh)
	}

	go func() {
		counter := 0
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Main goroutine exiting")
				close(mainCh)
				return

			case mainCh <- counter:
				counter++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	<-signalCh
	fmt.Println("Received shutdown signal")
	cancel()

	for i := 0; i < workerCount; i++ {
		<-doneCh
	}

	fmt.Println("Graceful shutdown complete")
}
