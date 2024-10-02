package main

import (
	"fmt"
	"strings"
)

func main() {
	//Valid palindrome
	str := "Malayalam"
	isPal := isPalindrome(str)
	fmt.Println("String is palindrome", isPal)
	//Reverse string with two pointer approach
	str1 := "hello"
	revStr := toReverseString(str1)
	fmt.Println("After reverese the string", revStr)
	// First Non-Repating charcter in string
	str2 := "swiss"
	firstStr := nonRepatChar(str2)
	fmt.Println("First non repating char", firstStr)
	//Longest Substring Without Repeating Characters
	str3 := "aabcccccaaa"
	longSub := longestSubstring(str3)
	fmt.Println("Longest substring", longSub)

	//String builder is used to contact the string concatenation
	var sb strings.Builder
	for i := 0; i < 10; i++ {
		sb.WriteString("Go")
	}
	result := sb.String()
	fmt.Println("stirng is ", result)

	reve := "ratheesh"
	fmt.Println("Reverse", reverseStringRec(reve))
}

// If string is palindrome or not
func isPalindrome(s string) bool {
	tolow := toLower(s)
	str := []rune(tolow)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

// Convert the upper letter to small letter case
func toLower(s string) string {
	str := []rune(s)

	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			str[i] = char + 'a' - 'A'
		} else {
			str[i] = char
		}
	}
	return string(str)
}

// Reverse string - Two pointer approach
func toReverseString(s string) string {
	word := []rune(s)
	l := 0
	r := len(s) - 1

	for l < r {
		word[l], word[r] = word[r], word[l]
		l++
		r--
	}
	return string(word)
}

// First Non-Repating charcter in string
func nonRepatChar(s string) string {
	str := make(map[rune]int)
	//count the string
	for _, i := range s {
		str[i]++
	}

	//Find the first non - repating char
	for _, i := range s {
		if str[i] == 1 {
			return string(i)
		}
	}
	return ""
}

// Longest Substring Without Repeating Characters
func longestSubstring(s string) string {
	charInx := make(map[rune]int)
	maxlength := 0
	maxStr := ""
	left := 0

	for right, char := range s {
		if indx, exist := charInx[char]; exist {
			left = max(left, indx+1)
		}
		charInx[char] = right
		if right-left+1 > maxlength {
			maxlength = right - left + 1
			maxStr = s[left : right+1]
		}

	}
	return string(maxStr)
}

func reverseStringRec(s string) string {
	if len(s) <= 0 {
		return s
	}
	return reverseStringRec(s[1:] + string(s[0]))
}
