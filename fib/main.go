package main

import "fmt"

func fibRec(n int) int {

	if n < 2 {
		return n
	}
	a, b := fibRec(n-1), fibRec(n-2)
	return a + b
}

func fibRec1(m int) int {
	//fmt.Println(m)
	if m <= 0 || m <= 1 {
		return m
	}
	return fibRec1(m-1) + fibRec1(m-2)
}

//without recursion fibSeries

func fibSeries(num int) int {
	seq := []int{0, 1}

	for len(seq) <= num {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
	}
	return seq[len(seq)-1]
}

func main() {
	n := 9
	m := 6
	k := 7
	fmt.Println("fib normal series", fibSeries(k))
	fmt.Println("fibseries", fibRec(n))
	fmt.Println("fibseries", fibRec1(m))
}
