package main

import "fmt"

func heapSort(arr []int) {
	n := len(arr)
	buildHeap(arr, n)
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}

}

func buildHeap(arr []int, n int) {
	for i := (n / 2) - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
}

func heapify(arr []int, idx int, n int) {
	lar := idx
	left := 2*idx + 1
	right := 2*idx + 2

	if left < n && arr[left] > arr[lar] {
		lar = left
	}
	if right < n && arr[right] > arr[lar] {
		lar = right
	}
	if lar != idx {
		arr[idx], arr[lar] = arr[lar], arr[idx]
		heapify(arr, lar, n)
	}
}

func findthKthelement(arr []int, k int) int {
	heapSort(arr)
	return arr[len(arr)-k]
}

func main() {
	arr := []int{5, 2, 8, 4, 2, 7, 1, 6, 3} //12345678
	k := findthKthelement(arr, 3)
	fmt.Println("kth eleement is", k)

}
