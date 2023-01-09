package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := &sync.RWMutex{}
	wc := &sync.WaitGroup{}
	stream := []int{0}
	wc.Add(3)
	go func(mu *sync.RWMutex, wc *sync.WaitGroup) {
		defer wc.Done()
		fmt.Println("In goroutines ONE")
		mu.Lock()
		stream = append(stream, 1)
		mu.Unlock()
	}(mu, wc)
	go func(mu *sync.RWMutex, wc *sync.WaitGroup) {
		defer wc.Done()
		fmt.Println("In goroutines TWO")
		mu.Lock()
		stream = append(stream, 2)
		mu.Unlock()
	}(mu, wc)
	go func(mu *sync.RWMutex, wc *sync.WaitGroup) {
		defer wc.Done()
		fmt.Println("In goroutines THREE")
		mu.Lock()
		stream = append(stream, 3)
		mu.Unlock()
	}(mu, wc)
	wc.Wait()
	mu.RLock()
	fmt.Println(stream)
	mu.RUnlock()
}
