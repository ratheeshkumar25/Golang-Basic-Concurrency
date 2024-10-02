package main

import (
	"fmt"
	"sync"
)

func count(word string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	for i := 1; i < 20; i++ {
		c <- word
	}

}
func count1(word string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	for i := 1; i < 20; i++ {
		c <- word
	}

}

func main() {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(2)
	go func() {
		count("Dog", &wg, c)

	}()
	go func() {
		count1("cat", &wg, c)

	}()

	msg := <-c
	msg1 := <-c
	wg.Wait()
	fmt.Println(msg)
	fmt.Println(msg1)

}
