package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func (s *SafeMap) Get(k string) (int, bool) {
	s.mu.RLock()
	value, ok := s.m[k]
	s.mu.RUnlock()
	return value, ok
}

func (s *SafeMap) Set(k string, v int) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *SafeMap) Inc(k string) {
	s.mu.Lock()
	s.m[k]++
	s.mu.Unlock()
}

func main() {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				sm.Inc("counter")
			}
		}(i)
	}

	wg.Wait()

	if v, ok := sm.Get("counter"); ok {
		fmt.Println("counter =", v)
	}
}
