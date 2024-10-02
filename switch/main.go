package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("This is weekend")
	default:
		fmt.Println("This weekday")
	}

	whatIam := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Dont know the type %T", t)
		}
	}
	whatIam(true)
	whatIam(1)
	whatIam("heye")

}
