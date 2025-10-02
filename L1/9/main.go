package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 4, 8, 16, 32, 64, 128}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch1)
		for _, val := range x {
			ch1 <- val
		}
	}()
	go func() {
		defer close(ch2)
		for _, val := range x {
			ch2 <- val * 2
		}
	}()

	for val := range ch2 {
		fmt.Printf("%v ", val)
	}
}
