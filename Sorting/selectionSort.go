package main

import "fmt"

func SelectionSort(arr [5]int, size int) [5]int {
	for i := 0; i < size-1; i++ {
		min_idx := i
		for j := i + 1; j < size; j++ {
			if arr[min_idx] < arr[j] {
				min_idx = j

			}
			temp := arr[j]
			arr[j] = arr[min_idx]
			arr[min_idx] = temp
		}

	}

	return arr
}

func PrintList(arr [5]int) {
	for _, k := range arr {
		fmt.Print(k, " ")
	}
}

func main() {
	unSorted := [5]int{50, 40, 30, 20, 10}
	size := len(unSorted)
	fmt.Println("Unsorted list :")
	PrintList(unSorted)
	fmt.Println()

	sortedList := SelectionSort(unSorted, size)
	fmt.Println("Sorted list :")
	PrintList(sortedList)
}
