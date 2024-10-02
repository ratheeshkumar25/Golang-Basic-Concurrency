package main

import "fmt"

func main() {
	arr := []int{5, 6, 7, 2, 3, 5, 8}
	heapSort(arr)
	fmt.Println("arr", arr)

	nth := nthLargest(arr, 5)
	fmt.Println(nth)

}

func heapSort(arr []int) {
	n := len(arr)
	buildHeap(arr, n)
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func buildHeap(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func heapify(arr []int, idx int, n int) {
	lar := n
	left := 2*n + 1
	right := 2*n + 2

	if left < idx && arr[left] > arr[lar] {
		lar = left
	}
	if right < idx && arr[right] > arr[lar] {
		lar = right
	}
	if lar != n {
		arr[n], arr[lar] = arr[lar], arr[n]
		heapify(arr, idx, lar)
	}
}

func nthLargest(arr []int, k int) int {
	heapSort(arr)
	return arr[len(arr)-k]
}
