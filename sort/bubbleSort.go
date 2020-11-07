package main

import "fmt"

// BubbleSort sort a slice of integers
func BubbleSort(integers []int) {
	size := len(integers)
	for i := 0; i < size-1; i++ { // O(n*n)
		for j := 0; j < size-i-1; j++ {
			swap(integers, j)
		}
	}
	printArr(integers)
}

func swap(integers []int, i int) {
	if integers[i] > integers[i+1] {
		// swapping integers[i] = integers[i+1] && integers[i+1] = integers[i]
		integers[i], integers[i+1] = integers[i+1], integers[i]
	}
}

func printArr(arr []int) {
	fmt.Print("Sorted array: ")
	fmt.Println(arr)
}

func main() {
	arr := []int{43, 54, 7, 2, 1, 0, 6, 88}
	BubbleSort(arr)
	arr = []int{3, 44, 908, 51, 28, -12, 7, 89, 12, -3}
	BubbleSort(arr)
}
