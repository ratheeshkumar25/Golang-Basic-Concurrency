package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println(msg)

}

func printMessage2(msg string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		ch <- msg
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(4)
	go printMessage("Hello from thread1", &wg)
	go printMessage("Hello from thread2", &wg)
	go printMessage2("Message1", ch, &wg)
	go printMessage2("Message2", ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i, <-ch)
	}
	// wg.Wait()

}
