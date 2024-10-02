package main

import (
	"fmt"
)

func main() {
	s := "anagram"
	t := "nagarm"

	var isAnagram = func(s, t string) bool {
		if len(s) != len(t) {
			return false
		}

		count := make(map[byte]int)

		for i := range s {
			count[s[i]]++
		}

		for i := range t {
			count[t[i]]--
			if count[t[i]] < 0 {
				return false
			}
		}
		return true
	}

	checkAngaram := isAnagram(s, t)
	fmt.Println("Given string is ", checkAngaram)
}
