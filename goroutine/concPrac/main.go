package main

import (
	"fmt"
	"time"
)

// Basic understanding of goroutine
func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHello function ended successfully")
}
func sayHi() {
	fmt.Println("Hi Ratheesh")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("learning goroutine")
	go sayHello()
	sayHi()
	//Wait for a moment to allow goroutine to finish
	time.Sleep(2000 * time.Millisecond)
	//Output := learning gorountine , Hi Ratheesh , Hello, say hello function ended successfully
}
