package main

import "fmt"

func bubbleSort(arr [5]int, size int) [5]int {

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = swap(arr[j], arr[j+1])
			}
		}

	}

	return arr
}

func swap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp

	return a, b
}

func printlist(arr [5]int) {
	for _, k := range arr {
		fmt.Print(k, " ")
	}
}

func main() {
	list := [5]int{41, 51, 21, 11, 0}
	size := len(list)

	fmt.Println("Unsorted List :")
	printlist(list)
	fmt.Println()

	sortedList := bubbleSort(list, size)
	fmt.Println("Sorted List :")
	printlist(sortedList)
}
