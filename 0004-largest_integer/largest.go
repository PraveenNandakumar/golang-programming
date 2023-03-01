// Largest element in a array

package main

import "fmt"

func main() {
	var arr []int
	var n, largest int

	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&n)

	arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	largest = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	fmt.Printf("Largest element is %d", largest)
}
