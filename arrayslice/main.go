package main

import (
	"fmt"
)

func main() {

	//Array :- An array is a collection of elements that belong to the same type. For example, the collection of integers 5, 8, 9, 79, 76 forms an array

	var arr [3]int //int array with length 3
	//intialize the value in array
	arr[0] = 50
	arr[1] = 60
	arr[2] = 70
	fmt.Println(arr)

	//short hand declaration to create array
	arr1 := [4]int{2, 3, 4, 5}
	arr2 := [3]int{12}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// makes the compiler determine the length
	arr3 := [...]int{22, 23, 24, 25}
	fmt.Println("Array with length of:", arr3, len(arr3))

	// array are the value types :Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable,
	//a copy of the original array is assigned to the new variable.
	//If changes are made to the new variable, it will not be reflected in the original array.
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is", b)
	//Sum of array
	arr4 := [...]float64{67.7, 75.5, 88.80, 21.8, 30.8}
	sum := float64(0)

	for i, v := range arr4 {
		fmt.Printf("%d the element of arr is %f\n", i, v)
		sum += v
	}
	fmt.Println("\n sum of all element of array", sum)
	// n := 10
	// recu(n)
	// fmt.Print(n)
	//Map
	myApp := make(map[string]int)
	myApp["Apple"] = 5
	fmt.Println(myApp["Apple"])
	//If else standard conditional logical in go
	x := 6
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is 5 or less")
	}

	//goroutine and channel
	ch := make(chan string)
	go func() {
		ch <- "Running go routine"
	}()
	receive := <-ch
	fmt.Println(receive)

	//for loop in golang

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j) //Output 0,1,2
	}

	for i := range 3 {
		fmt.Println(i) // Output 0,1,2
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//Twod array
	var twoDarr [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDarr[i][j] = i + j
		}
	}
	fmt.Println("Two D Array", twoDarr)

	twoDarr = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("2d", twoDarr)
	fmt.Println("len of array I", len(twoDarr), cap(twoDarr))

	//Slice

	var s []string
	fmt.Println("uint", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, "len", len(s), "cap", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("after appending s", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)
	l := s[1:2]
	fmt.Println("sl1", l)
	l = s[1:]
	fmt.Println("sl1", l)
	l = s[:2]
	fmt.Println("sl1", l)
	l = s[2:]
	fmt.Println("sl1", l)

	//Twod Slice
	twoDslice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		interLen := i + 1
		twoDslice[i] = make([]int, interLen)
		for j := 0; j < interLen; j++ {
			twoDslice[i][j] = i + j
		}
	}
	fmt.Println("twoDslice", twoDslice)

	//Maps - maps are inbuilt datatype or datastructure (sometimes its called hashmap or dic in other langauage)

	mp := make(map[string]int)
	mp["k1"] = 7
	mp["k3"] = 13
	fmt.Println("map is", mp)

	v1 := mp["k1"]
	fmt.Println("v1 is", v1)
	fmt.Println("length of mp", len(mp))
	delete(mp, "k3")
	fmt.Print("mp is", mp)
	clear(mp)
	fmt.Println(mp)
	_, prs := mp["k3"]
	fmt.Println("prs", prs)
	vardic(123, 123)

	nestedDict := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"address": map[string]interface{}{
			"city":  "New York",
			"state": "NY",
		},
		"contacts": []interface{}{
			map[string]interface{}{
				"type":  "email",
				"value": "john.doe@example.com",
			},
			map[string]interface{}{
				"type":  "phone",
				"value": "123-456-7890",
			},
		},
	}

	fmt.Printf("%v+n", nestedDict)

	address := nestedDict["address"].(map[string]interface{})
	address["city"] = "Log Angles"
	contact := nestedDict["contacts"].([]interface{})
	email := contact[0].(map[string]interface{})
	email["value"] = "test@gmail.com"
	phone := contact[1].(map[string]interface{})
	phone["value"] = "9961429910"

	fmt.Printf("%v+n", nestedDict)

	//range
	nums2 := []int{5, 6, 7, 8, 9}
	sum2 := 0
	for _, num := range nums2 {
		sum2 += num
	}
	fmt.Println("Sum of Array", sum2)

	for i, num := range nums2 {
		if num == 7 {
			fmt.Println("index", i)
		}
	}
	// Type assertion
	var h interface{} = "hello"

	if l, ok := h.(int); ok {
		fmt.Println(l)
	} else {
		fmt.Println("type assertion not happening ")
	}

}

func vardic(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("total is", total)

}
