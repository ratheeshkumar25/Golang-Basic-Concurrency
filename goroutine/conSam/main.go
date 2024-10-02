package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := []int{1, 2, 3, 4, 5, 6}
	result := make(chan int, len(arr))

	for _, d := range arr {
		wg.Add(1)
		go process(d, result, &wg)
	}
	wg.Wait()
	close(result)
	for resul := range result {
		fmt.Println("result", resul)
	}

}

func process(data int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	process := data * 2
	result <- process
}
