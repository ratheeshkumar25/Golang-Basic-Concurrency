package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Basic understanding how the go routine works
	go start()
	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finishes")

	arr := []int{10, 20, 30, 40, 50, 60}
	var wg sync.WaitGroup
	chsum := make(chan int, 2)
	mid := len(arr) / 2
	wg.Add(2)
	go func() {
		defer wg.Done()
		sumArr(arr[mid:], chsum)
	}()

	go func() {
		defer wg.Done()
		sumArr(arr[:mid], chsum)
	}()
	sum := <-chsum
	sum1 := <-chsum

	wg.Wait()
	close(chsum)

	if sum > sum1 {
		fmt.Println("Max sum is ", sum)
	} else {
		fmt.Println("Max sum is", sum1)
	}

	fmt.Println("Array sum is ", sum)
	fmt.Println("Array sum is half", sum1)

}

func sumArr(arr []int, c chan int) {

	var sum int
	for _, val := range arr {
		sum += val

	}
	c <- sum
}

func start() {
	fmt.Println("In goroutines")
}
