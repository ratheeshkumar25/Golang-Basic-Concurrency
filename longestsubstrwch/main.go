package main

import "fmt"
func main(){
	s := "abcabcbb"

	var longestSubstring = func(s string)int{
		str := make(map[rune]bool)
		i := 0
		res := 0

		for j,sRune := range s{
			for str[sRune] {
				str[rune(s[i])] = false
				i++
			}
			str[sRune] = true 
			if j-i+1 > res{
				res = j-i+1
			}
		}
		return res
	}

	longSub := longestSubstring(s)
	fmt.Println(longSub)
}