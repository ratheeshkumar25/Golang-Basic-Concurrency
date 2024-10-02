package main

import "fmt"

func main() {

	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	group := groupAnagram(strs)
	fmt.Println("group of anagram", group)

}

func groupAnagram(strs []string) [][]string {
	var sortString = func(s string) string {
		r := []rune(s)
		for i := 0; i < len(r); i++ {
			for j := 0; j < len(r)-i-1; j++ {
				if r[j] > r[j+1] {
					r[j], r[j+1] = r[j+1], r[j]
				}

			}
		}
		return string(r)
	}

	findAnagram := make(map[string][]string)

	for _, str := range strs {
		sortStr := sortString(str)
		findAnagram[sortStr] = append(findAnagram[sortStr], str)
	}

	var result [][]string

	for _, group := range findAnagram {
		result = append(result, group)
	}
	return result
}
