package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	containDup := containDuplicate(nums)
	fmt.Println("Duplicate contain in the array", containDup)

}

func containDuplicate(nums []int) bool {
	var sortNum = func(nums []int) []int {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums)-i-1; j++ {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
		return nums
	}

	sortArr := sortNum(nums)
	dupMap := make(map[int]bool)

	for _, val := range sortArr {
		if _, ok := dupMap[val]; ok {
			return true
		}
		dupMap[val] = true
	}
	return false

}
