package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func Factorial(n int) int {
	k := 1
	for i := 1; i <= n; i++ {
		k = k * i
	}
	return k
}

func occurence(s string, char byte) string {
	if len(s) == 0 {
		return s
	}

	if s[0] == char {
		return occurence(s[1:], char)
	}
	return string(s[0]) + occurence(s[1:], char)
}

func main() {
	fmt.Println("Factorial of number", fact(5))
	fmt.Println("Factorial of number with normal function", Factorial(5))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(" Fibonocy of number of ", fib(7))

	str := "ratheesh"

	removeLett := occurence(str, 'e')
	fmt.Println(removeLett)
}
