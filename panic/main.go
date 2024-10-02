package main

import "fmt"

// func handle() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recover from panic", r)
// 	}
// }

func main() {
	var handle = func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic")
		}
	}
	defer handle()
	panic("Something went wrong")

}
