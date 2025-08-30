package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main1() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(30 * time.Microsecond):
				fmt.Println("continue")
			}
		}
	}()
	cancel()
	wg.Wait()
}

func main2() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("ctx timeout")
			return
		}
	}()
	wg.Wait()
}

func main3() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine exit")
		runtime.Goexit()
	}()
	wg.Wait()
}

func main4() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(150 * time.Millisecond):
			return
		}
	}()
	wg.Wait()
}

func main5() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	wg.Wait()
}

func main6() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("goroutine exit with done signal")
				return
			case <-time.After(30 * time.Microsecond):
				fmt.Println("continue")
			}
		}
	}()
	wg.Wait()
}
