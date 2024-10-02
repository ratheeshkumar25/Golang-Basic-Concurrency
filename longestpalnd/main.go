package main

import (
	"fmt"
)

func main() {

	s := "A man, a plan, a canal: Panama"
	isPalindrome := isValidPalindrome(s)
	fmt.Println("String is", isPalindrome)

}

func isValidPalindrome(s string) bool {
	newStr := []rune{}
	toSmall := toLower(s)
	fmt.Println("****", toSmall)

	for _, char := range toSmall {
		if IsLetter(char) || IsDigit(char) {
			newStr = append(newStr, char)
			fmt.Println("//////", string(newStr))
		}
	}

	for i := 0; i < len(newStr); i++ {
		if newStr[i] != newStr[len(newStr)-i-1] {
			return false
		}
	}
	return true

}

func toLower(s string) string {
	str := []rune(s)
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			str[i] = char + 'a' - 'A'
		} else {
			str[i] = char
		}
	}
	fmt.Println("****", string(str))
	return string(str)
}

func IsLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}
