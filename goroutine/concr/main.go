package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10

	evenCh := make(chan int)
	oddCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go isPrime(n, evenCh, oddCh, &wg)

	go func() {
		wg.Wait()
		close(evenCh)
		close(oddCh)
	}()

	go func() {
		for even := range evenCh {
			fmt.Println("Even number is", even)
		}
	}()
	go func() {
		for odd := range oddCh {
			fmt.Println("Odd number is", odd)
		}
	}()

	wg.Wait()

}

func isPrime(n int, evenCh chan int, odd chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			odd <- i
		}
	}

}
