package main

import "fmt"

func isSubString(str string, str1 string) bool {
	m, n := len(str), len(str1)
	for i := 0; i <= m-n; i++ {
		j := 0
		for j < n && str[i+j] == str1[j] {
			j++
		}
		if j == n {
			return true
		}

	}
	return false
}

func main() {
	str := "ratheesh"
	str1 := "rathee"
	isSub := isSubString(str, str1)

	fmt.Println("is substring", isSub)

}
