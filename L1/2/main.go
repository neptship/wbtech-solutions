package main

import (
	"fmt"
	"os"
	"sync"
)

func SquareGoroutines(nums []int) {
	wg := &sync.WaitGroup{}
	result := make([]int, len(nums))
	for i, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[i] = num * num
		}()
	}
	wg.Wait()
	for _, n := range result {
		fmt.Fprintf(os.Stdout, "%d ", n)
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	SquareGoroutines(nums)
}
