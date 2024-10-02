package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result

}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	//fmt.Println("ARRYLEFT", left)
	right := mergeSort(arr[mid:])
	//fmt.Println("ARRAYRIGHT", right)
	return merge(left, right)

}

func main() {
	arr := []int{7, 8, 2, 3, 4, 5, 3}
	sort := mergeSort(arr)
	fmt.Println(sort)
}
